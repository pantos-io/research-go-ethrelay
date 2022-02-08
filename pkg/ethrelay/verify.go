package ethrelay

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

type VerificationResult struct {
	returnCode uint8
}

func (result VerificationResult) String() string {
	switch result.returnCode {
	case 0:
		return "Merkle proof verified"
	case 1:
		return "failed to verify Merkle proof"
	default:
		return fmt.Sprintf("VerificationResult: { returnCode: %d }", result.returnCode)
	}
}

func (c Client) BlockHeaderExists(chainId string, blockHash common.Hash) (bool, error) {
	return c.DstChain(chainId).ethrelay.IsHeaderStored(nil, blockHash)
}

func (c Client) GetOriginalBlockHeader(chainId string, blockHash common.Hash) (*types.Block, error) {
	return c.SrcChain(chainId).client.BlockByHash(context.Background(), common.BytesToHash(blockHash[:]))
}

func (c Client) GetRequiredVerificationFee(chainId string) (*big.Int, error) {
	return c.DstChain(chainId).ethrelay.GetRequiredVerificationFee(nil)
}

func (c Client) GenerateMerkleProofForTx(chainId string, txHash common.Hash) ([]byte, *MerkleProof, error) {
	chain := c.SrcChain(chainId)
	txReceipt, err := chain.client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return []byte{}, nil, err
	}

	block, err := chain.client.BlockByHash(context.Background(), txReceipt.BlockHash)
	if err != nil {
		return []byte{}, nil, err
	}

	rlpEncodedHeader, err := rlp.EncodeToBytes(block.Header())
	if err != nil {
		return []byte{}, nil, err
	}

	proof, err := NewMerkleProof(block.Transactions(), txReceipt.TransactionIndex)
	if err != nil {
		return []byte{}, nil, err
	}

	return rlpEncodedHeader, proof, nil
}

func receiptsByTransactions(ctx context.Context, c *ethclient.Client, transactions []*types.Transaction) (types.Receipts, error) {
	receipts := make([]*types.Receipt, len(transactions))
	for i := 0; i < len(receipts); i++ {
		receipt, err := c.TransactionReceipt(ctx, transactions[i].Hash())
		if err != nil {
			return nil, err
		}
		receipts[i] = receipt
	}
	return receipts, nil
}

func (c Client) GenerateMerkleProofForReceipt(chainId string, txHash common.Hash) ([]byte, *MerkleProof, error) {
	chain := c.SrcChain(chainId)
	txReceipt, err := chain.client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return []byte{}, nil, err
	}

	block, err := chain.client.BlockByHash(context.Background(), txReceipt.BlockHash)
	if err != nil {
		return []byte{}, nil, err
	}

	receipts, err := receiptsByTransactions(context.Background(), chain.client, block.Transactions())
	if err != nil {
		return []byte{}, nil, err
	}

	rlpEncodedHeader, err := rlp.EncodeToBytes(block.Header())
	if err != nil {
		return []byte{}, nil, err
	}

	proof, err := NewMerkleProof(receipts, txReceipt.TransactionIndex)
	if err != nil {
		return []byte{}, nil, err
	}

	return rlpEncodedHeader, proof, nil
}

func (c Client) VerifyMerkleProof(chainId string, feeInWei *big.Int, rlpHeader []byte, trieValueType TrieValueType,
	proof *MerkleProof, noOfConfirmations uint8) {

	var tx *types.Transaction
	var err error
	chain := c.DstChain(chainId)
	auth := prepareTransaction(c.account, c.privateKey, &chain.Chain, feeInWei)

	switch trieValueType {
	case ValueTypeTransaction:
		tx, err = chain.ethrelay.VerifyTransaction(auth, feeInWei, rlpHeader,
			noOfConfirmations, proof.Value, proof.Path, proof.Nodes)
	case ValueTypeReceipt:
		tx, err = chain.ethrelay.VerifyReceipt(auth, feeInWei, rlpHeader, noOfConfirmations,
			proof.Value, proof.Path, proof.Nodes)
	case ValueTypeState:
		tx, err = chain.ethrelay.VerifyState(auth, feeInWei, rlpHeader, noOfConfirmations,
			proof.Value, proof.Path, proof.Nodes)
	default:
		log.Fatal("Unexpected trie value type: ", trieValueType)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	receipt, err := awaitTxReceipt(chain.client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(chain.client, c.account, tx, receipt.BlockNumber)
		fmt.Printf("Tx failed: %s\n", reason)
		return
	}

	var verificationResult *VerificationResult

	switch trieValueType {
	case ValueTypeTransaction:
		verificationResult, err = c.getVerifyTransactionEvent(chainId, receipt)
	case ValueTypeReceipt:
		verificationResult, err = c.getVerifyReceiptEvent(chainId, receipt)
	case ValueTypeState:
		verificationResult, err = c.getVerifyStateEvent(chainId, receipt)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx successful: %s\n", verificationResult)
}

func (c Client) getVerifyTransactionEvent(chainId string, receipt *types.Receipt) (*VerificationResult, error) {
	eventIterator, err := c.DstChain(chainId).ethrelay.EthrelayFilterer.FilterVerifyTransaction(
		&bind.FilterOpts{
			Start:   receipt.BlockNumber.Uint64(),
			End:     nil,
			Context: nil,
		})
	if err != nil {
		return nil, err
	}
	if eventIterator.Next() {
		return &VerificationResult{
			returnCode: eventIterator.Event.Result,
		}, nil
	}
	return nil, fmt.Errorf("no event found")
}

func (c Client) getVerifyReceiptEvent(chainId string, receipt *types.Receipt) (*VerificationResult, error) {
	eventIterator, err := c.DstChain(chainId).ethrelay.EthrelayFilterer.FilterVerifyReceipt(
		&bind.FilterOpts{
			Start:   receipt.BlockNumber.Uint64(),
			End:     nil,
			Context: nil,
		})
	if err != nil {
		return nil, err
	}
	if eventIterator.Next() {
		return &VerificationResult{
			returnCode: eventIterator.Event.Result,
		}, nil
	}
	return nil, fmt.Errorf("no event found")
}

func (c Client) getVerifyStateEvent(chainId string, receipt *types.Receipt) (*VerificationResult, error) {
	eventIterator, err := c.DstChain(chainId).ethrelay.EthrelayFilterer.FilterVerifyState(
		&bind.FilterOpts{
			Start:   receipt.BlockNumber.Uint64(),
			End:     nil,
			Context: nil,
		})
	if err != nil {
		return nil, err
	}
	if eventIterator.Next() {
		return &VerificationResult{
			returnCode: eventIterator.Event.Result,
		}, nil
	}
	return nil, fmt.Errorf("no event found")
}