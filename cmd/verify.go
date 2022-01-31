// This file contains logic executed if the command "verify" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"github.com/pantos-io/go-ethrelay/ethrelay"
	"github.com/spf13/cobra"
)

var verifyFlagSrcChain string
var verifyFlagDstChain string

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies a transaction or a block from a source chain on a destination chain",
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	verifyCmd.PersistentFlags().StringVarP(&verifyFlagSrcChain, "source", "s", "mainnet", "a source chain from which to read data for verification")
	verifyCmd.RegisterFlagCompletionFunc("source", chainCompletionFn(ethrelay.ChainTypeSrc))
	verifyCmd.PersistentFlags().StringVarP(&verifyFlagDstChain, "destination", "d", "local", "a destination chain whose data needs to be verified")
	verifyCmd.RegisterFlagCompletionFunc("destination", chainCompletionFn(ethrelay.ChainTypeDst))
}
