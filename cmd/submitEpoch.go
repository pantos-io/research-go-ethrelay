// This file contains logic executed if the command "submit epoch" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/pantos-io/go-ethrelay/ethereum/ethash"
	"github.com/pantos-io/go-ethrelay/typedefs"

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
			fileName := writeEpochAsJson(epochData, epoch)
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


func writeEpochAsJson(epochData typedefs.EpochData, epoch *big.Int) string {
	f, err := os.Create(fmt.Sprintf("./epoch_%s.json", epoch))
	checkError(err)
	defer f.Close()

	bytes, err := json.MarshalIndent(epochData, "", "\t")
	checkError(err)

	_, err = f.Write(bytes)
	checkError(err)

	return f.Name()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
