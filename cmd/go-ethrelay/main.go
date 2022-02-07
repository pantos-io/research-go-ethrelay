// All commands are child commands of the root command.
// Authors: Marten Sigwart, Philipp Frauenthaler

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pantos-io/go-ethrelay/internal/io"
	"github.com/pantos-io/go-ethrelay/pkg/ethrelay"
	"github.com/spf13/cobra"
)

var cfgFile string
var client *ethrelay.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-ethrelay",
	Short: "The CLI to interact with the ETH Relay prototype",
}

// Adds all child commands to the root command and sets flags appropriately.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(readConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./ethrelay.yml", "YAML config file")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readConfig() {
	var err error
	if client, err = io.ReadConfig(cfgFile); err != nil {
		log.Fatalln("Failed to read in config:", err)
	}
}
