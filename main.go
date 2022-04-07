package main

import (
	"github.com/Schrodinger-Hat/Daje/pkg/cmd/root"
	"os"
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
