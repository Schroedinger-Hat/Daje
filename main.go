package main

import (
	"log"
	"os"

	"github.com/Schroedinger-Hat/Daje/constants"
	"github.com/Schroedinger-Hat/Daje/internal/config"
	"github.com/Schroedinger-Hat/Daje/pkg/cmd/root"
	"github.com/spf13/viper"
)

type exitCode int

const (
	exitOk    exitCode = 0
	exitError exitCode = 1
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName(constants.ConfigFileName)
	err := config.LoadConfig()
	if err != nil {
		errorMessage := "main->" + err.Error()
		log.Fatal(errorMessage)
	}
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	cmd := root.NewCmdRoot()
	if _, err := cmd.ExecuteC(); err != nil {
		return exitError
	}

	return exitOk
}
