package utils

import (
	"blog/pkg/app/config"
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfig(file string) (config.ServerSetting, error) {
	var config config.ServerSetting
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println("Cannot open file%v", err.Error())
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}