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
	"github.com/spf13/cobra"
	"log"
	"math/big"
)

var submitFlagSrcChain uint8
var submitFlagDestChain uint8
var submitFlagRandomize bool

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit [blocknumber]",
	Short: "Submits a block header to the destination chain",
	Long: `Submits the specified block header from the source chain to the destination chain`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var blockNumber *big.Int = nil
		if len(args) > 0 {
			var ok bool
			blockNumber = new(big.Int)
			blockNumber, ok = blockNumber.SetString(args[0], 10)
			if !ok {
				log.Fatalf("Illegal block number '%s'", args[0])
			}
		}
		if submitFlagRandomize {
			randomizedRlpHeader, err := testimoniumClient.RandomizeHeader(blockNumber, submitFlagSrcChain)
			if err != nil {
				log.Fatal("Failed to randomize block: " + err.Error())
			}
			testimoniumClient.SubmitRLPHeader(randomizedRlpHeader, submitFlagDestChain)
			return
		}
		testimoniumClient.SubmitHeader(blockNumber, submitFlagSrcChain, submitFlagDestChain)
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	submitCmd.Flags().BoolP("live", "l", false, "live mode (continuously submits most recent block headers)")
	submitCmd.Flags().Uint8Var(&submitFlagSrcChain, "source", 2, "source chain (default: 2)")
	submitCmd.Flags().Uint8Var(&submitFlagDestChain, "destination", 1, "destination chain (default: 1)")
	submitCmd.Flags().BoolVar(&submitFlagRandomize, "randomize", false, "randomize block (default: false)")
}
