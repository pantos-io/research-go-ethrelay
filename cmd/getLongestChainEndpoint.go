// This file contains logic executed if the command "get longestchainendpoint" is typed in.
// Authors: Leonhard Esterbauer

package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var testimoniumContractChain string

// getLongestChainEndpointCmd represents the block command
var getLongestChainEndpointCmd = &cobra.Command{
	Use:   "longestchainendpoint",
	Short: "Retrieves the blockhash of the longest chain on a verifying endpoint",
	Long: `Retrieves the blockhash of the longest chain on a verifying endpoint`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		testimoniumClient = createTestimoniumClient()

		blockHash, err := testimoniumClient.GetLongestChainEndpoint(testimoniumContractChain)
		if err != nil {
			log.Fatalf("Failed to retrieve longest chain blockHash from chain '%s': %s", testimoniumContractChain, err)
		}

		fmt.Println("LongestChainEndpointBlockHash: ", common.BytesToHash(blockHash[:]).String())
	},
}

func init() {
	getCmd.AddCommand(getLongestChainEndpointCmd)

	getLongestChainEndpointCmd.PersistentFlags().StringVarP(&testimoniumContractChain, "verifying", "v", "local", "The blockchain where the contract was deployed")
}
