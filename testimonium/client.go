// This file contains functions called by the various commands. These functions are used to interact with smart contracts
// (Ethash, Testimonium)
// Authors: Marten Sigwart, Philipp Frauenthaler

package testimonium

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/pantos-io/go-ethrelay/ethereum/ethash"
	"github.com/pantos-io/go-ethrelay/ethereum/ethashsol"
	"github.com/pantos-io/go-ethrelay/typedefs"
)

type ChainConfig map[string]interface{}

type ChainsConfig map[uint8]ChainConfig

type Chain struct {
	client                     *ethclient.Client
	testimoniumContractAddress common.Address
	testimoniumContract        *Testimonium
	ethashContractAddress      common.Address
	ethashContract             *ethashsol.Ethashsol
	fullUrl                    string
}

type Client struct {
	chains     map[uint8]*Chain
	account    common.Address
	privateKey *ecdsa.PrivateKey
}

type Header struct {
	Hash            [32]byte
	BlockNumber     *big.Int
	TotalDifficulty *big.Int
}

type FullHeader struct {
	Parent                    [32]byte
	UncleHash                 [32]byte
	StateRoot                 [32]byte
	TransactionsRoot          [32]byte
	ReceiptsRoot              [32]byte
	BlockNumber               *big.Int
	GasLimit                  *big.Int
	GasUsed                   *big.Int
	RlpHeaderHashWithoutNonce [32]byte
	Timestamp                 *big.Int
	Nonce                     *big.Int
	Difficulty                *big.Int
	ExtraData                 *byte
}

type VerificationResult struct {
	returnCode uint8
}

type TrieValueType int

const (
	VALUE_TYPE_TRANSACTION TrieValueType = 0
	VALUE_TYPE_RECEIPT     TrieValueType = 1
	VALUE_TYPE_STATE       TrieValueType = 2
)

func (header FullHeader) String() string {
	return fmt.Sprintf(`BlockHeader: {
Parent: %s,
StateRoot: %s,
TransactionsRoot: %s,
ReceiptsRoot: %s,
BlockNumber: %s,
RlpHeaderHashWithoutNonce: %s,
Nonce: %s,
Difficulty: %s }`,
		common.Bytes2Hex(header.Parent[:]),
		common.Bytes2Hex(header.StateRoot[:]),
		common.Bytes2Hex(header.TransactionsRoot[:]),
		common.Bytes2Hex(header.ReceiptsRoot[:]),
		header.BlockNumber.String(),
		common.Bytes2Hex(header.RlpHeaderHashWithoutNonce[:]),
		header.Nonce.String(),
		header.Difficulty.String())
}

func (t TestimoniumSubmitBlock) String() string {
	return fmt.Sprintf("SubmitBlockEvent: { Hash: %s }", common.BytesToHash(t.BlockHash[:]).String())
}

func (event TestimoniumRemoveBranch) String() string {
	return fmt.Sprintf("RemoveBranchEvent: { Root: %s }", common.BytesToHash(event.Root[:]).String())
}

func (event TestimoniumPoWValidationResult) String() string {
	return fmt.Sprintf("PoWValidationResultEvent: { returnCode: %d, errorInfo: %d }", event.ReturnCode, event.ErrorInfo)
}

func (result VerificationResult) String() string {
	return fmt.Sprintf("VerificationResult: { returnCode: %d }", result.returnCode)
}

