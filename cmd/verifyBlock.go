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
// This file contains logic executed if the command "verify block" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// verifyBlockCmd represents the block command
// TODO: this command only compares the hashes and checks for existence on the respective chain, not for equality
//  even though a tampering is hard to achieve, this does not mean the blocks are equal
// 	another unlikely misbehaviour of this command is occurs if one tries to verify a block to kick off a dispute
// 	- here it is possible that a block is valid on the destination chain and can't be disputed by participants,
//  however, this would means that an "attacker" can produce blocks in a faster way than the source blockchain
//  and is for that reason very unlikely
var verifyBlockCmd = &cobra.Command{
	Use:   "block [blockHash]",
	Short: "Verifies a block",
	Long:  `Gets sure a block with [blockHash] from a source blockchain is also present on a destination blockchain`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockHash := common.HexToHash(args[0])

		ethrelayClient = createEthrelayClient()

		headerExists, err := ethrelayClient.BlockHeaderExists(verifyFlagDstChain, blockHash)
		if err != nil {
			log.Fatal("Could not verify block header on destination chain: " + err.Error())
		}

		if !headerExists {
			fmt.Printf("No header stored for block %s on destination chain\n", ShortHexString(args[0]))
			return
		}

		_, err = ethrelayClient.GetOriginalBlockHeader(verifyFlagSrcChain, blockHash)
		if err != nil {
			log.Fatal("Could not get original block on source chain: " + err.Error())
		}

		fmt.Printf("Block %s is valid\n", ShortHexString(args[0]))
	},
}

func init() {
	verifyCmd.AddCommand(verifyBlockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyBlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyBlockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ShortHexString(hex string) string {
	if len(hex) <= 12 {
		return hex
	}

	return fmt.Sprintf("%s...%s", hex[:6], hex[len(hex)-4:])
}
