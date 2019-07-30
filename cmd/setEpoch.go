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
	"github.com/pf92/testimonium-cli/ethereum/ethash"
	"math/big"

	"log"

	"github.com/spf13/cobra"
)

var flagChain uint8

// setEpochCmd represents the command for setting epoch data (Ethash contract)
var setEpochCmd = &cobra.Command{
	Use:   "set-epoch [epoch]",
	Short: "Sets the epoch data for the specified epoch",
	Long: `Sets the epoch data for the specified epoch`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var epoch *big.Int = nil
		var ok bool
		epoch = new(big.Int)
		epoch, ok = epoch.SetString(args[0], 10)
		if !ok {
			log.Fatalf("Illegal epoch number '%s'", args[0])
		}

		epochData := ethash.GenerateEpochData(epoch.Uint64())
		testimoniumClient.SetEpochData(epochData, flagChain)
	},
}

func init() {
	rootCmd.AddCommand(setEpochCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disputeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disputeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setEpochCmd.Flags().Uint8VarP(&flagChain, "chain", "c", 1, "the chain ID")
}
