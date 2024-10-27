package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/schroedinger-Hat/Daje/constants"
)

func InitEmptyDaje() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}
	dotFile := filepath.Join(currentUser.HomeDir, constants.DajeDotFile)
	if err = os.Mkdir(dotFile, os.ModeDir); err != nil {
		return err
	}
	return nil
}

func IsDajeInitialized() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dotFile := filepath.Join(currentUser.HomeDir, constants.DajeDotFile)
	if _, err := os.Stat(dotFile); err != nil {
		if os.IsNotExist(err) {
			return false
		}

		log.Fatal(err)
	}

	return true
}
