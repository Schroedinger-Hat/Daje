package root

import (
	"github.com/spf13/cobra"

	"github.com/Schroedinger-Hat/Daje/pkg/cmd/checkhealth"
	initCmd "github.com/Schroedinger-Hat/Daje/pkg/cmd/init"
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

	cmd.AddCommand(initCmd.CmdInit())
	cmd.AddCommand(checkhealth.CmdCheckhealth())

	return cmd
}
