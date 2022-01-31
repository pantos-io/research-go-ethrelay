// All commands are child commands of the root command.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/pantos-io/go-ethrelay/pkg/ethrelay"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var client *ethrelay.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-ethrelay",
	Short: "The CLI to interact with the ETH Relay prototype",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
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
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Can't read config file:", err)
		return
	}

	chainsConfig := viper.Get("chains").(map[string]interface{})
	privateKey := viper.Get("privateKey").(string)

	client = ethrelay.NewClient(privateKey, chainsConfig)
}
