// This file contains logic executed if the command "account" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Prints the address of the current account",
	Long: `Prints the address of the current account`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(createEthrelayClient().Account())
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
