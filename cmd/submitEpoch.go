// This file contains logic executed if the command "submit epoch" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
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
	Short: "Sets the epoch data for the specified epoch on the verifying chain",
	Long: `Sets the epoch data for the specified epoch on the verifying chain`,
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
			writeEpochAsJson(epochData, epoch)
			fmt.Printf("Wrote epoch data to %s.json\n", epoch.String())
			return
		}
		testimoniumClient = createTestimoniumClient()
		testimoniumClient.SetEpochData(submitFlagDestChain, epochData)
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


func writeEpochAsJson(epochData typedefs.EpochData, epoch *big.Int) {
	f, err := os.Create(fmt.Sprintf("./%s.json", epoch.String()))
	checkError(err)
	defer f.Close()

	_, err = fmt.Fprint(f, "{\n")
	_, err = fmt.Fprintf(f, "  \"epoch\": \"%s\",\n", epochData.Epoch.String())
	_, err = fmt.Fprintf(f, "  \"fullSizeIn128Resolution\": \"%s\",\n", epochData.FullSizeIn128Resolution.String())
	_, err = fmt.Fprintf(f, "  \"branchDepth\": \"%s\",\n", epochData.BranchDepth.String())
	_, err = fmt.Fprintf(f, "  \"merkleNodes\": ",)
	writeElementsToFile(f, epochData.MerkleNodes)
	_, err = fmt.Fprint(f, "\n}")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeElementsToFile(file *os.File, arr []*big.Int) {
	_, err := fmt.Fprintf(file, "[")
	checkError(err)
	if len(arr) > 0 {
		_, err = fmt.Fprintf(file, "\"%s\"", arr[0].String())
		checkError(err)
	}

	for i := 1; i < len(arr); i++ {
		_, err = fmt.Fprintf(file, ", \"%s\"", arr[i].String())
		checkError(err)
	}
	_, err = fmt.Fprintf(file, "]")
	checkError(err)
}
