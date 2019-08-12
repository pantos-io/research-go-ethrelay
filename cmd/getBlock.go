// This file contains logic executed if the command "get block" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"log"
)

var headerFlag bool

// verifyBlockCmd represents the block command
var getBlockCmd = &cobra.Command{
	Use:   "block [blockHash]",
	Short: "Retrieves a block",
	Long: `Retrieves the block with the specified hash ('blockHash')`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockHash := common.HexToHash(args[0])

		testimoniumClient = createTestimoniumClient()

		if headerFlag {
			header, err := testimoniumClient.HeaderByHash(blockHash, getFlagChain)
			if err != nil {
				log.Fatal("Failed to retrieve header: " + err.Error())
			}
			printHeader(header)
			return
		}

		block, err := testimoniumClient.BlockByHash(blockHash, getFlagChain)
		if err != nil {
			log.Fatal("Failed to retrieve block: " + err.Error())
		}
		printBlock(block)
		if detailFlag {
			printTransactions(block)
		}
	},
}

func init() {
	getCmd.AddCommand(getBlockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getBlockCmd.Flags().BoolVar(&headerFlag, "header", false, "Get the header of the block")
	getBlockCmd.Flags().BoolVarP(&detailFlag, "detail", "d", false, "Show transaction details of block")
}

func printHeader(header *types.Header) {
	fmt.Printf("Hash: %s\n", header.Hash().String())
	fmt.Printf("Number: %s\n", header.Number.String())
	fmt.Printf("Nonce: %d\n", header.Nonce.Uint64())
	fmt.Printf("StateRoot: %s\n", header.Root.String())
	fmt.Printf("TxHash: %s\n", header.TxHash.String())
	fmt.Printf("ReceiptHash: %s\n", header.ReceiptHash.String())
}

func printBlock(block *types.Block) {
	fmt.Printf("Hash: %s\n", block.Hash().String())
	fmt.Printf("Number: %d\n", block.Number())
	fmt.Printf("Nonce: %d\n", block.Nonce())
	fmt.Printf("Transaction Count: %d\n", len(block.Transactions()))
}

func printTransactions(block *types.Block) {
	fmt.Printf("Transactions:\n")
	for index, tx := range(block.Transactions()) {
		fmt.Printf("%d: %s\n", index, tx.Hash().String())
	}
}
