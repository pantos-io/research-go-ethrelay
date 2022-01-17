// This file contains logic executed if the command "stake" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pantos-io/go-ethrelay/ethereum/utils"
)

var stakeFlagChain string

// stakeCmd represents the stake command
var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Shows the stake stored on the specified chain",
	Long: `Shows the stake stored on the specified chain`,
	Run: func(cmd *cobra.Command, args []string) {
		testimoniumClient = createTestimoniumClient()
		stakeInWei, err := testimoniumClient.GetStake(stakeFlagChain)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Stake balance: %s ETH\n", utils.WeiToEther(stakeInWei))
	},
}

func init() {
	rootCmd.AddCommand(stakeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	stakeCmd.PersistentFlags().StringVar(&stakeFlagChain, "chain", "local", "chain")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
