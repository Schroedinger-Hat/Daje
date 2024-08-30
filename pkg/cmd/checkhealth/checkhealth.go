package checkhealth

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Schroedinger-Hat/Daje/constants"
)

func CmdCheckhealth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checkhealth [flags]",
		Short: "Check Daje health",
		RunE: func(cmd *cobra.Command, args []string) error { //nolint:all
			return submitCmdCheckhealth()
		},
	}

	return cmd
}

func submitCmdCheckhealth() error {
	log.Println("[Checkhealth]:[LoadConfig]", "Configuration Path:"+viper.GetViper().ConfigFileUsed())

	for _, value := range constants.ConfigParameters {
		log.Println("[Checkhealth]:[LoadConfig]", value+": "+viper.GetString(value))
	}

	return nil
}