func (event TestimoniumWithdrawStake) String() string {
	return fmt.Sprintf("TestimoniumWithdrawStakeEvent: { client: %s, withdrawnStake: %d }", common.Bytes2Hex(event.Client.Bytes()), event.WithdrawnStake)
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
	client := new(Client)
	client.chains = make(map[uint8]*Chain)

	for k, v := range chainsConfig {
		chainId, err := strconv.ParseInt(k, 10, 8)
		if err != nil {
			log.Fatal(err)
		}

		chainConfig := v.(map[string]interface{})

		// create client connection
		var ethClient *ethclient.Client
		fullUrl, err := createConnectionUrl(chainConfig)
		if err != nil {
			fmt.Printf("WARNING: Could not read url specified for chain %d (%s)\n", chainId, err)
			continue
		}

		ethClient, err = ethclient.Dial(fullUrl)
		if err != nil {
			fmt.Printf("WARNING: Cannot connect to chain %d (%s): %s\n", chainId, fullUrl, err)
			continue // --> even if we cannot connect to this chain, we still try to connect to the other ones
		}

		chain := new(Chain)
		chain.client = ethClient
		chain.fullUrl = fullUrl

		// create testimonium contract instance
		var testimoniumContract *Testimonium
		addressHex := chainConfig["ethrelayaddress"]
		if addressHex != nil {
			ethrelayAddress := common.HexToAddress(addressHex.(string))
			testimoniumContract, err = NewTestimonium(ethrelayAddress, ethClient)
			if err != nil {
				fmt.Printf("WARNING: No Testimonium contract deployed at address %s on chain %d (%s)\n", addressHex, chainId, fullUrl)
			} else {
				chain.testimoniumContract = testimoniumContract
				chain.testimoniumContractAddress = ethrelayAddress
			}
		}

		// create ethash contract instance
		var ethashContract *ethashsol.Ethashsol
		addressHex = chainConfig["ethashaddress"]
		if addressHex != nil {
			ethashAddress := common.HexToAddress(addressHex.(string))
			ethashContract, err = ethashsol.NewEthashsol(ethashAddress, ethClient)
			if err != nil {
				fmt.Printf("WARNING: No Ethash contract deployed at address %s on chain %d (%s)\n", addressHex, chainId, fullUrl)
			} else {
				chain.ethashContract = ethashContract
				chain.ethashContractAddress = ethashAddress
			}
		}

		client.chains[uint8(chainId)] = chain
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
	var totalBalance = new(big.Int)
	for k := range c.chains {
		balance, err := c.Balance(k)
		if err != nil {
			return nil, err
		}
		totalBalance.Add(totalBalance, balance)
	}
	return totalBalance, nil
}

func (c Client) Balance(chainId uint8) (*big.Int, error) {
	var totalBalance = new(big.Int)

	_, exists := c.chains[chainId]
	if !exists {
		return nil, fmt.Errorf("chain %s does not exist", chainId)
	}

	balance, err := c.chains[chainId].client.BalanceAt(context.Background(), c.account, nil)
	if err != nil {
		return nil, err
	}

	totalBalance.Add(totalBalance, balance)

	return totalBalance, nil
}

func (c Client) GetStake(chainId uint8) (*big.Int, error) {
	_, exists := c.chains[chainId]
	if !exists {
		return nil, fmt.Errorf("chain %s does not exist", chainId)
	}
	stake, err := c.chains[chainId].testimoniumContract.GetStake(
		&bind.CallOpts{
			From: c.account,
		})
	if err != nil {
		return nil, err
	}
	return stake, nil
}

func (c Client) DepositStake(chainId uint8, amountInWei *big.Int) error {
	_, exists := c.chains[chainId]
	if !exists {
		return fmt.Errorf("chain %s does not exist", chainId)
	}

	auth := prepareTransaction(c.account, c.privateKey, c.chains[chainId], amountInWei)

	_, err := c.chains[chainId].testimoniumContract.DepositStake(auth, amountInWei)
	if err != nil {
		return err
	}

	// fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	return nil
}

func (c Client) WithdrawStake(chainId uint8, amountInWei *big.Int) error {
	_, exists := c.chains[chainId]
	if !exists {
		return fmt.Errorf("chain %s does not exist", chainId)
	}

	auth := prepareTransaction(c.account, c.privateKey, c.chains[chainId], big.NewInt(0))

	tx, err := c.chains[chainId].testimoniumContract.WithdrawStake(auth, amountInWei)
	if err != nil {
		return err
	}

	// fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	receipt, err := awaitTxReceipt(c.chains[chainId].client, tx.Hash())
	if err != nil {
		return err
	}

	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(c.chains[chainId].client, c.account, tx, receipt.BlockNumber)
		return fmt.Errorf("Tx failed: %s\n", reason)
	}

	// Transaction is successful
	eventIterator, err := c.chains[chainId].testimoniumContract.TestimoniumFilterer.FilterWithdrawStake(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		return err
	}

	if eventIterator.Next() {
		// fmt.Printf("Tx successful: %s\n", eventIterator.Event.String())

		if eventIterator.Event.WithdrawnStake.Cmp(amountInWei) != 0 {
			return errors.New("withdraw not successful, reason: more than 'amount' stake is locked in contract")
		}

		return nil
	}

	return errors.New("uncaught error")
}

func (c Client) BlockHeaderExists(blockHash [32]byte, chain uint8) (bool, error) {
	_, exists := c.chains[chain]
	if !exists {
		return false, fmt.Errorf("chain %s does not exist", chain)
	}

	return c.chains[chain].testimoniumContract.IsHeaderStored(nil, blockHash)
}

