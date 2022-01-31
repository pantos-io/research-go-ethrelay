package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pantos-io/go-ethrelay/pkg/ethrelay"
	"github.com/spf13/cobra"
)

func writeToJson(fileName string, data interface{}) string {
	f, err := os.Create(fmt.Sprintf("./%s.json", fileName))
	checkError(err)
	defer f.Close()

	bytes, err := json.MarshalIndent(data, "", "\t")
	checkError(err)

	_, err = f.Write(bytes)
	checkError(err)

	return f.Name()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func chainCompletionFn(chainType ethrelay.ChainType) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		readConfig()
		ids := client.Chains(chainType)
		filteredIds := make([]string, 0, len(ids))

		for _, id := range ids {
			if strings.HasPrefix(id, toComplete) {
				filteredIds = append(filteredIds, id)
			}
		}

		return filteredIds, cobra.ShellCompDirectiveNoFileComp
	}
}

var flagsMap = map[string]func(*cobra.Command, *string){
	"source": func(cmd *cobra.Command, ptr *string) {
		cmd.Flags().StringVarP(ptr, "source", "s", "mainnet", "The identifier of a source blockchain, as set up in the config file")
		cmd.RegisterFlagCompletionFunc("source", chainCompletionFn(ethrelay.ChainTypeSrc))
	},
	"destination": func(cmd *cobra.Command, ptr *string) {
		cmd.Flags().StringVarP(ptr, "destination", "d", "local", "The identifier of a destination blockchain, as set up in the config file")
		cmd.RegisterFlagCompletionFunc("destination", chainCompletionFn(ethrelay.ChainTypeDst))
	},
	"chain": func(cmd *cobra.Command, ptr *string) {
		cmd.Flags().StringVar(ptr, "chain", "mainnet", "The identifier of a (source or destination) blockchain, as set up in the config file")
		cmd.RegisterFlagCompletionFunc("chain", chainCompletionFn(ethrelay.ChainTypeAny))
	},

}

func addCommonFlag(cmd *cobra.Command, flagName string, ptr *string) {
	flagsMap[flagName](cmd, ptr)
}