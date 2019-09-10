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
	"github.com/pantos-io/go-testimonium/testimonium"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the Testimonium client",
	Long: `This command initializes the Testimonium client. 
This command sets up the testimonium.yml file in the current directory.
The file contains connection configurations for the different blockchains, e.g.,
private key, url, port, etc.

The default testimonium.yml file looks like this:

    privateKey: <YOUR PRIVATE KEY>
    chains:
        0:
            url: mainnet.infura.io
        1:
            type: http
            url: localhost
            port: 7545

Chain ID 0 contains connection configuration for the main Ethereum chain (via Infura).
Chain ID 1 contains connection configuration for a local chain (e.g., run via Ganache).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setting up testimonium.yml...")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the private key of your account (the account will be used on all chains): ")
		privateKey, _ := reader.ReadString('\n')

		viper.Set("privateKey", privateKey[:len(privateKey)-1])

		chainsConfig := make(map[uint8]interface{})

		mainnetConfig := testimonium.CreateChainConfig("", "mainnet.infura.io", 0)
		chainsConfig[0] = mainnetConfig

		ganacheConfig := testimonium.CreateChainConfig("http", "localhost", 8545)
		chainsConfig[1] = ganacheConfig

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
				if response == "Y\n" {
					fmt.Println("Overwriting...")
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
