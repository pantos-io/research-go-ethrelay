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
	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

var noOfConfirmations uint8

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction [txHash]",
	Short: "Verify a transaction",
	Long: `Verify a transaction from the source chain on the destination chain

Behind the scene, the command queries the transaction with the specified hash ('txHash') from the source chain.
It then generates a Merkle Proof contesting the existence of the transaction within a specific block.
This information gets sent to the destination chain, where not only the existence of the block but also the Merkle Proof are verified`,
	Run: func(cmd *cobra.Command, args []string) {
		txHash := BytesToBytes32(common.Hex2Bytes(args[0]))
		testimoniumClient.VerifyTransaction(txHash, noOfConfirmations, verifyFlagDestChain)
	},
}

func init() {
	verifyCmd.AddCommand(transactionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	transactionCmd.Flags().Uint8VarP(&noOfConfirmations, "confirmations", "c", 4, "Number of block confirmations (default: 4)")
}
