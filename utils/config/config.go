package config

import (
	"encoding/json"
	"os"
)

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
