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

package main

import (
	"github.com/spf13/cobra"
)

// ethashCmd represents the ethash command
var ethashCmd = &cobra.Command{
	Use:   "ethash",
	Short: "Deploys the Ethash smart contract on the specified blockchain",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		deployedAddress := client.DeployEthash(deployFlagDstChain)

		updateChainsConfig(deployedAddress, deployFlagDstChain, "ethashAddress")
	},
}

func init() {
	deployCmd.AddCommand(ethashCmd)
}
