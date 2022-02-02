// This file contains logic executed if the command "verify" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
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

	addCommonFlag(verifyCmd, "source", &verifyFlagSrcChain)
	addCommonFlag(verifyCmd, "destination", &verifyFlagDstChain)
}
