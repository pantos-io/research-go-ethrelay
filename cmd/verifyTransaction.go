// This file contains logic executed if the command "verify transaction" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pantos-io/go-ethrelay/ethrelay"
	"github.com/spf13/cobra"
)

var noOfConfirmations uint8

// verifyTransactionCmd represents the transaction command
var verifyTransactionCmd = &cobra.Command{
	Use:   "transaction txHash",
	Aliases: []string{"tx"},
	Short: "Verifies a transaction",
	Long: `Verifies a transaction from a source chain on a destination chain
Behind the scene, the command queries the transaction with the specified hash from the source chain.
It then generates a Merkle Proof contesting the existence of the transaction within a specific block.
This information gets sent to the destination chain, where not only the existence of the block but also the Merkle Proof are verified.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])

		rlpHeader, proof, err := client.GenerateMerkleProofForTx(verifyFlagSrcChain, txHash)
		if err != nil {
			log.Fatal("Failed to generate Merkle Proof: " + err.Error())
		}

		// TODO: this produces a merkle proof for the transaction and does not verify the transaction
		//  maybe it is better to introduce a new command for this behaviour as it is quite confusing to
		//  call verifyTransaction and no transaction is verified
		if jsonFlag {
			fileName := writeToJson(fmt.Sprint("tx_", txHash), proof)
			fmt.Println("Wrote merkle proof to", fileName)
			return
		}

		feesInWei, err := client.GetRequiredVerificationFee(verifyFlagDstChain)
		if err != nil {
			log.Fatal(err)
		}

		client.VerifyMerkleProof(verifyFlagDstChain, feesInWei, rlpHeader, ethrelay.ValueTypeTransaction, proof, noOfConfirmations)
	},
}

func init() {
	verifyCmd.AddCommand(verifyTransactionCmd)

	verifyTransactionCmd.Flags().Uint8Var(&noOfConfirmations, "confirmations", 4, "Number of block confirmations")
	verifyTransactionCmd.Flags().BoolVar(&jsonFlag, "json", false, "save merkle proof to a json file")
}
