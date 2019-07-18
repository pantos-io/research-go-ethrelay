package testimonium

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"math/big"
	"strconv"
	"time"
)

type Chain struct {
	client *ethclient.Client
	contractAddress common.Address
	contract *Testimonium
}

type Client struct {
	chains map[uint8]*Chain
	account common.Address
	privateKey *ecdsa.PrivateKey
}

type BlockHeader struct {
	Parent                    [32]byte
	StateRoot                 [32]byte
	TransactionsRoot          [32]byte
	ReceiptsRoot              [32]byte
	BlockNumber               *big.Int
	RlpHeaderHashWithoutNonce [32]byte
	Nonce                     *big.Int
	LockedUntil               *big.Int
	TotalDifficulty           *big.Int
	OrderedIndex              *big.Int
	IterableIndex             *big.Int
	LatestFork                [32]byte
}

func (header BlockHeader) String() string {
	return fmt.Sprintf(`BlockHeader: {
Parent: %s,
StateRoot: %s,
TransactionsRoot: %s,
ReceiptsRoot: %s,
BlockNumber: %s,
RlpHeaderHashWithoutNonce: %s,
Nonce: %s,
LockedUntil: %s,
TotalDifficulty: %s,
OrderedIndex: %s,
IterableIndex: %s,
LatestFork: %s }`,
common.Bytes2Hex(header.Parent[:]),
common.Bytes2Hex(header.StateRoot[:]),
common.Bytes2Hex(header.TransactionsRoot[:]),
common.Bytes2Hex(header.ReceiptsRoot[:]),
header.BlockNumber.String(),
common.Bytes2Hex(header.RlpHeaderHashWithoutNonce[:]),
header.Nonce.String(),
header.LockedUntil.String(),
header.TotalDifficulty.String(),
header.OrderedIndex.String(),
header.IterableIndex.String(),
common.Bytes2Hex(header.LatestFork[:]))
}

func (t TestimoniumSubmitBlockHeader) String() string {
	return fmt.Sprintf("SubmitBlockHeaderEvent: { Hash: %s, HashWithoutNonce: %s, Nonce: %s, Parent: %s }",
		common.BytesToHash(t.Hash[:]).String(),
		common.BytesToHash(t.HashWithoutNonce[:]).String(),
		t.Nonce.String(),
		common.BytesToHash(t.Parent[:]).String())
}

func (event TestimoniumRemoveBranch) String() string {
	return fmt.Sprintf("RemoveBranchEvent: { Root: %s }", common.BytesToHash(event.Root[:]).String())
}

func NewClient(privateKey string, chainsConfig map[string]interface{}) *Client {
	client := new(Client)
	client.chains = make(map[uint8]*Chain)

	for k,v := range chainsConfig {
		chainId, err := strconv.ParseInt(k,10, 8)
		if err != nil {
			log.Fatal(err)
		}

		chainConfig := v.(map[string]interface{})

		// create client connection
		var ethClient *ethclient.Client
		fullUrl, err := createConnectionUrl(chainConfig)
		if err != nil {
			fmt.Printf("WARNING: No url specified for chain %d\n", chainId)
			continue
		}

		ethClient, err = ethclient.Dial(fullUrl)
		if err != nil {
			fmt.Printf("WARNING: Cannot connect to chain %d (%s): %s\n", chainId, fullUrl, err)
			continue	// --> even if we cannot connect to this chain, we still try to connect to the other ones
		}

		// create contract instance
		addressHex := chainConfig["address"].(string)
		address := common.HexToAddress(addressHex)
		var contract *Testimonium
		contract, err = NewTestimonium(address, ethClient)
		if err != nil {
			fmt.Printf("WARNING: No Testimonium contract deployed at address %s on chain %d (%s)\n", addressHex, chainId, fullUrl)
		}
		client.chains[uint8(chainId)] = &Chain{ethClient, address, contract}
	}

	// get public address
	privateKeyBytes, err := hexutil.Decode(privateKey);
	if err != nil {
		log.Fatal(err)
	}
	ecdsaPrivateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	client.privateKey = ecdsaPrivateKey
	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	client.account = crypto.PubkeyToAddress(*publicKeyECDSA)

	return client
}

