// This file contains logic executed if the command "stake withdraw" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/spf13/cobra"

	"github.com/pantos-io/go-ethrelay/pkg/ethereum/utils"
)

// stakeDepositCmd represents the command 'stake deposit <amount>'
var stakeWithdrawCmd = &cobra.Command{
	Use:   "withdraw amountInWei",
	Short: "Withdraws the specified amount of Wei.",
	Long:  `Withdraws the specified amount of Wei, i.e., the client's stake is decreased by the specified amount`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		amountInWei := new(big.Int)
		amountInWei, ok := amountInWei.SetString(args[0], 10)
		if !ok {
			log.Fatal("Can not parse amountInWei parameter")
		}

		err := client.WithdrawStake(stakeFlagChain, amountInWei)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Successfully withdrew stake: %f ETH\n", utils.WeiToEther(amountInWei))
	},
}

func init() {
	stakeCmd.AddCommand(stakeWithdrawCmd)
}
