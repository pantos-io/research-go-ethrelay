package ethrelay

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pantos-io/go-ethrelay/pkg/ethereum/ethash"
)

func getRlpHeaderByEvent(chain *DestinationChain, blockHash common.Hash) ([]byte, error) {
	eventIterator, err := chain.ethrelay.FilterNewBlock(nil)
	if err != nil {
		return nil, err
	}

	// TODO: the search could be enhanced if we use index event parameters, but this causes a little more cost and changes to the contract,
	//  evaluation is necessary - neglected in the first case, as the usual case is, that the event is at most the lock period behind,
	//  because it is not meant to need this for anything else than disputing

	first := true
	for eventIterator.Next() {
		// As no block hash can be submitted twice, we found an event if the event data equals the block hash
		if bytes.Equal(eventIterator.Event.BlockHash[:], blockHash[:]) {

			// get the hash where the event was emitted
			txHash := eventIterator.Event.Raw.TxHash

			// get the full transaction by txhash
			tx, isPending, err := chain.client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				return nil, err
			}

			// if the transaction is pending, we don't know if it will be included
			if isPending {
				return nil, fmt.Errorf("transaction where block was submitted is currently pending...")
			}

			// get raw abi-encoded bytes of transaction data
			txData := tx.Data()

			// load contract ABI
			ethrelayAbi, err := abi.JSON(strings.NewReader(EthrelayABI))
			if err != nil {
				return nil, err
			}

			var method *abi.Method
			var inputs []byte

			// The constructor is always the first function that emits this event
			if first {
				method = &ethrelayAbi.Constructor

				// Constructor arguments are appended to the bytecode of the contract
				inputs = txData[len(common.FromHex(EthrelayMetaData.Bin)):]
			} else {
				// parse method-id, the first 4 bytes are always the first 4 bytes of the encoded message signature
				id := txData[0:4]
				inputs = txData[4:]

				// recover method from signature and ABI
				method, err = ethrelayAbi.MethodById(id)
				if err != nil {
					return nil, err
				}
			}

			// unpack method inputs
			parameter, err := method.Inputs.Unpack(inputs)
			if err != nil {
				return nil, err
			}

			return parameter[0].([]byte), nil
		}

		first = false
	}

	return nil, fmt.Errorf("no submit event for block '%s' found", common.Bytes2Hex(blockHash[:]))
}

func (c Client) DisputeBlock(chainId string, blockHash common.Hash) {
	chain := c.DstChain(chainId)

	fmt.Println("Disputing block...")

	rlpEncodedBlockHeader, err := getRlpHeaderByEvent(chain, blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// decode block header from rlp encoded block header
	blockHeader, err := decodeHeaderFromRLP(rlpEncodedBlockHeader)
	if err != nil {
		log.Fatal(err)
	}

	// the last thing needed for calling dispute is the parent rlp encoded block header
	rlpEncodedParentBlockHeader, err := getRlpHeaderByEvent(chain, blockHeader.ParentHash)
	if err != nil {
		log.Fatal(err)
	}

	auth := prepareTransaction(c.account, c.privateKey, &chain.Chain, big.NewInt(0))

	// take the encoded block header and encode it without the nonce and the mixed hash
	blockHeaderWithoutNonce, err := EncodeHeaderWithoutNonceToRLP(blockHeader)
	if err != nil {
		log.Fatal(err)
	}

	// create a hash to get the block hash without nonce needed for the ethash metadata construction
	blockHeaderHashWithoutNonce := crypto.Keccak256(blockHeaderWithoutNonce)

	// keccak256 returns a dynamic byte array, but we need a 32 byte fixed size byte array for the ethash block meta data
	var blockHeaderHashWithoutNonceLength32 common.Hash
	copy(blockHeaderHashWithoutNonceLength32[:], blockHeaderHashWithoutNonce)

	// get DAG and compute dataSetLookup and witnessForLookup
	blockMetaData := ethash.NewBlockMetaData(blockHeader.Number.Uint64(), blockHeader.Nonce.Uint64(), blockHeaderHashWithoutNonceLength32)
	dataSetLookUp := blockMetaData.DAGElementArray()
	witnessForLookup := blockMetaData.DAGProofArray()

	tx, err := chain.ethrelay.DisputeBlockHeader(auth, rlpEncodedBlockHeader, rlpEncodedParentBlockHeader, dataSetLookUp, witnessForLookup)
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

	// get RemoveBranch event
	eventIteratorRemoveBranch, err := chain.ethrelay.EthrelayFilterer.FilterRemoveBranch(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	if eventIteratorRemoveBranch.Next() {
		fmt.Printf("Tx successful: %s\n", eventIteratorRemoveBranch.Event)
	}

	// get PoW Verification event
	eventIteratorPoWResult, err := chain.ethrelay.EthrelayFilterer.FilterPoWValidationResult(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	if eventIteratorPoWResult.Next() {
		fmt.Printf("Tx successful: %s\n", eventIteratorPoWResult.Event)
	}
}