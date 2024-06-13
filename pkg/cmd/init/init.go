package init

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Schrodinger-Hat/Daje/internal/config"
)

func NewCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [flags]",
		Short: "Initialize daje on your system",
		RunE: func(cmd *cobra.Command, args []string) error { //nolint:all
			return submitAction()
		},
	}

	return cmd
}

func submitAction() error {
	if !config.IsDajeInitialized() {
		err := config.InitEmptyDaje()
		if err != nil {
			fmt.Println(err)
			return nil
		}

		fmt.Println("Daje has been initialized successfully!")
	} else {
		fmt.Println("Daje has been already initialized in the system.")
		return nil
	}

	return nil
}
