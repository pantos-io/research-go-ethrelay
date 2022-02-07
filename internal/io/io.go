package io

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pantos-io/go-ethrelay/pkg/ethrelay"
	"github.com/spf13/viper"
)

func ReadConfig(cfgFile string) (*ethrelay.Client, error) {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("can't read config file: %s", err)
	}

	chainsConfig := viper.Get("chains").(map[string]interface{})
	privateKey := viper.Get("privateKey").(string)

	return ethrelay.NewClient(privateKey, chainsConfig), nil
}

func WriteToJson(fileName string, data interface{}) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	
	defer f.Close()

	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}