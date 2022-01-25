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
	Use:   "receipt [txHash]",
	Short: "Verifies a receipt",
	Long: `Verifies a receipt from a source chain on a destination chain

Behind the scene, the command queries the receipt with the specified hash ('txHash') from the source chain.
It then generates a Merkle Proof contesting the existence of the receipt within a specific block.
This information gets sent to the destination chain, where not only the existence of the block but also the Merkle Proof are verified`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])

		ethrelayClient = createEthrelayClient()

		rlpHeader, proof, err := ethrelayClient.GenerateMerkleProofForReceipt(verifyFlagSrcChain, txHash)
		if err != nil {
			log.Fatal("Failed to generate Merkle Proof: " + err.Error())
		}

		feesInWei, err := ethrelayClient.GetRequiredVerificationFee(verifyFlagDstChain)
		if err != nil {
			log.Fatal(err)
		}

		ethrelayClient.VerifyMerkleProof(verifyFlagDstChain, feesInWei, rlpHeader, ethrelay.ValueTypeReceipt, proof, noOfConfirmations)
	},
}

func init() {
	verifyCmd.AddCommand(verifyReceiptCmd)

	verifyReceiptCmd.Flags().Uint8VarP(&noOfConfirmations, "confirmations", "c", 4, "Number of block confirmations")
}
