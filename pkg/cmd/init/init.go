package init

import "github.com/spf13/cobra"

func NewCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [flags]",
		Short: "Initialize daje on your system",
		RunE: func(cmd *cobra.Command, args []string) error {
			return createRun()
		},
	}

	return cmd
}

func createRun() error {
	return submitAction()
}

func submitAction() error {
	return nil
}
