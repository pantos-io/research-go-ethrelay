// This file contains logic executed if the command "get transaction" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

var receiptFlag bool

// getTransactionCmd represents the transaction command
var getTransactionCmd = &cobra.Command{
	Use:   "transaction txHash",
	Short: "Retrieves a transaction",
	Long: `Retrieves the transaction with the specified hash`,
	Aliases: []string{"tx"},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])

		ethrelayClient = createEthrelayClient()

		if receiptFlag {
			txReceipt, err := ethrelayClient.TransactionReceipt(getFlagChain, txHash)
			if err != nil {
				log.Fatal("Failed to retrieve transaction receipt: " + err.Error())
			}
			printTransactionReceipt(txReceipt)
			return
		}

		tx, _, err := ethrelayClient.Transaction(getFlagChain, txHash)
		if err != nil {
			log.Fatal("Failed to retrieve transaction: " + err.Error())
		}
		printTransaction(tx)
	},
}

func init() {
	getCmd.AddCommand(getTransactionCmd)

	getTransactionCmd.Flags().BoolVarP(&receiptFlag, "receipt", "r", false, "Get the receipt of the transaction")
}

func printTransaction(tx *types.Transaction) {
	fmt.Printf("Hash: %s\n", tx.Hash())
	fmt.Printf("To: %s\n", tx.To())
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("Value: %d\n", tx.Value())
	fmt.Printf("GasPrice: %d\n", tx.GasPrice())
	fmt.Printf("Gas: %d\n", tx.Gas())
}

func printTransactionReceipt(receipt *types.Receipt) {
	fmt.Printf("TxHash: %s\n", receipt.TxHash)
	fmt.Printf("BlockHash: %s\n", receipt.BlockHash)
	fmt.Printf("Status: %d\n", receipt.Status)
	fmt.Printf("BlockNumber: %d\n", receipt.BlockNumber)
	fmt.Printf("GasUsed: %d\n", receipt.GasUsed)
	fmt.Printf("CumulativeGasUsed: %d\n", receipt.CumulativeGasUsed)
	fmt.Printf("TransactionIndex: %d\n", receipt.TransactionIndex)
	fmt.Printf("ContractAddress: %s\n", receipt.ContractAddress)
}
