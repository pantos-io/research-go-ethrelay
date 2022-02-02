// This file contains logic executed if the command "stake" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
	"github.com/spf13/cobra"
)

var stakeFlagChain string

// stakeCmd represents the stake command
var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Retrieves, deposits or withdraws stake from a destination chain",
}

func init() {
	rootCmd.AddCommand(stakeCmd)

	addCommonFlag(stakeCmd, "destination", &stakeFlagChain)
}
