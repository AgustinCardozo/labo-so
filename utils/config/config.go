package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func InitConfig(filePath string, config interface{}) {
	err := SetupConfig(filePath, &config)
	if err != nil {
		fmt.Errorf("Error setting up configuration file: %v", err)
		panic(err)
	}
}

func SetupConfig(filePath string, config interface{}) error {
	configFile, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)

	if err := jsonParser.Decode(&config); err != nil {
		return err
	}

	return nil
}
