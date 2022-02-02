// This file contains logic executed if the command "get longestchainendpoint" is typed in.
// Authors: Leonhard Esterbauer

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var ethrelayContractChain string

// getLongestChainEndpointCmd represents the block command
var getLongestChainEndpointCmd = &cobra.Command{
	Use:   "longestchainendpoint",
	Short: "Retrieves the blockhash of the longest chain on a destination blockchain",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		blockHash, err := client.GetLongestChainEndpoint(ethrelayContractChain)
		if err != nil {
			log.Fatalf("Failed to retrieve longest chain blockHash from chain '%s': %s", ethrelayContractChain, err)
		}

		fmt.Println("LongestChainEndpointBlockHash:", common.BytesToHash(blockHash[:]))
	},
}

func init() {
	getCmd.AddCommand(getLongestChainEndpointCmd)

	addCommonFlag(getLongestChainEndpointCmd, "destination", &ethrelayContractChain)
}
