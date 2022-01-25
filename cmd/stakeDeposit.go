// This file contains logic executed if the command "stake deposit" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"
	"math/big"

	"github.com/spf13/cobra"

	"github.com/pantos-io/go-ethrelay/ethereum/utils"
)

// stakeDepositCmd represents the command 'stake deposit <amount>'
var stakeDepositCmd = &cobra.Command{
	Use:   "deposit amountInWei",
	Short: "Deposits the specified amount of Wei.",
	Long: `Deposits the specified amount of Wei, i.e., the client's stake is increased by the specified amount'`,
	Run: func(cmd *cobra.Command, args []string) {
		ethrelayClient = createEthrelayClient()

		amountInWei := new(big.Int)
		amountInWei, ok := amountInWei.SetString(args[0], 10)
		if !ok {
			log.Fatal("Can not parse amountInWei parameter")
		}

		err := ethrelayClient.DepositStake(stakeFlagChain, amountInWei)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Successfully deposited stake: %f ETH\n", utils.WeiToEther(amountInWei))
	},
}

func init() {
	stakeCmd.AddCommand(stakeDepositCmd)
}
