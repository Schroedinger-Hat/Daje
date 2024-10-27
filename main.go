package main

import (
	"os"

	"github.com/schroedinger-Hat/Daje/pkg/cmd/root"
)

type exitCode int

const (
	exitOk    exitCode = 0
	exitError exitCode = 1
)

func main() {
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
