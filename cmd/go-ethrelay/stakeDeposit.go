// This file contains logic executed if the command "stake deposit" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/spf13/cobra"

	"github.com/pantos-io/go-ethrelay/internal/ethereum/conversions"
)

// stakeDepositCmd represents the command 'stake deposit <amount>'
var stakeDepositCmd = &cobra.Command{
	Use:   "deposit amountInWei",
	Short: "Deposits the specified amount of Wei.",
	Long:  "Deposits the specified amount of Wei, i.e., the client's stake is increased by the specified amount",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		amountInWei := new(big.Int)
		amountInWei, ok := amountInWei.SetString(args[0], 10)
		if !ok {
			log.Fatal("Can not parse amountInWei parameter")
		}

		err := client.DepositStake(stakeFlagChain, amountInWei)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Successfully deposited stake: %f ETH\n", conversions.WeiToEther(amountInWei))
	},
}

func init() {
	stakeCmd.AddCommand(stakeDepositCmd)
}
