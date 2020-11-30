// This file contains logic executed if the command "get longestchainendpoint" is typed in.
// Authors: Leonhard Esterbauer

package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var testimoniumContractChain uint8

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
			log.Fatal("Failed to retrieve longest chain blockHash from chain " + strconv.Itoa(int(testimoniumContractChain)) + ":" + err.Error())
		}

		fmt.Printf("LongestChainEndpointBlockHash: %s\n", common.BytesToHash(blockHash[:]).String())
	},
}

func init() {
	getCmd.AddCommand(getLongestChainEndpointCmd)

	getLongestChainEndpointCmd.PersistentFlags().Uint8VarP(&testimoniumContractChain, "verifying", "v", 1, "The blockchain where the contract was deployed")
}
