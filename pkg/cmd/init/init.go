package init

import "github.com/spf13/cobra"

func NewCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daje init [flags]",
		Short: "Initialize daje on your system",
	}

	return cmd
}
