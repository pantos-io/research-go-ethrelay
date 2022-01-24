// This file contains functions called by the various commands. These functions are used to interact with smart contracts
// (Ethash, Testimonium)
// Authors: Marten Sigwart, Philipp Frauenthaler

package testimonium

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/pantos-io/go-ethrelay/ethereum/ethashsol"
)

type Chain struct {
	client                     *ethclient.Client
	fullUrl                    string
}

type SourceChain struct {
	Chain
}

type DestinationChain struct {
	Chain
	testimonium        	*Testimonium
	testimoniumAddress	common.Address
	ethash             	*ethashsol.Ethashsol
	ethashAddress		common.Address
}

type Client struct {
	chains		map[string]*Chain
	srcChains	map[string]*SourceChain
	dstChains	map[string]*DestinationChain
	account    	common.Address
	privateKey 	*ecdsa.PrivateKey
}

type Header struct {
	Hash            common.Hash
	BlockNumber     *big.Int
	TotalDifficulty *big.Int
}

type TrieValueType int
const (
	ValueTypeTransaction TrieValueType = iota
	ValueTypeReceipt
	ValueTypeState
)

type PoWValidationResult int
const (
	PoWValid		= 0
	PoWEpoch		= 1
	PoWDifficulty	= 2
)

func (event TestimoniumRemoveBranch) String() string {
	return fmt.Sprintf("branch with root hash %s removed", common.BytesToHash(event.Root[:]))
}

func (event TestimoniumPoWValidationResult) String() string {
	switch event.ReturnCode.Int64() {
	case PoWValid:
		return "PoW was successfully validated"
	case PoWEpoch:
		return fmt.Sprintf("epoch data for epoch %d not set", event.ErrorInfo)
	case PoWDifficulty:
		return fmt.Sprintf("calculated difficulty of %d too low", event.ErrorInfo)
	default:
		return fmt.Sprintf("PoWValidationResultEvent: { returnCode: %d, errorInfo: %d }", event.ReturnCode, event.ErrorInfo)
	}
}

func CreateChainConfig(connectionType string, connectionUrl string, connectionPort uint64) map[string]interface{} {
	chainConfig := make(map[string]interface{})

	chainConfig["url"] = connectionUrl

	if connectionType != "" {
		chainConfig["type"] = connectionType
	}
	if connectionPort != 0 {
		chainConfig["port"] = connectionPort
	}
	return chainConfig
}

func NewClient(privateKey string, chainsConfig map[string]interface{}) *Client {

	sources := chainsConfig["sources"].(map[string]interface{})
	destinations := chainsConfig["destinations"].(map[string]interface{})

	client := new(Client)
	client.chains = make(map[string]*Chain, len(sources) + len(destinations))
	client.srcChains = make(map[string]*SourceChain, len(sources))
	client.dstChains = make(map[string]*DestinationChain, len(destinations))

	fillInCommonProps := func(chain *Chain, chainConfig map[string]interface{}) error {
		// create client connection
		var ethClient *ethclient.Client
		fullUrl, err := createConnectionUrl(chainConfig)
		if err != nil {
			return fmt.Errorf("could not read url specified: %s", err);
		}
	
		ethClient, err = ethclient.Dial(fullUrl)
		if err != nil {
			// --> even if we cannot connect to this chain, we still try to connect to the other ones
			return fmt.Errorf("cannot connect to chain at %s: %s", fullUrl, err)
		}
	
		chain.client = ethClient
		chain.fullUrl = fullUrl
	
		return nil;
	}

	for chainId, v := range sources {
		chainConfig := v.(map[string]interface{})
		srcChain := new(SourceChain)

		fillInCommonProps(&srcChain.Chain, chainConfig)

		if _, exists := client.chains[chainId]; exists {
			fmt.Printf("WARNING: Duplicate chain ID '%s', overwriting previous configuration\n", chainId)
		}

		client.chains[chainId] = &srcChain.Chain
		client.srcChains[chainId] = srcChain
	}

	for chainId, v := range destinations {
		chainConfig := v.(map[string]interface{})
		dstChain := new(DestinationChain)

		fillInCommonProps(&dstChain.Chain, chainConfig)

		// Create ETH Relay contract instance
		addressHex := chainConfig["ethrelayaddress"]
		if addressHex == nil {
			fmt.Printf("WARNING: Address for ETH Relay instance for chain '%s' not configured\n", chainId)
		} else {
			ethrelayAddress := common.HexToAddress(addressHex.(string))
			testimoniumContract, err := NewTestimonium(ethrelayAddress, dstChain.client)
			if err == nil {
				dstChain.testimonium = testimoniumContract
				dstChain.testimoniumAddress = ethrelayAddress
			} else {
				fmt.Printf("WARNING: No ETH Relay contract deployed on chain '%s' at address %s: %s\n", chainId, addressHex, err)
			}
		}

		// Create Ethash contract instance
		addressHex = chainConfig["ethashaddress"]
		if addressHex == nil {
			fmt.Printf("WARNING: Address for Ethash instance for chain '%s' not configured\n", chainId)
		} else {
			ethashAddress := common.HexToAddress(addressHex.(string))
			ethashContract, err := ethashsol.NewEthashsol(ethashAddress, dstChain.client)
			if err == nil {
				dstChain.ethash = ethashContract
				dstChain.ethashAddress = ethashAddress
			} else {
				fmt.Printf("WARNING: No Ethash contract deployed on chain '%s' at address %s: %s\n", chainId, addressHex, err)
			}
		}

		if _, exists := client.chains[chainId]; exists {
			fmt.Printf("WARNING: Duplicate chain ID '%s', overwriting previous configuration\n", chainId)
		}

		client.chains[chainId] = &dstChain.Chain
		client.dstChains[chainId] = dstChain
	}

	// get public address
	privateKeyBytes, err := hexutil.Decode(privateKey)
	if err != nil {
		fmt.Println("Could not decode private key. Is it a correct hex string (0x...)?")
		os.Exit(1)
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
		// port can be parsed as int
		if port, ok := chainConfig["port"].(int); ok {
			fullUrl = fmt.Sprintf("%s:%d", fullUrl, port)
			return fullUrl, nil
		}

		// port is a string but could still be a legal port
		port, err := strconv.ParseUint(chainConfig["port"].(string), 10, 64)
		if err != nil {
			return "", fmt.Errorf("llegal port: %s", chainConfig["port"].(string))
		}
		fullUrl = fmt.Sprintf("%s:%d", fullUrl, port)
	}
	return fullUrl, nil
}

