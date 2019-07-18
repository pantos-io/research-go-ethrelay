// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

var verifyFlagSrcChain uint8
var verifyFlagDestChain uint8

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies a transaction or a block from the source chain on the destination chain",
	Long: `Verifies a transaction or a block from the source chain on the destination chain`,
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	verifyCmd.PersistentFlags().Uint8Var(&verifyFlagSrcChain, "source", 2, "source chain (default: 2)")
	verifyCmd.PersistentFlags().Uint8Var(&verifyFlagDestChain, "destination", 1, "destination chain (default: 1)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
