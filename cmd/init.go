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

	"github.com/pantos-io/go-ethrelay/testimonium"
	"github.com/spf13/viper"

	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the ETH Relay client",
	Long: `This command initializes the ETH Relay client. 
This command sets up the testimonium.yml file in the current directory.
The file contains connection configurations for the different blockchains, e.g.,
private key, URL, port, etc.

The default testimonium.yml file looks like this:

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
	privateKey: <YOUR PRIVATE KEY>

Websocket-Connection is required for submitting blocks in live mode.
Chains under "sources" contain connection configurations for the source chains, defaulting to the main Ethereum chain (via Infura).
Chains under "destinations" contain connection configurations for the verifying chains, defaulting to a local chain (e.g. via Ganache).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setting up testimonium.yml...")
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

		mainnetConfig := testimonium.CreateChainConfig("wss", "mainnet.infura.io/ws/v3/1e835672adba4b9b930a12a3ec58ebad", 0)
		sources["mainnet"] = mainnetConfig

		localConfig := testimonium.CreateChainConfig("http", "localhost", 7545)
		destinations["local"] = localConfig

		chainsConfig["sources"] = sources
		chainsConfig["destinations"] = destinations

		viper.Set("chains", chainsConfig)

		err := viper.SafeWriteConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				err := ioutil.WriteFile("./testimonium.yml", []byte(""), 0644)
				if err != nil {
					fmt.Printf("Unable to write file: %v", err)
					return
				}
			} else {
				// File already exists
				fmt.Print("File testimonium.yml already exists. Overwrite? (n/Y):")
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
		fmt.Println("Created testimonium.yml.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