func (c Client) Chain(id string) *Chain {
	if _, exists := c.chains[id]; !exists {
		log.Fatalf("Chain '%s' does not exist", id)
	}
	return c.chains[id]
}

func (c Client) SrcChain(id string) *SourceChain {
	if _, exists := c.srcChains[id]; !exists {
		log.Fatalf("Source chain '%s' does not exist", id)
	}
	return c.srcChains[id]
}

func (c Client) DstChain(id string) *DestinationChain {
	if _, exists := c.dstChains[id]; !exists {
		log.Fatalf("Destination chain '%s' does not exist", id)
	}
	return c.dstChains[id]
} 

func (c Client) Chains() []string {
	keys := make([]string, len(c.chains))

	i := 0
	for k := range c.chains {
		keys[i] = k
		i++
	}
	return keys
}

func (c Client) BlockByHash(chainId string, blockHash common.Hash) (*types.Block, error) {
	return c.Chain(chainId).client.BlockByHash(context.Background(), blockHash)
}

func (c Client) HeaderByNumber(chainId string, blockNumber *big.Int) (*types.Header, error) {
	return c.Chain(chainId).client.HeaderByNumber(context.Background(), blockNumber)
}

func (c Client) HeaderByHash(chainId string, blockHash common.Hash) (*types.Header, error) {
	return c.Chain(chainId).client.HeaderByHash(context.Background(), blockHash)
}

type TotalDifficulty struct {
	TotalDifficulty string `json:"totalDifficulty"       gencodec:"required"`
}

func (c Client) Transaction(chainId string, txHash common.Hash) (*types.Transaction, bool, error) {
	return c.Chain(chainId).client.TransactionByHash(context.Background(), txHash)
}

func (c Client) TransactionReceipt(chainId string, txHash common.Hash) (*types.Receipt, error) {
	return c.Chain(chainId).client.TransactionReceipt(context.Background(), txHash)
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
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
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
		header.BaseFee,
	})

	// be careful when passing byte-array as buffer, the pointer can change if the buffer is used again
	return buffer.Bytes(), err
}

func decodeHeaderFromRLP(bytes []byte) (*types.Header, error) {
	header := new(types.Header)

	err := rlp.DecodeBytes(bytes, header)

	return header, err
}

func encodeHeaderWithoutNonceToRLP(header *types.Header) ([]byte, error) {
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
		header.BaseFee,
	})

	return buffer.Bytes(), err
}

func prepareTransaction(from common.Address, privateKey *ecdsa.PrivateKey, chain *Chain, valueInWei *big.Int) *bind.TransactOpts {
	nonce, err := chain.client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := chain.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.From = from
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = valueInWei // in wei
	auth.GasPrice = gasPrice

	// one could also set the gas limit, however it seems that the right gas limit is only estimated
	// if the gas limit is not set specifically
	return auth
}

func awaitTxReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	const TimeoutLength = 2
	receipts := make(chan *types.Receipt)

	go func(chan *types.Receipt) {
		for {
			receipt, _ := client.TransactionReceipt(context.Background(), txHash)

			if receipt != nil {
				receipts <- receipt
			}
		}
	}(receipts)

	select {
	case receipt := <-receipts:
		return receipt, nil
	case <-time.After(TimeoutLength * time.Minute):
		return nil, fmt.Errorf("timeout: did not receive receipt after %d minutes", TimeoutLength)
	}

	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{chain.testimoniumContractAddress},
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
