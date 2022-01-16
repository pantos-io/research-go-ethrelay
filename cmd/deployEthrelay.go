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
)

var deployFlagSourceChain string
var deployFlagGenesisNumber uint64

// ethrelayCmd represents the ethrelay command
var ethrelayCmd = &cobra.Command{
	Use:   "ethrelay",
	Short: "Deploys the ETH Relay smart contract on the specified blockchain",
	Long:  `Deploys the ETH Relay smart contract on the specified blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		testimoniumClient = createTestimoniumClient()
		deployedAddress := testimoniumClient.DeployTestimonium(deployFlagTargetChain, deployFlagSourceChain, deployFlagGenesisNumber)

		updateChainsConfig(deployedAddress, deployFlagTargetChain, "ethrelayAddress")
	},
}

func init() {
	deployCmd.AddCommand(ethrelayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethrelayCmd.PersistentFlags().String("foo", "", "A help for foo")
	ethrelayCmd.Flags().StringVarP(&deployFlagSourceChain, "source", "s", "mainnet", "The source chain containing the specified genesis block")
	ethrelayCmd.Flags().Uint64VarP(&deployFlagGenesisNumber, "genesis", "g", 1, "The number of the block (of the source chain) that should be used as genesis block")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethrelayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
