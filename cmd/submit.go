// This file contains logic executed if the command "submit" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"github.com/spf13/cobra"
)

var submitFlagDestChain string

// verifyCmd represents the verify command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submits a block of the target chain or epoch data to the verifying chain",
	Long: `Submits a block of the target chain or epoch data to the verifying chain`,
}

func init() {
	rootCmd.AddCommand(submitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	submitCmd.PersistentFlags().StringVar(&submitFlagDestChain, "destination", "local", "verifying chain")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
