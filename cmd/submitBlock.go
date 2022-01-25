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
	Use:   "block {blockNumber | blockHash}",
	Short: "Submits a block from a source chain to a destination chain",
	Long:  "Queries the given block from a source chain and submits it to a destination chain",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if submitFlagLiveMode {
			ethrelayClient = createEthrelayClient()
			// TODO: live mode should be variable, outsource this to terminal
			ethrelayClient.SubmitHeaderLive(submitFlagDstChain, submitFlagSrcChain, 5*time.Minute)

			return
		}

		var header *types.Header = nil
		var err error

		ethrelayClient = createEthrelayClient()

		if len(args) > 0 {
			if strings.HasPrefix(args[0], "0x") {
				blockHash := common.HexToHash(args[0])
				header, err = ethrelayClient.HeaderByHash(getFlagChain, blockHash)
			} else {
				var ok bool
				var blockNumber *big.Int = nil
				blockNumber = new(big.Int)
				blockNumber, ok = blockNumber.SetString(args[0], 10)

				if !ok {
					log.Fatalf("Illegal block number '%s'", args[0])
				}

				header, err = ethrelayClient.HeaderByNumber(submitFlagSrcChain, blockNumber)
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
			header = ethrelayClient.RandomizeHeader(submitFlagSrcChain, header)
		}

		fmt.Printf("Submitting block %s of chain '%s' to chain '%s'...\n", header.Number, submitFlagSrcChain, submitFlagDstChain)

		//header.Nonce = types.EncodeNonce(header.Nonce.Uint64() + 1)  // can be used for testing PoW validation

		err = ethrelayClient.SubmitHeader(submitFlagDstChain, header)
		if err != nil {
			log.Fatal("Failed to submit header: " + err.Error())
		}
	},
}

func init() {
	submitCmd.AddCommand(submitBlockCmd)

	submitBlockCmd.Flags().BoolVarP(&submitFlagLiveMode, "live", "l", false, "live mode (continuously submits most recent block headers)")
	submitBlockCmd.Flags().StringVar(&submitFlagSrcChain, "source", "mainnet", "source chain")
	submitBlockCmd.Flags().BoolVarP(&submitFlagRandomize, "randomize", "r", false, "randomize block")
	submitBlockCmd.Flags().StringVarP(&submitFlagParent, "parent", "p", "", "set parent explicitly")
}
