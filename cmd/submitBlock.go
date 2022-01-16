// This file contains logic executed if the command "submit block" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

var submitFlagSrcChain string
var submitFlagRandomize bool
var submitFlagParent string
var submitFlagLiveMode bool

// submitCmd represents the submit command
var submitBlockCmd = &cobra.Command{
	Use:   "block [blockNumber or blockHash]",
	Short: "Submits a block header from source chain to verifying chain",
	Long:  `Queries the given block from the source chain and submits it to the verifying chain`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if submitFlagLiveMode {
			testimoniumClient = createTestimoniumClient()
			// TODO: live mode should be variable, outsource this to terminal
			testimoniumClient.SubmitHeaderLive(submitFlagDestChain, submitFlagSrcChain, 5*time.Minute)

			return
		}

		var header *types.Header = nil
		var err error

		testimoniumClient = createTestimoniumClient()

		if len(args) > 0 {
			if strings.HasPrefix(args[0], "0x") {
				blockHash := common.HexToHash(args[0])
				header, err = testimoniumClient.HeaderByHash(blockHash, getFlagChain)
			} else {
				var ok bool
				var blockNumber *big.Int = nil
				blockNumber = new(big.Int)
				blockNumber, ok = blockNumber.SetString(args[0], 10)

				if !ok {
					log.Fatalf("Illegal block number '%s'", args[0])
				}

				header, err = testimoniumClient.HeaderByNumber(blockNumber, submitFlagSrcChain)
			}
		}

		if err != nil {
			log.Fatal("Failed to retrieve header: " + err.Error())
		}

		if len(submitFlagParent) > 0 {
			fmt.Printf("Modifying parent...\n")
			header.ParentHash = common.HexToHash(submitFlagParent)
		}

		if submitFlagRandomize {
			fmt.Printf("Randomizing header...\n")
			header = testimoniumClient.RandomizeHeader(header, submitFlagSrcChain)
		}

		fmt.Printf("Submitting block %s of chain '%s' to chain '%s'...\n", header.Number.String(), submitFlagSrcChain, submitFlagDestChain)

		//header.Nonce = types.EncodeNonce(header.Nonce.Uint64() + 1)  // can be used for testing PoW validation

		err = testimoniumClient.SubmitHeader(header, submitFlagDestChain)
		if err != nil {
			log.Fatal("Failed to submit header: " + err.Error())
		}
	},
}

func init() {
	submitCmd.AddCommand(submitBlockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	submitBlockCmd.Flags().BoolVarP(&submitFlagLiveMode, "live", "l", false, "live mode (continuously submits most recent block headers)")
	submitBlockCmd.Flags().StringVar(&submitFlagSrcChain, "source", "mainnet", "source chain")
	submitBlockCmd.Flags().BoolVarP(&submitFlagRandomize, "randomize", "r", false, "randomize block")
	submitBlockCmd.Flags().StringVarP(&submitFlagParent, "parent", "p", "", "set parent explicitly")
}