func (c Client) GetLongestChainEndpoint(chain uint8) ([32]byte, error) {
	_, exists := c.chains[chain]
	if !exists {
		return [32]byte{}, fmt.Errorf("chain %s does not exist", chain)
	}

	return c.chains[chain].testimoniumContract.GetLongestChainEndpoint(nil)
}

func (c Client) GetBlockHeader(blockHash [32]byte, chain uint8) (Header, error) {
	return c.chains[chain].testimoniumContract.GetHeader(nil, blockHash)
}

func (c Client) GetOriginalBlockHeader(blockHash [32]byte, chain uint8) (*types.Block, error) {
	return c.chains[chain].client.BlockByHash(context.Background(), common.BytesToHash(blockHash[:]))
}

func (c Client) SubmitHeader(header *types.Header, chain uint8) error {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	fmt.Printf("Submitting block: \nNo: %s\nHash: %s\n", header.Number.String(), header.Hash().String())

	rlpHeader, err := encodeHeaderToRLP(header)
	if err != nil {
		log.Fatal("Failed to encode header to RLP: " + err.Error())
	}

	return c.SubmitRLPHeader(rlpHeader, chain)
}

func (c Client) SubmitHeaderLive(destinationChain uint8, sourceChain uint8, lockTime time.Duration) {
	// Check preconditions
	if _, exists := c.chains[destinationChain]; !exists {
		log.Fatalf("Chain '%d' does not exist", destinationChain)
	}

	if _, exists := c.chains[sourceChain]; !exists {
		log.Fatalf("Chain '%d' does not exist", sourceChain)
	}

	/*
		there is much more to care about here:
		- 	if the genesis block of the testimonium contract is not on the current main chain,
			it is unpossible to find a block from where we can start submitting new blocks.
			this should be taken into account when creating a contract by selecting a genesis block
			far enough in the past that is ensured to not get into this scenario. the best case would be
			to start at block 0 of the source blockchain as this can never be a forked orphan branch.
			if this is the case, the search would go to block no. 0 to realize the genesis block is not on
			a given branch. this could be prevented by including a contract-creation search and check if the
			given blockNumber is not smaller than the genesis block, else it is useless and we can stop immediately.
		- 	if new blocks occur on the main chain, it is not safe if thy get into the longest chain or not.
			a participant in the relay-network may not take the risk to submit transactions with a high transaction-
			cost if they will never get any fees for this block. here, the participant with the fastest
			internet-connection (receiving new blocks from the main-chain and sending the blocks to the
			destination chain only requires latency and no hard problems) can relay all new blocks and
			wins the race so all fees go to the single participant. maybe this is not a big problem as relays
			can be deployed in the cloud and run in a very cheap and cost-intensive way.
		-	at the very first time the search and submit-phase takes a very long time because first one
			has to go back all the blocks and than submit all single blocks. here, the only way this could
			be enhanced is batch processing, which the testimonium contract already supports and some kind of
			binary search to effectively search for the latest submitted block which is part of the longest chain
			and part in the testimonium contract
		-	there may be not enough stake to deposit all blocks so one have to wait until the blocks are unlocked and
			the stake is freed again.
		-	another one has already submitted the block between finding this block with the search and submitting the block
		-	implement a live mode by subscribing to new blocks, here the same errors as above can occur
	*/

	// as the relay makes only sense if it's on the latest state of the main chain and it should never get older
	// than a few blocks behind the main chain it should be more performant and easier to do a simple backward-scan
	// starting with the newest block of the main chain, on the other hand, a binary search is always very fast and
	// if in the backwards search more than log2(n) blocks are stored, a binary search is always faster, so maybe
	// implement a binary search as default here

	genesis, err := c.chains[destinationChain].testimoniumContract.GetGenesisBlockHash(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Getting sure ETH Relay genesis block 0x%s from destination chain %d really exists on source chain %d\n", common.Bytes2Hex(genesis[:]), sourceChain, destinationChain)

	// returns an error if genesis was not found
	_, err = c.chains[sourceChain].client.HeaderByHash(context.Background(), genesis)
	if err != nil {
		log.Fatal(err)
	}

	// at the beginning this is nil - which returns the most recent block
	var blockNumber *big.Int = nil
	one := big.NewInt(1)

	var header *types.Header = nil

	// find the most recent block that was already submitted
	for {
		// get newest, longest header from source chain
		header, err = c.HeaderByNumber(blockNumber, sourceChain)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nSearching for block No. %s from source chain %d on destination chain %d", header.Number.String(), sourceChain, destinationChain)

		isHeaderStored, err := c.chains[destinationChain].testimoniumContract.IsHeaderStored(nil, header.Hash())
		if err != nil {
			log.Fatal(err)
		}

		if isHeaderStored {
			break
		}

		blockNumber = header.Number

		blockNumber.Sub(blockNumber, one)
	}

	fmt.Printf("\n\nlatest block No. submitted to destination chain: %s\n\n", header.Number.String())

	requiredStake, err := c.chains[destinationChain].testimoniumContract.GetRequiredStakePerBlock(nil)
	if err != nil {
		log.Fatal(err)
	}

	stake, err := c.GetStake(destinationChain)
	if err != nil {
		log.Fatal(err)
	}

	// check if there is enough stake
	if stake.Cmp(requiredStake) < 0 {
		log.Fatal("not enough stake deposited")
	}

	// has to be bigger than one as we checked above
	// TODO: if stake is currently locked, the routine tries to use it, but an error occurs
	maxBlocksWithStake := big.NewInt(0)
	maxBlocksWithStake.Div(stake, requiredStake)

	// calculate max. block submissions with stake
	var queue []time.Time

	// blockNumber was updated, so the destination chain is a few blocks behind source chain - updating now
	if blockNumber != nil {
		// submit all blocks to the most recent one
		for {
			if len(queue) >= int(maxBlocksWithStake.Uint64()) {
				timeUntilNextBlockIsUnlocked := queue[0].Add(lockTime)
				waitingTime := timeUntilNextBlockIsUnlocked.Sub(time.Now())

				if waitingTime > 0 {
					fmt.Printf("all stake is locked, waiting for %fs to continue\n", waitingTime.Seconds())
					time.Sleep(waitingTime)
				}

				queue = queue[1:]
			}

			// increase by one as we only want blocks that are new
			blockNumber.Add(blockNumber, one)

			header, err := c.HeaderByNumber(blockNumber, sourceChain)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Stake queue-length: ", len(queue), "\n")

			// TODO: a check for enough free/unlocked stake is required here, though a time based workaround is already implemented
			err = c.SubmitHeader(header, destinationChain)
			if err != nil {
				log.Fatal(err)
			}

			// add now + 1m for latency and whatever
			queue = append(queue, time.Now().Add(time.Second))

			// get newest, longest header from source chain
			header, err = c.HeaderByNumber(nil, sourceChain)
			if err != nil {
				log.Fatal(err)
			}

			// we caught up all the blocks... continue
			if header.Number.Cmp(blockNumber) == 0 {
				break
			}
		}
	}

	fmt.Printf("\nstarting live mode...\n\n")

	headers := make(chan *types.Header)

	sub, err := c.chains[sourceChain].client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			if len(queue) >= int(maxBlocksWithStake.Uint64()) {
				timeUntilNextBlockIsUnlocked := queue[0].Add(lockTime)
				waitingTime := timeUntilNextBlockIsUnlocked.Sub(time.Now())

				if waitingTime > 0 {
					fmt.Printf("all stake is locked, waiting for %fs to continue\n", waitingTime.Seconds())
					time.Sleep(waitingTime)
				}

				queue = queue[1:]
			}

			fmt.Println("Stake queue-length: ", len(queue), "\n")

			err = c.SubmitHeader(header, destinationChain)
			if err != nil {
				log.Fatal(err)
			}

			queue = append(queue, time.Now().Add(time.Second))
		}
	}
}

