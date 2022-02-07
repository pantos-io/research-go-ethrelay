package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pantos-io/go-ethrelay/internal/ethereum/encoding"
	"github.com/pantos-io/go-ethrelay/internal/ethereum/ethash"
	"github.com/pantos-io/go-ethrelay/internal/io"
	"github.com/pantos-io/go-ethrelay/pkg/ethrelay"
	"github.com/spf13/cobra"
)

var generateFlagChain string
var genesisBlockNumber = new(big.Int)

const BASE_DIR = "data"
var TX_DIR = filepath.Join(BASE_DIR, "transactions")
var RCP_DIR = filepath.Join(BASE_DIR, "receipts")
var POW_DIR = filepath.Join(BASE_DIR, "pows")

var generateCmd = &cobra.Command{
	Use: 	"generate [genesisBlock]",
	Short: 	"Generates and exports test data for the Ethrelay project",
	Long:	`Generates Merkle Proofs and PoWs for specific transactions / blocks and exports them in JSON format.
The data are intended to be used for the tests residing in the Ethrelay project.

If no genesis block is given, a recent block is chosen by the application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}

		if (len(args) == 0) {
			return nil
		}

		if _, success := genesisBlockNumber.SetString(args[0], 10); !success {
			return errors.New("failed to parse genesis block")
		}
		
		// TODO Read in latest block number
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) { run() },
}

func init() {
	rootCmd.AddCommand(generateCmd)

	addCommonFlag(generateCmd, "source", &generateFlagChain)
}

func run() {
	rand.Seed(time.Now().Unix())

	err := os.MkdirAll(BASE_DIR, fs.ModePerm)
	if err != nil {
		fmt.Println("Failed to create directory for test data:", err)
		return
	}

	writeEpoch()

	genesisBlock, err := client.BlockByNumber(generateFlagChain, genesisBlockNumber)
	blockErrorCheck(err)

	genesisPlus1Block, err := client.BlockByNumber(generateFlagChain, new(big.Int).Add(genesisBlockNumber, big.NewInt(1)))
	blockErrorCheck(err)

	genesisPlus6Block, err := client.BlockByNumber(generateFlagChain, new(big.Int).Add(genesisBlockNumber, big.NewInt(6)))
	blockErrorCheck(err)

	writeTransactionsAndReceipts(genesisBlock, genesisPlus1Block, genesisPlus6Block)

	genesisPlus2Block, err := client.BlockByNumber(generateFlagChain, new(big.Int).Add(genesisBlockNumber, big.NewInt(2)))
	blockErrorCheck(err)

	writePoWs(genesisPlus2Block)
}

func blockErrorCheck(err error) {
	if err != nil {
		log.Fatalln("Failed to fetch block:", err)
	}
}

func writeEpoch() {
	const EPOCH_FILE = "epoch.json"

	epoch := new(big.Int).Div(genesisBlockNumber, big.NewInt(30000)).Uint64()

	fmt.Printf("Fetching and writing epoch data for epoch %d...\n", epoch)
	epochData := ethash.GenerateEpochData(epoch)
	
	path := filepath.Join(BASE_DIR, EPOCH_FILE)
	if err := io.WriteToJson(path, epochData); err != nil {
		fmt.Println("Failed to write epoch data:", err)
	} else {
		fmt.Println("Wrote epoch data to", path)
	}
}

func writeTransactionsAndReceipts(genesis, genesisPlus1, genesisPlus6 *types.Block) {
	err := os.MkdirAll(TX_DIR, fs.ModePerm)
	if err != nil {
		fmt.Println("Failed to create directory for transactions:", err)
		return
	}

	err = os.MkdirAll(RCP_DIR, fs.ModePerm)
	if err != nil {
		fmt.Println("Failed to create directory for receipts:", err)
		return
	}

	writeTransactionAndReceipt(genesis, 		"genesis")
	writeTransactionAndReceipt(genesisPlus1, 	"genesisPlus1")
	writeTransactionAndReceipt(genesisPlus6, 	"genesisPlus6")
}

func writeTransactionAndReceipt(block *types.Block, fileName string) {
	txHash, proof, err := getRandomTransactionAndReceiptProof(block)
	if (err != nil) {
		fmt.Printf("Failed to generate Merkle Proof for block %s: %s", block.Number(), err)
		return
	}

	txProof 	:= proof[0]
	rcpProof 	:= proof[1]

	txPath		:= filepath.Join(TX_DIR, fmt.Sprint(fileName, ".json"))
	rcpPath		:= filepath.Join(RCP_DIR, fmt.Sprint(fileName, ".json"))
	

	err = io.WriteToJson(txPath, txProof)
	if err == nil {
		fmt.Println("Wrote Merkle Proof for transaction", txHash, "to", txPath)
	} else {
		fmt.Printf("Failed to write Merkle Proof for transaction %s to file %s: %s\n", txHash, txPath, err)
	}

	err = io.WriteToJson(rcpPath, rcpProof)
	if err == nil {
		fmt.Println("Wrote Merkle Proof for receipt of transaction", txHash, "to", rcpPath)
	} else {
		fmt.Printf("Failed to write Merkle Proof for receipt of transaction %s to file %s: %s\n", txHash, rcpPath, err)
	}
}

func getRandomTransactionAndReceiptProof(block *types.Block) (common.Hash, [2]*ethrelay.MerkleProof, error) {
	txHash := getRandomTxHash(block)

	_, txProof, err := client.GenerateMerkleProofForTx(generateFlagChain, txHash)
	if err != nil {
		fmt.Println()
		return common.Hash{}, [2]*ethrelay.MerkleProof{}, fmt.Errorf("failed to generate Merkle Proof for transaction %s: %s", txHash, err)
	}
	_, rcpProof, err := client.GenerateMerkleProofForTx(generateFlagChain, txHash)
	if err != nil {
		fmt.Println()
		return common.Hash{}, [2]*ethrelay.MerkleProof{}, fmt.Errorf("failed to generate Merkle Proof for receipt of transaction %s: %s", txHash, err)
	}

	return txHash, [2]*ethrelay.MerkleProof{txProof, rcpProof}, nil
}

func getRandomTxHash(block *types.Block) common.Hash {
	i := rand.Intn(block.Transactions().Len())
	return block.Transactions()[i].Hash()
}

type powJson struct {
	DatasetLookUp []*big.Int;
	WitnessForLookup []*big.Int;
}

func writePoWs(genesisPlus2 *types.Block) {
	err := os.MkdirAll(POW_DIR, fs.ModePerm)
	if err != nil {
		fmt.Println("Failed to create directory for PoWs:", err)
		return
	}

	writePoW(genesisPlus2, "genesisPlus2")
}

func writePoW(block *types.Block, fileName string) {
	blockHeader := block.Header()
	rlpHeader, err := encoding.EncodeHeaderWithoutNonceToRLP(blockHeader)
	if (err != nil) {
		fmt.Println("Failed to encode header of block", block.Number(), "to RLP:", err)
		return
	}

	blockMetaData := ethash.NewBlockMetaData(blockHeader.Number.Uint64(), blockHeader.Nonce.Uint64(), crypto.Keccak256Hash(rlpHeader))

	fmt.Printf("Generating PoW for block %s...\n", block.Number())
	
	path := filepath.Join(POW_DIR, fmt.Sprint(fileName, ".json"))
	err = io.WriteToJson(path, powJson{blockMetaData.DAGElementArray(), blockMetaData.DAGProofArray()})
	if err == nil {
		fmt.Println("Wrote PoW for block", block.Number(), "to", path)
	} else {
		fmt.Printf("Failed to write PoW for block %s to file %s: %s\n", block.Number(), path, err)
	}
}