// This file contains logic executed if the command "submit" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"github.com/pantos-io/go-ethrelay/ethrelay"
	"github.com/spf13/cobra"
)

var submitFlagDstChain string

// verifyCmd represents the verify command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submits a block or epoch data of a source chain to a destination chain",
}

func init() {
	rootCmd.AddCommand(submitCmd)

	submitCmd.PersistentFlags().StringVarP(&submitFlagDstChain, "destination", "d", "local", "a destination chain to which to submit data to")
	submitCmd.RegisterFlagCompletionFunc("destination", chainCompletionFn(ethrelay.ChainTypeDst))
}