func (c Client) SubmitRLPHeader(rlpHeader []byte, chain uint8) error {
	// Check preconditions
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	// Submit Transfer Transaction
	auth := prepareTransaction(c.account, c.privateKey, c.chains[chain], big.NewInt(0))
	//auth.GasLimit = lastBlock.GasLimit()
	tx, err := c.chains[chain].testimoniumContract.SubmitBlock(auth, rlpHeader)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	receipt, err := awaitTxReceipt(c.chains[chain].client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(c.chains[chain].client, c.account, tx, receipt.BlockNumber)
		// fmt.Printf("Tx failed: %s\n", reason)
		return errors.New(reason)
	}

	// Transaction is successful
	eventIterator, err := c.chains[chain].testimoniumContract.TestimoniumFilterer.FilterSubmitBlock(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	// TODO: is this really the next event on the same chain? what if a transaction is included into one block,
	//  but get's overtaken by a submit-event of another fork? if it is guaranteed that the client is not reconnecting
	//  to other nodes within a usage - this may also be the case on every other transaction call
	//  workaround: check that the transaction from eventIterator's event is the same as the submitted transaction above
	if eventIterator.Next() {
		// fmt.Printf("Tx successful: %s\n", eventIterator.Event.String())

		// TODO: this is only 1 special hash value emitted by the contract for too small stake and not a read error code
		if eventIterator.Event.BlockHash == [32]byte{0} {
			return errors.New("block was not submitted, reason: too small stake deposited")
		}

		return nil
	}

	return errors.New("uncaught error")
}

func (c Client) BlockByHash(blockHash common.Hash, chain uint8) (*types.Block, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	return c.chains[chain].client.BlockByHash(context.Background(), blockHash)
}

func (c Client) BlockByNumber(blockNumber uint64, chain uint8) (*types.Block, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	return c.chains[chain].client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
}

func (c Client) HeaderByNumber(blockNumber *big.Int, chain uint8) (*types.Header, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	return c.chains[chain].client.HeaderByNumber(context.Background(), blockNumber)
}

type TotalDifficulty struct {
	TotalDifficulty string `json:"totalDifficulty"       gencodec:"required"`
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}

	return hexutil.EncodeBig(number)
}

