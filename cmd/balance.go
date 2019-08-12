// This file contains logic executed if the command "balance" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

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

		testimoniumClient = createTestimoniumClient()
		if len(args) > 0 {
			chainId, err := strconv.ParseInt(args[0], 10, 8)
			if err != nil {
				log.Fatal(err)
			}
			balance, err := testimoniumClient.Balance(uint8(chainId))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%.4f ETH\n", getDecimal(balance, 18))
			return
		}

		if detailFlag {
			totalBalance := big.NewInt(0)
			for _, chainId := range testimoniumClient.Chains() {
				balance, err := testimoniumClient.Balance(uint8(chainId))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Chain %d: %.4f ETH\n", chainId, getDecimal(balance, 18))
				totalBalance = totalBalance.Add(totalBalance, balance)
			}
			fmt.Printf("Total  : ")
		}
		balance, err := testimoniumClient.TotalBalance()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	balanceCmd.Flags().BoolVarP(&detailFlag, "detail", "d", false, "display detailed balance information")

}
