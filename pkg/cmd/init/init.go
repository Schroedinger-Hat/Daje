package init

import (
	"errors"
	"log"

	"github.com/spf13/cobra"

	"github.com/Schroedinger-Hat/Daje/internal/config"
	"github.com/Schroedinger-Hat/Daje/internal/dotfiles"
)

func CmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [flags]",
		Short: "Initialize daje on your system",
		RunE: func(cmd *cobra.Command, args []string) error { //nolint:all
			return submitCmdInit()
		},
	}

	return cmd
}

func submitCmdInit() error {
	funcErrorHandle := func(err error) error {
		errorMessage := "CmdInit->" + err.Error()
		log.Fatal(errorMessage)
		return errors.New(errorMessage)
	}
	functionOrder := []func() error{
		config.LoadConfig,
		dotfiles.InitializeDotfiles,
	}

	for _, function := range functionOrder {
		if err := function(); err != nil {
			return funcErrorHandle(err)
		}
	}
	return nil
}