func (c Client) TotalDifficulty(blockNumber *big.Int, chain uint8) (*big.Int, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	client, err := rpc.Dial(c.chains[chain].fullUrl)
	if err != nil {
		log.Fatal("Failed to connect to chain", err)
	}

	var totalDifficulty *TotalDifficulty
	err = client.CallContext(context.Background(), &totalDifficulty, "eth_getBlockByNumber", toBlockNumArg(blockNumber), false)
	if err == nil && totalDifficulty == nil {
		return big.NewInt(0), ethereum.NotFound
	}

	diff, err := hexutil.DecodeBig(totalDifficulty.TotalDifficulty)
	if err != nil {
		return big.NewInt(0), err
	}

	return diff, nil
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

func getRlpHeaderByTestimoniumSubmitEvent(chain *Chain, blockHash [32]byte) ([]byte, error) {
	eventIterator, err := chain.testimoniumContract.FilterSubmitBlock(nil)
	if err != nil {
		return nil, err
	}

	// TODO: the search could be enhanced if we use index event parameters, but this causes a little more cost and changes to the contract,
	//  evaluation is necessary - neglected in the first case, as the usual case is, that the event is at most the lock period behind,
	//  because it is not meant to need this for anything else than disputing

	for eventIterator.Next() {
		// according to the contract, the submit header event has exactly one parameter/data-item that is submitted
		// as no block-hash can be submitted twice, we found an event if the event data does equal the block-hash
		if bytes.Equal(eventIterator.Event.Raw.Data, blockHash[:]) {

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

			// parse method-id, the first 4 bytes are always the first 4 bytes of the encoded message signature
			methodId := txData[0:4]
			methodInputs := txData[4:]

			// load contract ABI
			testimoniumAbi, err := abi.JSON(strings.NewReader(TestimoniumABI))
			if err != nil {
				return nil, err
			}

			// recover method from signature and ABI
			method, err := testimoniumAbi.MethodById(methodId)
			if err != nil {
				return nil, err
			}

			// unpack method inputs
			parameter, err := method.Inputs.Unpack(methodInputs)
			if err != nil {
				return nil, err
			}

			return parameter[0].([]byte), nil
		}
	}

	return nil, fmt.Errorf("no submit event for block '%s' found", common.Bytes2Hex(blockHash[:]))
}

func (c Client) DisputeBlock(blockHash [32]byte, chain uint8) {
	fmt.Println("Disputing block ...")

	rlpEncodedBlockHeader, err := getRlpHeaderByTestimoniumSubmitEvent(c.chains[chain], blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// decode block header from rlp encoded block header
	blockHeader, err := decodeHeaderFromRLP(rlpEncodedBlockHeader)
	if err != nil {
		log.Fatal(err)
	}

	// the last thing needed for calling dispute is the parent rlp encoded block header
	rlpEncodedParentBlockHeader, err := getRlpHeaderByTestimoniumSubmitEvent(c.chains[chain], blockHeader.ParentHash)
	if err != nil {
		log.Fatal(err)
	}

	auth := prepareTransaction(c.account, c.privateKey, c.chains[chain], big.NewInt(0))

	// take the encoded block header and encode it without the nonce and the mixed hash
	blockHeaderWithoutNonce, err := encodeHeaderWithoutNonceToRLP(blockHeader)
	if err != nil {
		log.Fatal(err)
	}

	// create a hash to get the block hash without nonce needed for the ethash metadata construction
	blockHeaderHashWithoutNonce := crypto.Keccak256(blockHeaderWithoutNonce)

	// keccak256 returns a dynamic byte array, but we need a 32 byte fixed size byte array for the ethash block meta data
	var blockHeaderHashWithoutNonceLength32 [32]byte
	copy(blockHeaderHashWithoutNonceLength32[:], blockHeaderHashWithoutNonce)

	// get DAG and compute dataSetLookup and witnessForLookup
	blockMetaData := ethash.NewBlockMetaData(blockHeader.Number.Uint64(), blockHeader.Nonce.Uint64(), blockHeaderHashWithoutNonceLength32)
	dataSetLookUp := blockMetaData.DAGElementArray()
	witnessForLookup := blockMetaData.DAGProofArray()

	tx, err := c.chains[chain].testimoniumContract.DisputeBlockHeader(auth, rlpEncodedBlockHeader, blockHeaderHashWithoutNonceLength32, rlpEncodedParentBlockHeader, dataSetLookUp, witnessForLookup)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

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

	// get RemoveBranch event
	eventIteratorRemoveBranch, err := c.chains[chain].testimoniumContract.TestimoniumFilterer.FilterRemoveBranch(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	if eventIteratorRemoveBranch.Next() {
		fmt.Printf("Tx successful: %s\n", eventIteratorRemoveBranch.Event.String())
	}

	// get PoW Verification event
	eventIteratorPoWResult, err := c.chains[chain].testimoniumContract.TestimoniumFilterer.FilterPoWValidationResult(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	if eventIteratorPoWResult.Next() {
		fmt.Printf("Tx successful: %s\n", eventIteratorPoWResult.Event.String())
	}
}

func (c Client) GetRequiredVerificationFee(chain uint8) (*big.Int, error) {
	return c.chains[chain].testimoniumContract.GetRequiredVerificationFee(nil)
}

func (c Client) GenerateMerkleProofForTx(txHash [32]byte, chain uint8) ([]byte, []byte, []byte, []byte, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	txReceipt, err := c.chains[chain].client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return []byte{}, []byte{}, []byte{}, []byte{}, err
	}

	block, err := c.chains[chain].client.BlockByHash(context.Background(), txReceipt.BlockHash)
	if err != nil {
		return []byte{}, []byte{}, []byte{}, []byte{}, err
	}

	// create transactions trie
	indexBuffer := new(bytes.Buffer)
	txBuffer := new(bytes.Buffer)
	merkleTrie := new(trie.Trie)
	transactions := block.Transactions()

	for i := 0; i < transactions.Len(); i++ {
		indexBuffer.Reset()
		txBuffer.Reset()
		err = rlp.Encode(indexBuffer, uint(i))
		if err != nil {
			return []byte{}, []byte{}, []byte{}, []byte{}, err
		}
		transactions.EncodeIndex(i, txBuffer)
		merkleTrie.Update(indexBuffer.Bytes(), txBuffer.Bytes())
	}

	txBuffer.Reset()
	transactions.EncodeIndex(int(txReceipt.TransactionIndex), txBuffer)

	// create Merkle proof
	rlpEncodedTx := txBuffer.Bytes()

	indexBuffer.Reset()
	err = rlp.Encode(indexBuffer, txReceipt.TransactionIndex)
	if err != nil {
		return []byte{}, []byte{}, []byte{}, []byte{}, err
	}
	path := make([]byte, len(indexBuffer.Bytes()))
	copy(path, indexBuffer.Bytes())

	merkleIterator := merkleTrie.NodeIterator(nil)
	var proofNodes [][]byte
	for merkleIterator.Next(true) {
		if merkleIterator.Leaf() && bytes.Equal(merkleIterator.LeafKey(), path) {
			// leaf node representing tx has been found --> create Merkle proof
			proofNodes = merkleIterator.LeafProof()
			break
		}
	}

	indexBuffer.Reset()
	err = rlp.Encode(indexBuffer, proofNodes)
	if err != nil {
		return []byte{}, []byte{}, []byte{}, []byte{}, err
	}
	rlpEncodedProofNodes := make([]byte, len(indexBuffer.Bytes()))
	copy(rlpEncodedProofNodes, indexBuffer.Bytes())

	indexBuffer.Reset()
	err = rlp.Encode(indexBuffer, block.Header())
	if err != nil {
		return []byte{}, []byte{}, []byte{}, []byte{}, err
	}
	rlpEncodedHeader := make([]byte, len(indexBuffer.Bytes()))
	copy(rlpEncodedHeader, indexBuffer.Bytes())

	return rlpEncodedHeader, rlpEncodedTx, path, rlpEncodedProofNodes, nil
}

func (c Client) GenerateMerkleProofForReceipt(txHash [32]byte, chain uint8) ([]byte, []byte, []byte, []byte, error) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	txReceipt, err := c.chains[chain].client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return []byte{}, []byte{}, []byte{}, []byte{}, err
	}

	block, err := c.chains[chain].client.BlockByHash(context.Background(), txReceipt.BlockHash)
	if err != nil {
		return []byte{}, []byte{}, []byte{}, []byte{}, err
	}

	var path []byte
	var rlpEncodedReceipt []byte

	// create receipts trie
	buffer := new(bytes.Buffer)
	merkleTrie := new(trie.Trie)
	for i := 0; i < block.Transactions().Len(); i++ {
		tx := block.Body().Transactions[i]

		receipt, err := c.chains[chain].client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return []byte{}, []byte{}, []byte{}, []byte{}, err
		}

		buffer.Reset()
		receipt.EncodeRLP(buffer)
		encodedReceipt := make([]byte, len(buffer.Bytes()))
		copy(encodedReceipt, buffer.Bytes())

		buffer.Reset()
		rlp.Encode(buffer, uint(i))

		if txReceipt.TxHash == receipt.TxHash {
			path = make([]byte, len(buffer.Bytes()))
			copy(path, buffer.Bytes())

			rlpEncodedReceipt = make([]byte, len(encodedReceipt))
			copy(rlpEncodedReceipt, encodedReceipt)
		}

		merkleTrie.Update(buffer.Bytes(), encodedReceipt)
	}

	// create Merkle proof

	merkleIterator := merkleTrie.NodeIterator(nil)
	var proofNodes [][]byte
	for merkleIterator.Next(true) {
		if merkleIterator.Leaf() && bytes.Equal(merkleIterator.LeafKey(), path) {
			// leaf node representing tx has been found --> create Merkle proof
			proofNodes = merkleIterator.LeafProof()
			break
		}
	}

	buffer.Reset()
	rlp.Encode(buffer, proofNodes)
	rlpEncodedProofNodes := make([]byte, len(buffer.Bytes()))
	copy(rlpEncodedProofNodes, buffer.Bytes())

	buffer.Reset()
	rlp.Encode(buffer, block.Header())
	rlpEncodedHeader := make([]byte, len(buffer.Bytes()))
	copy(rlpEncodedHeader, buffer.Bytes())

	return rlpEncodedHeader, rlpEncodedReceipt, path, rlpEncodedProofNodes, nil
}

