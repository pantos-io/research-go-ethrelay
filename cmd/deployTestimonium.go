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

var deployFlagTargetChain uint8
var deployFlagGenesisNumber uint64

// testimoniumCmd represents the testimonium command
var testimoniumCmd = &cobra.Command{
	Use:   "testimonium",
	Short: "Deploys the Testimonium smart contract on the specified blockchain",
	Long: `Deploys the Testimonium smart contract on the specified blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		testimoniumClient = createTestimoniumClient()
		deployedAddress := testimoniumClient.DeployTestimonium(deployFlagChain, deployFlagTargetChain, deployFlagGenesisNumber)

		updateChainsConfig(deployedAddress, deployFlagChain, "testimoniumAddress")
	},
}

func init() {
	deployCmd.AddCommand(testimoniumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testimoniumCmd.PersistentFlags().String("foo", "", "A help for foo")
	testimoniumCmd.Flags().Uint8VarP(&deployFlagTargetChain, "target", "t", 0, "The 'target' chain containing the specified genesis block")
	testimoniumCmd.Flags().Uint64VarP(&deployFlagGenesisNumber, "genesis", "g", 1, "The number of the block (of the target chain) that should be used as genesis block")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testimoniumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
