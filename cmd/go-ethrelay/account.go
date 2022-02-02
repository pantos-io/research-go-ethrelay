// This file contains logic executed if the command "account" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Prints the address of the current account",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(client.Account())
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
