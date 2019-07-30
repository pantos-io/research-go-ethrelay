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
	"github.com/pf92/testimonium-cli/ethereum"

	"log"

	"github.com/spf13/cobra"
)

var disputeFlagChain uint8

// disputeCmd represents the dispute command
var disputeCmd = &cobra.Command{
	Use:   "dispute [blockHash]",
	Short: "Disputes a submitted block header",
	Long: `Disputes the submitted block header with the specified hash ('blockHash')`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockHash := common.HexToHash(args[0])	// omit the first two chars "0x"
		blockHashBytes := blockHash.Bytes()
		var blockHashBytes32 [32]byte
		copy(blockHashBytes32[:], blockHashBytes)

		// get blockNumber, nonce and RlpHeaderHashWithoutNonce and generate dataSetLookup and witnessForLookup
		header, err := testimoniumClient.BlockHeader(blockHashBytes32, disputeFlagChain)
		if err != nil {
			log.Fatal("Failed to retrieve header from contract: " + err.Error())
		}

		fmt.Println("create DAG, compute dataSetLookup and witnessForLookup")
		// get DAG and compute dataSetLookup and witnessForLookup
		blockMetaData := ethereum.NewBlockMetaData(&header)
		dataSetLookup := blockMetaData.DAGElementArray()
		witnessForLookup := blockMetaData.DAGProofArray()
		//time.Sleep(20 * time.Second)
		//var dataSetLookup []*big.Int
		//var witnessForLookup []*big.Int
		//fmt.Println("call disputeBlock(...)")
		fmt.Println(disputeFlagChain)
		testimoniumClient.DisputeBlock(blockHash, dataSetLookup, witnessForLookup, disputeFlagChain)
		//fmt.Println(header)
		//fmt.Println(dataSetLookup)
		//fmt.Println(witnessForLookup)

		//fmt.Println(blockHash)
	},
}

func init() {
	rootCmd.AddCommand(disputeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disputeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disputeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	disputeCmd.Flags().Uint8VarP(&disputeFlagChain, "chain", "c", 1, "the disputed chain ID")
}