func createConnectionUrl(chainConfig map[string]interface{}) (string, error) {
	fullUrl := ""
	if chainConfig["type"] != nil {
		fullUrl += chainConfig["type"].(string) + "://"
	} else {
		fullUrl += "https://"
	}
	if chainConfig["url"] == nil {
		return "", fmt.Errorf("no url specified")
	}

	fullUrl += chainConfig["url"].(string)
	if chainConfig["port"] != nil {
		fullUrl = fmt.Sprintf("%s:%d", fullUrl, chainConfig["port"].(int))
	}
	return fullUrl, nil
}

func (c Client) Chains() []uint8 {
	keys := make([]uint8, len(c.chains))

	i := 0
	for k := range c.chains {
		keys[i] = k
		i++
	}
	return keys
}

func (c Client) Account() string {
	return c.account.Hex()
}

func (c Client) TotalBalance() (*big.Int, error) {
	var totalBalance = new(big.Int);
	for k,_ := range c.chains {
		balance, err := c.Balance(k)
		totalBalance.Add(totalBalance, balance)
		if err != nil {
			return nil, err
		}
	}
	return totalBalance, nil
}

func (c Client) Balance(chainId uint8) (*big.Int, error) {
	var totalBalance = new(big.Int);
	_, exists := c.chains[chainId]
	if !exists {
		return nil, fmt.Errorf("chain %s does not exist", chainId)
	}
	balance, err := c.chains[chainId].client.BalanceAt(context.Background(), c.account, nil)
	totalBalance.Add(totalBalance, balance)
	if err != nil {
		return nil, err
	}
	return totalBalance, nil
}

func (c Client) BlockHeaderExists(blockHash [32]byte, chain uint8) (bool, error) {
	return c.chains[chain].contract.IsBlock(nil, blockHash)
}

func (c Client) BlockHeader(blockHash [32]byte, chain uint8) (BlockHeader, error) {
	return c.chains[chain].contract.Headers(nil, blockHash)
}

func (c Client) OriginalBlockHeader(blockHash [32]byte, chain uint8) (*types.Block, error) {
	return c.chains[chain].client.BlockByHash(context.Background(), common.BytesToHash(blockHash[:]))
}

func (c Client) SubmitHeader(header *types.Header, chain uint8) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	rlpHeader, err := encodeHeaderToRLP(header)
	if err != nil {
		log.Fatal("Failed to encode header to RLP: " + err.Error())
	}
	//fmt.Printf("RLP: 0x%s\n", hex.EncodeToString(rlpHeader))
	c.SubmitRLPHeader(rlpHeader, chain)
}

func (c Client) SubmitRLPHeader(rlpHeader []byte, chain uint8) {
	// Check preconditions
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	// Submit Transfer Transaction
	auth := prepareTransaction(c.account, c.privateKey, c.chains[chain])
	tx, err := c.chains[chain].contract.SubmitHeader(auth, rlpHeader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	receipt, err := awaitTxReceipt(c.chains[chain].client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(c.chains[chain].client, c.account, tx, receipt.BlockNumber)
		fmt.Printf("Tx failed: %s\n", reason)
		return
	}

	// Transaction is successful
	event, err := c.chains[chain].contract.TestimoniumFilterer.ParseSubmitBlockHeader(*receipt.Logs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tx successful: %s\n", event.String())
}

func (c Client) Block(blockHash common.Hash, chain uint8) (*types.Block, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}
	return c.chains[chain].client.BlockByHash(context.Background(), blockHash)
}

func (c Client) HeaderByNumber(blockNumber *big.Int, chain uint8) (*types.Header, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}
	return c.chains[chain].client.HeaderByNumber(context.Background(), blockNumber)
}

func (c Client) HeaderByHash(blockHash common.Hash, chain uint8) (*types.Header, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}
	return c.chains[chain].client.HeaderByHash(context.Background(), blockHash)
}

func (c Client) Transaction(txHash common.Hash, chain uint8) (*types.Transaction, bool, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}
	return c.chains[chain].client.TransactionByHash(context.Background(), txHash)
}

