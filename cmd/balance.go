// This file contains logic executed if the command "balance" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/spf13/cobra"
)

var detailFlag bool

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance [chain]",
	Short: "Prints the balance of the current account",
	Long: `Prints the balance of the current account.
	If [chain] is set, it prints the balance of the current account on the specified chain.
	If not, it prints the total balance`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ethrelayClient = createEthrelayClient()
		if len(args) > 0 {
			balance, err := ethrelayClient.Balance(args[0])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%.4f ETH\n", getDecimal(balance, 18))
			return
		}

		if detailFlag {
			totalBalance := big.NewInt(0)
			for _, chainId := range ethrelayClient.Chains() {
				balance, err := ethrelayClient.Balance(chainId)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Chain '%s': %.4f ETH\n", chainId, getDecimal(balance, 18))
				totalBalance = totalBalance.Add(totalBalance, balance)
			}
			fmt.Printf("Total: ")
		}
		balance, err := ethrelayClient.TotalBalance()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.4f ETH\n", getDecimal(balance, 18))
	},
}

func getDecimal(absolute *big.Int, decimals int) *big.Float {
	decimal := new(big.Float)
	decimal.SetString(absolute.String())
	return new(big.Float).Quo(decimal, big.NewFloat(math.Pow10(decimals)))
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.Flags().BoolVarP(&detailFlag, "detail", "d", false, "display detailed balance information")
}
