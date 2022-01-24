// This file contains logic executed if the command "submit epoch" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"math/big"

	"github.com/pantos-io/go-ethrelay/ethereum/ethash"

	"log"

	"github.com/spf13/cobra"
)

var jsonFlag bool

// submitEpochCmd represents the command for setting epoch data (Ethash contract)
var submitEpochCmd = &cobra.Command{
	Use:   "epoch [epoch]",
	Short: "Sets the epoch data for the specified epoch on a destination chain",
	Long: "Sets the epoch data for the specified epoch on a destination chain",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var epoch *big.Int = nil
		var ok bool
		epoch = new(big.Int)
		epoch, ok = epoch.SetString(args[0], 10)
		if !ok {
			log.Fatalf("Illegal epoch number '%s'", args[0])
		}

		epochData := ethash.GenerateEpochData(epoch.Uint64())

		if jsonFlag {
			fileName := writeToJson(epoch.String(), epochData)
			fmt.Println("Wrote epoch data to", fileName)
			return
		}
		ethrelayClient = createEthrelayClient()
		ethrelayClient.SetEpochData(submitFlagDstChain, epochData)
	},
}

func init() {
	submitCmd.AddCommand(submitEpochCmd)

	submitEpochCmd.Flags().BoolVar(&jsonFlag, "json", false, "creates a JSON file containing the epoch data without submitting it")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disputeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disputeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
