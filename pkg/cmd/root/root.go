package root

import (
	"github.com/spf13/cobra"

	initCmd "github.com/Schrodinger-Hat/Daje/pkg/cmd/init"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daje <command> <subcommand> [flags]",
		Short: "Daje CLI",
		Long:  "A command line to handle your dotfiles",
		Example: `
$ daje init
#=> Initializes Daje in your system
`,
	}

	cmd.AddCommand(initCmd.NewCmdInit())

	return cmd
}