func (c Client) TransactionReceipt(txHash common.Hash, chain uint8) (*types.Receipt, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}
	return c.chains[chain].client.TransactionReceipt(context.Background(), txHash)
}

func (c Client) RandomizeHeader(header *types.Header, chain uint8) *types.Header {
	temp := header.TxHash
	header.TxHash = header.ReceiptHash
	header.ReceiptHash = header.Root
	header.Root = temp
	return header
}

func (c Client) DisputeBlock(blockHash [32]byte, chain uint8) {
	auth := prepareTransaction(c.account, c.privateKey, c.chains[chain])
	tx, err := c.chains[chain].contract.DisputeBlock(auth, blockHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	receipt, err := awaitTxReceipt(c.chains[chain].client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(c.chains[chain].client, c.account, tx, receipt.BlockNumber)
		fmt.Printf("Tx failed: %s\n", reason)
		return
	}

	// Transaction is successful
	event, err := c.chains[chain].contract.TestimoniumFilterer.ParseRemoveBranch(*receipt.Logs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tx successful: %s\n", event)
}

func (c Client) GenerateMerkleProof(txHash [32]byte, chain uint8) ([32]byte, [32]byte, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}
	txReceipt, err := c.chains[chain].client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}
	return txHash, txReceipt.BlockHash, nil
}

func (c Client) VerifyTransaction(txHash [32]byte, blockHash [32]byte, noOfConfirmations uint8, chain uint8) bool {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	isValid, err := c.chains[chain].contract.VerifyTransaction(nil, txHash, blockHash, noOfConfirmations)
	if err != nil {
		log.Fatal("Failed to verify transaction: " + err.Error())
	}
	return isValid
}

func getFailureReason(client *ethclient.Client, from common.Address, tx *types.Transaction, blockNumber *big.Int) string {
	code, err := client.CallContract(context.Background(), createCallMsgFromTransaction(from, tx), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(string(code[67:]))
}

func createCallMsgFromTransaction(from common.Address, tx *types.Transaction) ethereum.CallMsg {
	return ethereum.CallMsg{
		From: from,
		To: tx.To(),
		Gas: tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value: tx.Value(),
		Data: tx.Data(),
	}
}

func encodeHeaderToRLP(header *types.Header) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := rlp.Encode(buffer, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra,
		header.MixDigest,
		header.Nonce,
	})
	//fmt.Printf("F RLP: 0x%s\n", hex.EncodeToString(buffer.Bytes()))
	return buffer.Bytes(), err
}


func prepareTransaction(from common.Address, privateKey *ecdsa.PrivateKey, chain *Chain) *bind.TransactOpts {
	nonce, err := chain.client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := chain.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Nonce: %d, Suggested Gas Price: %d\n", nonce, gasPrice)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.From = from
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasPrice = gasPrice
	// one could also set the gas limit, however it seems that the right gas limit is only estimated
	// if the gas limit is not set specifically
	return auth
}

func awaitTxReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	const TimeoutLength = 2
	receipts := make(chan *types.Receipt)

	go func(chan *types.Receipt) {
		for ;; {
			receipt, _ := client.TransactionReceipt(context.Background(), txHash)

			if receipt != nil {
				receipts <- receipt
			}
		}
	}(receipts)

	select {
		case receipt := <- receipts:
			return receipt, nil
		case <- time.After(TimeoutLength * time.Minute):
			return nil, fmt.Errorf("timeout: did not receive receipt after %d minutes", TimeoutLength)
	}

	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{chain.contractAddress},
	//}
	//
	//logs := make(chan types.Log)
	//
	//sub, err := chain.client.SubscribeFilterLogs(context.Background(), query, logs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for {
	//	select {
	//	case err := <-sub.Err():
	//		return nil, err
	//	case vLog := <-logs:
	//		// if the transaction hash of the event does not equal the passed
	//		// transaction hash we continue listening
	//		if vLog.TxHash.Hex() != txHash.Hex() {
	//			break
	//		}
	//		return parseEvent(vLog)
	//	}
	//}
}

