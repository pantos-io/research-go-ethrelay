// This file contains logic executed if the command "dispute" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var disputeFlagChain string

// disputeCmd represents the dispute command
var disputeCmd = &cobra.Command{
	Use:   "dispute blockHash",
	Short: "Disputes a submitted block header",
	Long: `Disputes the submitted block header with the specified hash`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockHash := common.HexToHash(args[0])

		client.DisputeBlock(disputeFlagChain, blockHash)
	},
}

func init() {
	rootCmd.AddCommand(disputeCmd)

	addCommonFlag(disputeCmd, "destination", &disputeFlagChain)
}
