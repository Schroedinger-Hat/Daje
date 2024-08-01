package dotfiles

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/Schroedinger-Hat/Daje/constants"
	"github.com/spf13/viper"
)

func InitializeDotfiles() error {
	err := os.MkdirAll(path.Join(
		constants.DajeBasePath,
		viper.GetString("dotfiles_directory"),
	), os.ModePerm)

	if err != nil {
		errorMessage := "InitializeDotfiles->" + err.Error()
		log.Fatal(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}
