// This file contains logic executed if the command "submit epoch" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
	"fmt"
	"math/big"

	"github.com/pantos-io/go-ethrelay/internal/ethereum/ethash"
	"github.com/pantos-io/go-ethrelay/internal/io"

	"log"

	"github.com/spf13/cobra"
)

var jsonFlag bool

// submitEpochCmd represents the command for setting epoch data (Ethash contract)
var submitEpochCmd = &cobra.Command{
	Use:   "epoch epochNumber",
	Short: "Sets the epoch data for the specified epoch on a destination chain",
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
			fileName := io.WriteToJson(epoch.String(), epochData)
			fmt.Println("Wrote epoch data to", fileName)
			return
		}
		client.SetEpochData(submitFlagDstChain, epochData)
	},
}

func init() {
	submitCmd.AddCommand(submitEpochCmd)

	submitEpochCmd.Flags().BoolVar(&jsonFlag, "json", false, "creates a JSON file containing the epoch data without submitting it")
}
