package main

import (
	"os"

	"github.com/Schroedinger-Hat/Daje/constants"
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
	viper.SetConfigName(constants.DajeConfigFileName)
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
