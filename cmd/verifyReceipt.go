// This file contains logic executed if the command "verify receipt" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pantos-io/go-ethrelay/ethrelay"
	"github.com/spf13/cobra"
)

// verifyReceiptCmd represents the receipt command
var verifyReceiptCmd = &cobra.Command{
	Use:   "receipt txHash",
	Short: "Verifies a receipt",
	Long:  `Verifies a receipt from a source chain on a destination chain.
Behind the scene, the command queries the receipt with the specified hash from the source chain.
It then generates a Merkle Proof contesting the existence of the receipt within a specific block.
This information gets sent to the destination chain, where not only the existence of the block but also the Merkle Proof are verified.`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])

		rlpHeader, proof, err := client.GenerateMerkleProofForReceipt(verifyFlagSrcChain, txHash)
		if err != nil {
			log.Fatal("Failed to generate Merkle Proof: " + err.Error())
		}

		feesInWei, err := client.GetRequiredVerificationFee(verifyFlagDstChain)
		if err != nil {
			log.Fatal(err)
		}

		client.VerifyMerkleProof(verifyFlagDstChain, feesInWei, rlpHeader, ethrelay.ValueTypeReceipt, proof, noOfConfirmations)
	},
}

func init() {
	verifyCmd.AddCommand(verifyReceiptCmd)

	verifyReceiptCmd.Flags().Uint8Var(&noOfConfirmations, "confirmations", 4, "Number of block confirmations")
}
