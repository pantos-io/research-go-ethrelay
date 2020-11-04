// This file contains logic executed if the command "get longestchainendpoint" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// getLongestChainEndpointCmd represents the block command
var getLongestChainEndpointCmd = &cobra.Command{
	Use:   "longestchainendpoint",
	Short: "Retrieves the blockhash of the longest chain on a verifying endpoint",
	Long: `Retrieves the blockhash of the longest chain on a verifying endpoint`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		testimoniumClient = createTestimoniumClient()

		blockHash, err := testimoniumClient.LongestChainEndpoint(getFlagChain)
		if err != nil {
			log.Fatal("Failed to retrieve longest chain blockHash from chain " + strconv.Itoa(int(getFlagChain)) + ":" + err.Error())
		}

		fmt.Printf("LongestChainEndpointBlockHash: { Hash: %s }\n", common.BytesToHash(blockHash[:]).String())
	},
}

func init() {
	getCmd.AddCommand(getLongestChainEndpointCmd)
}
