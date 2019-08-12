// This file contains logic executed if the command "dispute" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pf92/go-testimonium/ethereum/ethash"

	"log"

	"github.com/spf13/cobra"
)

var disputeFlagChain uint8

// disputeCmd represents the dispute command
var disputeCmd = &cobra.Command{
	Use:   "dispute [blockHash]",
	Short: "Disputes a submitted block header",
	Long: `Disputes the submitted block header with the specified hash ('blockHash')`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		blockHash := common.HexToHash(args[0])	// omit the first two chars "0x"
		blockHashBytes := blockHash.Bytes()
		var blockHashBytes32 [32]byte
		copy(blockHashBytes32[:], blockHashBytes)

		// get blockNumber, nonce and RlpHeaderHashWithoutNonce and generate dataSetLookup and witnessForLookup
		testimoniumClient = createTestimoniumClient()
		header, err := testimoniumClient.BlockHeader(blockHashBytes32, disputeFlagChain)
		if err != nil {
			log.Fatal("Failed to retrieve header from contract: " + err.Error())
		}

		fmt.Println("create DAG, compute dataSetLookup and witnessForLookup")
		// get DAG and compute dataSetLookup and witnessForLookup
		blockMetaData := ethash.NewBlockMetaData(header.BlockNumber.Uint64(), header.Nonce.Uint64(), header.RlpHeaderHashWithoutNonce)
		dataSetLookup := blockMetaData.DAGElementArray()
		witnessForLookup := blockMetaData.DAGProofArray()

		testimoniumClient.DisputeBlock(blockHash, dataSetLookup, witnessForLookup, disputeFlagChain)
	},
}

func init() {
	rootCmd.AddCommand(disputeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disputeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disputeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	disputeCmd.Flags().Uint8VarP(&disputeFlagChain, "chain", "c", 1, "the disputed chain ID")
}
