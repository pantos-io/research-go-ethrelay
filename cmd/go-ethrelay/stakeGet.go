package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pantos-io/go-ethrelay/internal/ethereum/conversions"
)

// stakeRetrieveCmd represents the stake retrieve command
var stakeRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Aliases: []string{"get", "show"},
	Short: "Retrieves the stake stored on the specified destination chain",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		stakeInWei, err := client.GetStake(stakeFlagChain)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Stake balance: %f ETH\n", conversions.WeiToEther(stakeInWei))
	},
}

func init() {
	stakeCmd.AddCommand(stakeRetrieveCmd)
}
