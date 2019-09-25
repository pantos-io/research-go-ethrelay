// This file contains logic executed if the command "stake deposit" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"
)

// stakeDepositCmd represents the command 'stake deposit <amount>'
var stakeDepositCmd = &cobra.Command{
	Use:   "deposit [amountInWei]",
	Short: "Deposits the specified amount of Wei.",
	Long: `Deposits the specified amount of Wei, i.e., the client's stake is increased by the specified amount'`,
	Run: func(cmd *cobra.Command, args []string) {
		testimoniumClient = createTestimoniumClient()
		amountInWei, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = testimoniumClient.DepositStake(stakeFlagChain, big.NewInt(int64(amountInWei)))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Deposit was successful")
	},
}

func init() {
	stakeCmd.AddCommand(stakeDepositCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	//stakeCmd.PersistentFlags().Uint8Var(&stakeFlagChain, "chain", 1, "chain")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
