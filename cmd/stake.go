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
	Run: func(cmd *cobra.Command, args []string) {
		ethrelayClient = createEthrelayClient()
		stakeInWei, err := ethrelayClient.GetStake(stakeFlagChain)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Stake balance: %f ETH\n", utils.WeiToEther(stakeInWei))
	},
}

func init() {
	rootCmd.AddCommand(stakeCmd)

	stakeCmd.PersistentFlags().StringVar(&stakeFlagChain, "chain", "local", "chain")
}
