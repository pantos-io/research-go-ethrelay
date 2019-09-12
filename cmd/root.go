// All commands are child commands of the root command.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"github.com/pantos-io/go-testimonium/testimonium"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-testimonium",
	Short: "Starts the Testimonium client",
	Long: `Starts the Testimonium client`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command called")
		client := createTestimoniumClient()
		sink := make(chan *testimonium.TestimoniumSubmitBlockHeader)

		_, err := client.WatchSubmitBlockHeader(1, sink)
		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case event := <-sink:
				fmt.Println(event)
			}
		}
	},
}

var testimoniumClient *testimonium.Client


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/testimonium.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name "testimonium" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("testimonium")
	}

	viper.AutomaticEnv() // read in environment variables that match



}

func createTestimoniumClient() (*testimonium.Client) {
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config file:", err)
	}

	chainsConfig := viper.Get("chains").(map[string]interface{})
	privateKey := viper.Get("privateKey").(string)

	return testimonium.NewClient(privateKey, chainsConfig)
}
