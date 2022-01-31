// This file contains logic executed if the command "get" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"github.com/pantos-io/go-ethrelay/ethrelay"
	"github.com/spf13/cobra"
)

var getFlagChain string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a block or transaction",
	Long: "Retrieves a block or transaction from one of the chains.",
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVar(&getFlagChain, "chain", "mainnet", "chain")
	getCmd.RegisterFlagCompletionFunc("chain", chainCompletionFn(ethrelay.ChainTypeAny))
}
