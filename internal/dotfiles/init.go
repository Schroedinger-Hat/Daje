package dotfiles

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

func InitializeDotfiles() error {
	err := os.MkdirAll(path.Join(viper.GetString("dotfiles.local")), os.ModePerm)
	if err != nil {
		errorMessage := "InitializeDotfiles->" + err.Error()
		log.Fatal(errorMessage)
		return errors.New(errorMessage)
	}

	// TODO: clone dotfiles.remote
	return nil
}
