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
	"github.com/pf92/testimonium-cli/testimonium"
	"github.com/spf13/cobra"
	"log"
)

// verifyReceiptCmd represents the receipt command
var verifyReceiptCmd = &cobra.Command{
	Use:   "receipt [txHash]",
	Short: "Verify a receipt",
	Long: `Verify a receipt from the source chain on the destination chain

Behind the scene, the command queries the receipt with the specified hash ('txHash') from the source chain.
It then generates a Merkle Proof contesting the existence of the receipt within a specific block.
This information gets sent to the destination chain, where not only the existence of the block but also the Merkle Proof are verified`,
	Aliases: []string{"tx"},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])
		blockHash, rlpEncodedReceipt, path, rlpEncodedProofNodes, err := testimoniumClient.GenerateMerkleProofForReceipt(txHash, verifyFlagSrcChain)
		if err != nil {
			log.Fatal("Failed to generate Merkle Proof: " + err.Error())
		}

		isValid := testimoniumClient.VerifyMerkleProof(blockHash, testimonium.VALUE_TYPE_RECEIPT, rlpEncodedReceipt, path, rlpEncodedProofNodes, noOfConfirmations, verifyFlagDestChain)
		fmt.Println("Receipt Validation Result: ", isValid)
	},
}

func init() {
	verifyCmd.AddCommand(verifyReceiptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	verifyReceiptCmd.Flags().Uint8VarP(&noOfConfirmations, "confirmations", "c", 4, "Number of block confirmations")
}
