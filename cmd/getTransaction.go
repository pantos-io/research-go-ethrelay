// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"log"
)

// getTransactionCmd represents the transaction command
var getTransactionCmd = &cobra.Command{
	Use:   "transaction [txHash]",
	Short: "Retrieves a transaction",
	Long: `Retrieves the transaction with the specified hash`,
	Aliases: []string{"tx"},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])
		tx, _, err := testimoniumClient.Transaction(txHash, getFlagChain)
		if err != nil {
			log.Fatal("Failed to retrieve transaction: " + err.Error())
		}
		printTransaction(tx)
	},
}

func init() {
	getCmd.AddCommand(getTransactionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTransactionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printTransaction(tx *types.Transaction) {
	fmt.Printf("Hash: %s\n", tx.Hash().String())
	fmt.Printf("To: %s\n", tx.To().String())
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("Value: %d\n", tx.Value())
	fmt.Printf("GasPrice: %d\n", tx.GasPrice())
	fmt.Printf("Gas: %d\n", tx.Gas())
}
