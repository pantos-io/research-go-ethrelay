package testimonium

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pantos-io/go-ethrelay/ethereum/ethashsol"
)

func (c Client) DeployTestimonium(dstChainId string, srcChainId string, genesisBlockNumber uint64) common.Address {
	header, err := c.HeaderByNumber(srcChainId, new(big.Int).SetUint64(genesisBlockNumber))
	if err != nil {
		log.Fatal("Failed to retrieve header from source chain: " + err.Error())
	}

	totalDifficulty, err := c.TotalDifficulty(srcChainId, new(big.Int).SetUint64(genesisBlockNumber))
	if err != nil {
		log.Fatalf("Failed to retrieve total difficulty of block %d: %s", genesisBlockNumber, err)
	}

	rlpHeader, err := encodeHeaderToRLP(header)
	if err != nil {
		log.Fatal("Failed to encode header to RLP: " + err.Error())
	}

	dstChain := c.DstChain(dstChainId)
	auth := prepareTransaction(c.account, c.privateKey, &dstChain.Chain, big.NewInt(0))

	addr, tx, _, err := DeployTestimonium(auth, dstChain.client, rlpHeader, totalDifficulty, dstChain.ethashAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	receipt, err := awaitTxReceipt(dstChain.client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(dstChain.client, c.account, tx, receipt.BlockNumber)
		fmt.Printf("Tx failed: %s\n", reason)
		return common.Address{}
	}

	fmt.Println("Contract has been deployed at address", addr)
	return addr
}

func (c Client) DeployEthash(chainId string) common.Address {
	chain := c.DstChain(chainId)
	auth := prepareTransaction(c.account, c.privateKey, &chain.Chain, big.NewInt(0))

	addr, tx, _, err := ethashsol.DeployEthashsol(auth, chain.client)
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
		return common.Address{}
	}

	fmt.Println("Contract has been deployed at address", addr)

	return addr
}

func (c Client) TotalDifficulty(chainId string, blockNumber *big.Int) (*big.Int, error) {
	client, err := rpc.Dial(c.Chain(chainId).fullUrl)
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

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}

	return hexutil.EncodeBig(number)
}