// This file contains logic executed if the command "dispute" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

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

		// copy dynamic byte array to fixed length byte array
		var blockHashBytes32 common.Hash
		blockHashBytes := blockHash.Bytes()
		copy(blockHashBytes32[:], blockHashBytes)

		client.DisputeBlock(disputeFlagChain, blockHash)
	},
}

func init() {
	rootCmd.AddCommand(disputeCmd)

	disputeCmd.Flags().StringVarP(&disputeFlagChain, "chain", "c", "local", "the disputed chain ID")
}
