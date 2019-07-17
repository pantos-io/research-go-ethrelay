package testimonium

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"io/ioutil"
	"log"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const abiFile = "Testimonium.abi"

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

type SubmitBlockHeaderEvent struct {//    event SubmitBlockHeader( bytes32 hash, bytes32 hashWithoutNonce, uint nonce, bytes32 parent );
	Hash [32]byte
	HashWithoutNonce [32]byte
	Nonce *big.Int
	Parent [32]byte
}

func (t SubmitBlockHeaderEvent) String() string {
	return fmt.Sprintf("SubmitBlockHeaderEvent: { Hash: %s, HashWithoutNonce: %s, Nonce: %s, Parent: %s }",
		common.BytesToHash(t.Hash[:]).String(),
		common.BytesToHash(t.HashWithoutNonce[:]).String(),
		t.Nonce.String(),
		common.BytesToHash(t.Parent[:]).String())
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
	event, err := parseEvent(*receipt.Logs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tx successful: %s\n", event.String())
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

func (c Client) SubmitHeader(blockNumber *big.Int, srcChain uint8, destChain uint8) {
	// Check preconditions
	if _, exists := c.chains[srcChain]; !exists {
		log.Fatalf("Source chain '%d' does not exist", srcChain)
	}
	if _, exists := c.chains[destChain]; !exists {
		log.Fatalf("Destination chain '%d' does not exist", destChain)
	}

	// todo: maybe also check that source and destination chain are different

	header, err := c.chains[srcChain].client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Submitting block %s of chain %d to chain %d...\n", header.Number.String(), srcChain, destChain)
	rlpHeader, err := encodeHeaderToRLP(header)
	//fmt.Printf("RLP: 0x%s\n", hex.EncodeToString(rlpHeader))
	c.SubmitRLPHeader(rlpHeader, destChain)
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

func (c Client) Watch() {

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
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice
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


func awaitSubmitBlockHeaderEvent(chain *Chain) (SubmitBlockHeaderEvent, error) {
	event, err := awaitEvent(chain, reflect.TypeOf(SubmitBlockHeaderEvent{}))
	return event.(SubmitBlockHeaderEvent), err
}

func awaitEvent(chain *Chain, desiredType reflect.Type) (fmt.Stringer, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{chain.contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := chain.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			return nil, err
		case vLog := <-logs:
			//var event interface{}
			event, err := parseEvent(vLog)
			if err != nil {
				return nil, err
			}
			if reflect.TypeOf(event) == desiredType {
				return event, nil
			}
			break
		}
	}
}

func readContractAbi(filename string) (abi.ABI, error)  {
	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		return abi.ABI{}, err
	}
	return abi.JSON(strings.NewReader(string(b)))
}

func parseEvent(vLog types.Log) (fmt.Stringer, error) {
	submitEventSignature := []byte("SubmitBlockHeader(bytes32,bytes32,uint256,bytes32)") // SubmitBlockHeader( bytes32 hash, bytes32 hashWithoutNonce, uint nonce, bytes32 parent );

	submitHash := crypto.Keccak256Hash(submitEventSignature)

	// read the contract abi from json file
	contractAbi, err := readContractAbi(abiFile)
	if err != nil {
		return nil, err
	}
	if vLog.Topics[0].Hex() == submitHash.Hex() {
		return parseSubmitEvent(contractAbi, vLog)
	}
	return nil, fmt.Errorf("retrieved unknown event")
}

func parseSubmitEvent(contractAbi abi.ABI, vLog types.Log) (SubmitBlockHeaderEvent, error) {
	event := SubmitBlockHeaderEvent{}
	err := contractAbi.Unpack(&event, "SubmitBlockHeader", vLog.Data)
	if err != nil {
		return event, err
	}
	return event, nil
}

