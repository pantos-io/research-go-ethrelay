// This file contains logic executed if the command "balance" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"
	"math/big"

	"github.com/pantos-io/go-ethrelay/ethereum/utils"
	"github.com/pantos-io/go-ethrelay/ethrelay"
	"github.com/spf13/cobra"
)

var detailFlag bool

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance [chain]",
	Short: "Prints the balance of the current account",
	Long:  `Prints the balance of the current account.
If a chain is set, it prints the balance of the current account on the specified chain.
If not, it prints the total balance`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			balance, err := client.Balance(args[0])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%f ETH\n", utils.WeiToEther(balance))
			return
		}

		if detailFlag {
			totalBalance := big.NewInt(0)
			for _, chainId := range client.Chains(ethrelay.ChainTypeAny) {
				balance, err := client.Balance(chainId)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Chain '%s': %f ETH\n", chainId, utils.WeiToEther(balance))
				totalBalance = totalBalance.Add(totalBalance, balance)
			}
			fmt.Printf("Total: ")
		}
		balance, err := client.TotalBalance()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%f ETH\n", utils.WeiToEther(balance))
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.Flags().BoolVarP(&detailFlag, "detail", "d", false, "display detailed balance information")
}
