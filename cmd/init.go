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
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/pantos-io/go-ethrelay/ethrelay"
	"github.com/spf13/viper"

	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the ETH Relay client",
	Long:  `This command initializes the ETH Relay client.
This command sets up the ethrelay.yml file in the current directory.
The file contains connection configurations for the different blockchains, e.g.,
private key, URL, port, etc.

The default ethrelay.yml file looks like this:

chains:
  sources:
    mainnet:
      type: wss
      url: mainnet.infura.io/ws/v3/1e835672adba4b9b930a12a3ec58ebad
  destinations:
    local:
      port: 7545
      type: http
      url: localhost
privatekey: 0x0

Websocket-Connection is required for submitting blocks in live mode.
Chains under "sources" contain connection configurations for the source chains, defaulting to the main Ethereum chain (via Infura).
Chains under "destinations" contain connection configurations for the destination chains, defaulting to a local chain (e.g. via Ganache).`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setting up ethrelay.yml...")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the private key of your account (the account will be used on all chains, input this in the format starting with '0x...'): ")
		privateKey, _ := reader.ReadString('\n')

		if !strings.HasPrefix(privateKey, "0x") {
			fmt.Println("Entered private key is not starting with '0x'.")
			return
		}

		viper.Set("privateKey", privateKey[:len(privateKey)-1])

		chainsConfig := make(map[string]interface{})
		sources := make(map[string]interface{})
		destinations := make(map[string]interface{})

		mainnetConfig := ethrelay.CreateChainConfig("wss", "mainnet.infura.io/ws/v3/1e835672adba4b9b930a12a3ec58ebad", 0)
		sources["mainnet"] = mainnetConfig

		localConfig := ethrelay.CreateChainConfig("http", "localhost", 7545)
		destinations["local"] = localConfig

		chainsConfig["sources"] = sources
		chainsConfig["destinations"] = destinations

		viper.Set("chains", chainsConfig)

		err := viper.SafeWriteConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				err := ioutil.WriteFile("./ethrelay.yml", []byte(""), 0644)
				if err != nil {
					fmt.Printf("Unable to write file: %v", err)
					return
				}
			} else {
				fmt.Print("File ethrelay.yml already exists. Overwrite? (n/Y):")
				response, _ := reader.ReadString('\n')
				overwrite, _ := regexp.MatchString("^[yY]?\n$", response)
				if overwrite {
					fmt.Println("Overwriting...")
				} else {
					return
				}
			}
			_ = viper.WriteConfig()
		}
		fmt.Println("Created ethrelay.yml.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
