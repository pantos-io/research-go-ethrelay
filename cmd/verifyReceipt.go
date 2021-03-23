// This file contains logic executed if the command "verify receipt" is typed in.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pantos-io/go-ethrelay/testimonium"
	"github.com/spf13/cobra"
)

// verifyReceiptCmd represents the receipt command
var verifyReceiptCmd = &cobra.Command{
	Use:   "receipt [txHash]",
	Short: "Verifies a receipt",
	Long: `Verifies a receipt from the source chain on the verifying chain

Behind the scene, the command queries the receipt with the specified hash ('txHash') from the source chain.
It then generates a Merkle Proof contesting the existence of the receipt within a specific block.
This information gets sent to the verifying chain, where not only the existence of the block but also the Merkle Proof are verified`,
	Aliases: []string{"tx"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(args[0])

		testimoniumClient = createTestimoniumClient()

		rlpHeader, rlpEncodedReceipt, path, rlpEncodedProofNodes, err := testimoniumClient.GenerateMerkleProofForReceipt(txHash, verifyFlagSrcChain)
		if err != nil {
			log.Fatal("Failed to generate Merkle Proof: " + err.Error())
		}

		feesInWei, err := testimoniumClient.GetRequiredVerificationFee(verifyFlagDestChain)
		if err != nil {
			log.Fatal(err)
		}

		testimoniumClient.VerifyMerkleProof(feesInWei, rlpHeader, testimonium.VALUE_TYPE_RECEIPT, rlpEncodedReceipt, path,
			rlpEncodedProofNodes, noOfConfirmations, verifyFlagDestChain)
	},
}

func init() {
	verifyCmd.AddCommand(verifyReceiptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	verifyReceiptCmd.Flags().Uint8VarP(&noOfConfirmations, "confirmations", "c", 4, "Number of block confirmations")
}