func (c Client) VerifyMerkleProof(feeInWei *big.Int, rlpHeader []byte, trieValueType TrieValueType, rlpEncodedValue []byte, path []byte,
	rlpEncodedProofNodes []byte, noOfConfirmations uint8, chain uint8) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	var tx *types.Transaction
	var err error
	auth := prepareTransaction(c.account, c.privateKey, c.chains[chain], feeInWei)

	switch trieValueType {
	case VALUE_TYPE_TRANSACTION:
		tx, err = c.chains[chain].testimoniumContract.VerifyTransaction(auth, feeInWei, rlpHeader,
			noOfConfirmations, rlpEncodedValue, path, rlpEncodedProofNodes)
	case VALUE_TYPE_RECEIPT:
		tx, err = c.chains[chain].testimoniumContract.VerifyReceipt(auth, feeInWei, rlpHeader, noOfConfirmations,
			rlpEncodedValue, path, rlpEncodedProofNodes)
	case VALUE_TYPE_STATE:
		tx, err = c.chains[chain].testimoniumContract.VerifyState(auth, feeInWei, rlpHeader, noOfConfirmations,
			rlpEncodedValue, path, rlpEncodedProofNodes)
	default:
		log.Fatal("Unexpected trie value type: ", trieValueType)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

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

	var verificationResult *VerificationResult

	switch trieValueType {
	case VALUE_TYPE_TRANSACTION:
		verificationResult, err = c.getVerifyTransactionEvent(chain, receipt)
	case VALUE_TYPE_RECEIPT:
		verificationResult, err = c.getVerifyReceiptEvent(chain, receipt)
	case VALUE_TYPE_STATE:
		verificationResult, err = c.getVerifyStateEvent(chain, receipt)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx successful: %s\n", verificationResult.String())
}

func (c Client) getVerifyTransactionEvent(chain uint8, receipt *types.Receipt) (*VerificationResult, error) {
	eventIterator, err := c.chains[chain].testimoniumContract.TestimoniumFilterer.FilterVerifyTransaction(
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

func (c Client) getVerifyReceiptEvent(chain uint8, receipt *types.Receipt) (*VerificationResult, error) {
	eventIterator, err := c.chains[chain].testimoniumContract.TestimoniumFilterer.FilterVerifyReceipt(
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

func (c Client) getVerifyStateEvent(chain uint8, receipt *types.Receipt) (*VerificationResult, error) {
	eventIterator, err := c.chains[chain].testimoniumContract.TestimoniumFilterer.FilterVerifyState(
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

func (c Client) SetEpochData(epochData typedefs.EpochData, chain uint8) {
	if _, exists := c.chains[chain]; !exists {
		log.Fatalf("Chain '%d' does not exist", chain)
	}

	nodes := []*big.Int{}
	start := big.NewInt(0)
	//fmt.Printf("No meaningful nodes: %d\n", len(epochData.MerkleNodes))
	for k, n := range epochData.MerkleNodes {
		nodes = append(nodes, n)
		if len(nodes) == 40 || k == len(epochData.MerkleNodes)-1 {
			mnlen := big.NewInt(int64(len(nodes)))
			fmt.Printf("Going to do tx\n")

			if k < 440 && epochData.Epoch.Uint64() == 128 {
				start.Add(start, mnlen)
				nodes = []*big.Int{}
				continue
			}

			auth := prepareTransaction(c.account, c.privateKey, c.chains[chain], big.NewInt(0))

			tx, err := c.chains[chain].ethashContract.SetEpochData(auth, epochData.Epoch, epochData.FullSizeIn128Resolution,
				epochData.BranchDepth, nodes, start, mnlen)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

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

			start.Add(start, mnlen)
			nodes = []*big.Int{}
		}
	}
}

func (c Client) DeployTestimonium(destinationChain uint8, sourceChain uint8, genesisBlockNumber uint64) common.Address {
	if _, exists := c.chains[destinationChain]; !exists {
		log.Fatalf("DestinationChain chain '%d' does not exist", destinationChain)
	}
	if _, exists := c.chains[sourceChain]; !exists {
		log.Fatalf("Source chain '%d' does not exist", sourceChain)
	}

	header, err := c.HeaderByNumber(new(big.Int).SetUint64(genesisBlockNumber), sourceChain)
	if err != nil {
		log.Fatal("Failed to retrieve header from source chain: " + err.Error())
	}

	totalDifficulty, err := c.TotalDifficulty(new(big.Int).SetUint64(genesisBlockNumber), sourceChain)
	if err != nil {
		log.Fatalf("Failed to retrieve total difficulty of block %d: %s", genesisBlockNumber, err)
	}

	rlpHeader, err := encodeHeaderToRLP(header)
	if err != nil {
		log.Fatal("Failed to encode header to RLP: " + err.Error())
	}

	auth := prepareTransaction(c.account, c.privateKey, c.chains[destinationChain], big.NewInt(0))

	addr, tx, _, err := DeployTestimonium(auth, c.chains[destinationChain].client, rlpHeader, totalDifficulty, c.chains[destinationChain].ethashContractAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	receipt, err := awaitTxReceipt(c.chains[destinationChain].client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(c.chains[destinationChain].client, c.account, tx, receipt.BlockNumber)
		fmt.Printf("Tx failed: %s\n", reason)
		return common.Address{}
	}

	fmt.Println("Contract has been deployed at address: ", addr.String())
	return addr
}

func (c Client) DeployEthash(destinationChain uint8) common.Address {
	if _, exists := c.chains[destinationChain]; !exists {
		log.Fatalf("DestinationChain chain '%d' does not exist", destinationChain)
	}

	auth := prepareTransaction(c.account, c.privateKey, c.chains[destinationChain], big.NewInt(0))

	addr, tx, _, err := ethashsol.DeployEthashsol(auth, c.chains[destinationChain].client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	receipt, err := awaitTxReceipt(c.chains[destinationChain].client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(c.chains[destinationChain].client, c.account, tx, receipt.BlockNumber)
		fmt.Printf("Tx failed: %s\n", reason)
		return common.Address{}
	}

	fmt.Println("Contract has been deployed at address: ", addr.String())

	return addr
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
