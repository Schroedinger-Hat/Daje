package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Schrodinger-Hat/Daje/constants"
)

func InitEmptyDaje() error {
	dajeConfigPath := filepath.Join(constants.DajeConfigBaseDir)
	if err := os.Mkdir(dajeConfigPath, 0755); err != nil {
		return err
	}

	dajeConfigFile := filepath.Join(constants.DajeConfigBaseDir, constants.DajeDotfileName)
	if _, err := os.Create(dajeConfigFile); err != nil {
		return err
	}

	return nil
}

func IsDajeInitialized() bool {
	dotFile := filepath.Join(constants.DajeConfigBaseDir, constants.DajeDotfileName)
	if _, err := os.Stat(dotFile); err != nil {
		if os.IsNotExist(err) {
			return false
		}

		log.Fatal(err)
	}

	return true
}
