package checkhealth

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Schrodinger-Hat/Daje/constants"
	"github.com/Schrodinger-Hat/Daje/internal/tuning"
)

func NewCmdCheckhealth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checkhealth [flags]",
		Short: "Check Daje health",
		RunE: func(cmd *cobra.Command, args []string) error {
			return submitAction()
		},
	}

	return cmd
}

func submitAction() error {

	fmt.Println("Version: ", constants.Version)
	fmt.Println("Configuration path: ", constants.DajeDotfilePath)
	fmt.Println("Tuning: ", tuning.IsSystemTuned())

	return nil
}
