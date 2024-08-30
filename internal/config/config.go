package config

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/Schroedinger-Hat/Daje/constants"
	"github.com/spf13/viper"
)

type configReturn struct {
	Value string
	Error error
}

func LoadConfig() error {
	configFilePath := ""
	viperReadInConfig := func() configReturn {
		viper.SetConfigFile(configFilePath)
		return configReturn{"", viper.ReadInConfig()}
	}
	functionOrder := []func() configReturn{
		checkBasePath,
		getConfigFilePath,
		viperReadInConfig,
	}

	for _, function := range functionOrder {
		result := function()
		if result.Value != "" {
			configFilePath = result.Value
		}
		if result.Error != nil {
			errorMessage := "LoadConfig->" + result.Error.Error()
			log.Fatal(errorMessage)
			return errors.New(errorMessage)
		}
	}

	pathRelativeToAbsolute(path.Dir(configFilePath))
	return nil
}

func checkBasePath() configReturn {
	if constants.ConfigBasepath == "" {
		homepath, err := os.UserHomeDir()
		if err != nil {
			return configReturn{"", errors.New("getConfigFilePath: User home not found")}
		}
		constants.ConfigBasepath = homepath
	}
	return configReturn{"", nil}
}

func pathRelativeToAbsolute(configDirectory string) {
	for _, value := range constants.ConfigParametersPath {
		configValue := viper.GetString(value)
		viper.Set(value, path.Join(configDirectory, configValue))
	}
}

func getConfigFilePath() configReturn {
	for _, value := range constants.ConfigPathOrder {
		currentFilepath := path.Join(constants.ConfigBasepath, value, constants.ConfigFileName)
		_, err := os.Stat(currentFilepath)
		if err == nil {
			return configReturn{currentFilepath, nil}
		}
	}
	return configReturn{"", errors.New("getConfigFilePath: Configuration not found")}
}
