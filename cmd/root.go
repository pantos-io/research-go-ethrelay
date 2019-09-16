// All commands are child commands of the root command.
// Authors: Marten Sigwart, Philipp Frauenthaler

package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pantos-io/go-testimonium/ethereum/ethash"
	"github.com/pantos-io/go-testimonium/testimonium"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"os"
	"sync"
)

var cfgFile string
var noSubmitFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-testimonium",
	Short: "Starts the Testimonium client",
	Long: `Starts the Testimonium client`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		client := createTestimoniumClient()

		latestBlockNumberChannel := make(chan *big.Int, 1)

		var wg sync.WaitGroup
		if !noSubmitFlag {
			wg.Add(1)
			go submitBlockHeaders(&wg, client, latestBlockNumberChannel)
		}

		wg.Add(1)
		go validateBlockHeaders(&wg, client, latestBlockNumberChannel)
		wg.Wait()
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
	rootCmd.Flags().BoolVar(&noSubmitFlag, "no-submit", false, "do not submit new block headers if this flag is set")
	//rootCmd.Flags().BoolVar(&noSubmitFlag, "no-dispute", false, "Do not validate blocks if this flag is set")
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

func submitBlockHeaders(wg *sync.WaitGroup, client *testimonium.Client, latestBlockNumberChannel chan *big.Int) {
	defer wg.Done()
	for {
		select {
			case latest := <-latestBlockNumberChannel:
				blockHeightToSubmit := new(big.Int)
				blockHeightToSubmit.Add(latest, big.NewInt(1))
				header, err := client.HeaderByNumber(blockHeightToSubmit, 0)
				if err != nil {
					fmt.Printf("WARNING: could not get block with height %d from source chain %d: %s\n", blockHeightToSubmit, 0, err)
				}
				fmt.Printf("Submitting block header %s to destination chain %d...\n", header.Hash().String(), 1)
				client.SubmitHeader(header, 1)
				// todo: react to success or failure accordingly
		}
	}
}

func validateBlockHeaders(wg *sync.WaitGroup, client *testimonium.Client, latestBlockNumberChannel chan<- *big.Int) {
	defer wg.Done()
	sink := make(chan *testimonium.TestimoniumSubmitBlockHeader, 1)

	// find latest valid block header of source chain on destination chain
	hash, err := client.LongestChainEndpoint(1)
	if err != nil {
		log.Fatal("could not read longest chain endpoint", err)
	}

	header, err := client.BlockHeader(hash, 1)
	if err != nil {
		log.Fatal("could not read endpoint header", err)
	}

	sink <- toSubmitEvent(hash, header) // create 'fake' submit event to pass to event sink

	// listen to newly submitted headers
	_, err = client.WatchSubmitBlockHeader(1, sink)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-sink:
			fmt.Printf("Validating block header %s (height %d)...\n", ShortHexString(common.BytesToHash(event.Hash[:]).String()), event.BlockNumber)
			if isValidBlockHeader(client, event.Hash, 0) {
				fmt.Printf("Block header %s is valid.\n", ShortHexString(common.BytesToHash(event.Hash[:]).String()))
				select {
					case latestBlockNumberChannel <- event.BlockNumber:
						break
					default:
						// non-blocking behaviour in case no routine is listening to the latestBlockNumberChannel
						// e.g., if the --no-submit flag is set
						break
				}

			} else {
				// Dispute header
				fmt.Printf("Validation Failed! Disputing block %s...\n", ShortHexString(common.BytesToHash(event.Hash[:]).String()))

				go func() {
					header, err := client.BlockHeader(event.Hash, 1)
					if err != nil {
						log.Fatal("Failed to retrieve header from contract: " + err.Error())
					}
					blockMetaData := ethash.NewBlockMetaData(header.BlockNumber.Uint64(), header.Nonce.Uint64(), header.RlpHeaderHashWithoutNonce)
					dataSetLookup := blockMetaData.DAGElementArray()
					witnessForLookup := blockMetaData.DAGProofArray()
					client.DisputeBlock(event.Hash, dataSetLookup, witnessForLookup, 1)

					// Query longest chain
					hash, err := client.LongestChainEndpoint(1)
					if err != nil {
						log.Fatal("could not read longest chain endpoint", err)
					}

					header, err = client.BlockHeader(hash, 1)
					if err != nil {
						log.Fatal("could not read endpoint header", err)
					}

					sink <- toSubmitEvent(hash, header) // create 'fake' submit event to pass to event sink
				}()

			}
		}
	}
}

func isValidBlockHeader(client *testimonium.Client, blockHash [32]byte, chain uint8) bool {
	_, err := client.OriginalBlockHeader(blockHash, chain)
	if err != nil {
		// if an error is returned, it means that no block with the specified block hash exists
		return false;
	}
	return true
}

func toSubmitEvent(hash common.Hash, header testimonium.BlockHeader) *testimonium.TestimoniumSubmitBlockHeader {
	submitEvent := new(testimonium.TestimoniumSubmitBlockHeader)
	submitEvent.Hash = hash
	submitEvent.BlockNumber = header.BlockNumber
	return submitEvent
}


func ShortHexString(hex string) string {
	if len(hex) <= 12 {
		return hex
	}
	return fmt.Sprintf("%s...%s", hex[:6], hex[len(hex)-4:])
}

