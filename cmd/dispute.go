// This file contains logic executed if the command "dispute" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var disputeFlagChain, sourceFlagChain uint8

// disputeCmd represents the dispute command
var disputeCmd = &cobra.Command{
	Use:   "dispute [blockHash]",
	Short: "Disputes a submitted block header",
	Long:  `Disputes the submitted block header with the specified hash ('blockHash')`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockHash := common.HexToHash(args[0])

		// copy dynamic byte array to fixed length byte array
		var blockHashBytes32 [32]byte
		blockHashBytes := blockHash.Bytes()
		copy(blockHashBytes32[:], blockHashBytes)

		// call disputeBlock in the testimonium client library
		testimoniumClient = createTestimoniumClient()
		testimoniumClient.DisputeBlock(blockHash, disputeFlagChain, sourceFlagChain)
	},
}

func init() {
	rootCmd.AddCommand(disputeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disputeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disputeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	disputeCmd.Flags().Uint8VarP(&disputeFlagChain, "disputed", "d", 1, "the disputed chain ID")
	disputeCmd.Flags().Uint8VarP(&sourceFlagChain, "source", "s", 0, "the source chain ID")
}
