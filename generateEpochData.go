package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pantos-io/go-testimonium/ethereum/ethash"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 2 {
		log.Fatal("Provide start block number as first and number of blocks as second argument.")
	}

	client, err := ethclient.Dial("https://mainnet.infura.io")
	checkError(err)

	startBlock, err := strconv.ParseInt(argsWithoutProg[0], 10, 64)
	checkError(err)

	noOfBlocks, err := strconv.ParseInt(argsWithoutProg[1], 10, 64)
	checkError(err)

	f, err := os.Create("./epoch-data.json")
	checkError(err)
	defer f.Close()

	_, err = fmt.Fprint(f, "{\n")
	checkError(err)
	blockNumber := big.NewInt(startBlock)
	for i := 0; int64(i) <= noOfBlocks; i++ {
		header, err := client.HeaderByNumber(context.Background(), blockNumber)
		checkError(err)

		fmt.Println("create DAG, compute dataSetLookup and witnessForLookup")
		// get DAG and compute dataSetLookup and witnessForLookup
		hashWithoutNonce := sealHash(header)
		blockMetaData := ethash.NewBlockMetaData(header.Number.Uint64(), header.Nonce.Uint64(), hashWithoutNonce)
		dataSetLookup := blockMetaData.DAGElementArray()
		witnessForLookup := blockMetaData.DAGProofArray()

		_, err = fmt.Fprintf(f, "\"%d\": [\n", blockNumber)
		checkError(err)
		writeElementsToFile(f, dataSetLookup)
		_, err = fmt.Fprintf(f, ",\n")
		checkError(err)
		writeElementsToFile(f, witnessForLookup)
		_, err = fmt.Fprintf(f, "\n],\n")
		checkError(err)
		blockNumber.Add(blockNumber, big.NewInt(int64(1)))
	}
	_, err = fmt.Fprint(f, "}")
	checkError(err)

}

func sealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	_ = rlp.Encode(hasher, []interface{}{
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
	})
	hasher.Sum(hash[:0])
	return hash
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeElementsToFile(file *os.File, arr []*big.Int) {
	_, err := fmt.Fprintf(file, "[")
	checkError(err)
	if len(arr) > 0 {
		_, err = fmt.Fprintf(file, "%s", arr[0].String())
		checkError(err)
	}

	for i := 1; i < len(arr); i++ {
		_, err = fmt.Fprintf(file, ", %s", arr[i].String())
		checkError(err)
	}
	_, err = fmt.Fprintf(file, "]")
	checkError(err)
}
