package config

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/Schroedinger-Hat/Daje/constants"
	"github.com/spf13/viper"
)

func ExtractConfigParameter(elementName string) (string, error) {
	return viper.GetString(elementName), nil
}

func LoadConfig() error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		errorMessage := "LoadConfig:getConfigFilePath->" + err.Error()
		log.Fatal(errorMessage)
		return errors.New(errorMessage)
	}
	viper.SetConfigFile(configFilePath)
	if err = viper.ReadInConfig(); err != nil {
		errorMessage := "LoadConfig:SetConfigFile->" + err.Error()
		log.Fatal(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	for _, value := range constants.DajeConfigPathOrder {
		currentFilepath := path.Join(constants.DajeBasePath, value, constants.DajeConfigFileName)
		if _, err := os.Stat(currentFilepath); err == nil {
			return currentFilepath, nil
		}
	}
	return "", errors.New("getConfigFilePath: Configuration not found")
}
