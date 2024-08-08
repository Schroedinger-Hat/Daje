package main

import (
	"os"
	"testing"
)

// For fun
func Test_MainRunShouldExitOkWhenNoSubcommand(t *testing.T) {
	_, w, _ := os.Pipe()
	stdout := os.Stdout
	stderr := os.Stderr

	os.Stdout = w
	os.Stderr = w

	if mainRun() != exitOk {
		t.Fatal("mainRun() not ok")
	}

	_ = w.Close()
	os.Stdout = stdout
	os.Stderr = stderr
}
