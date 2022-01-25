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
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var deployFlagDstChain string

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys a smart contract (Ethash or ETH Relay) on the specified blockchain",
	Long:  `Deploys a smart contract (Ethash or ETH Relay) on the specified blockchain`,
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.PersistentFlags().StringVarP(&deployFlagDstChain, "destination", "d", "local", "A destination blockchain to which the smart contract should be deployed")
}

func updateChainsConfig(deployedAddress common.Address, chainId string, key string) {
	chainsConfig := viper.Get("chains.destinations").(map[string]interface{})
	deployChainConfig := chainsConfig[chainId].(map[string]interface{})
	deployChainConfig[key] = deployedAddress.String()

	chainsConfig[chainId] = deployChainConfig
	viper.Set("chains.destinations", chainsConfig)

	_ = viper.WriteConfig()
}
