package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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