package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pantos-io/go-ethrelay/ethrelay"
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

func chainFlagCompletionFn(chainType ethrelay.ChainType) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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