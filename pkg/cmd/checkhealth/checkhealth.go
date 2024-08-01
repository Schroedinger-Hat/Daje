package checkhealth

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Schroedinger-Hat/Daje/constants"
	"github.com/Schroedinger-Hat/Daje/internal/config"
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
	err := config.LoadConfig()
	if err != nil {
		errorMessage := "Checkhealth->" + err.Error()
		log.Fatal(errorMessage)
		return errors.New(errorMessage)
	}
	log.Println("[Checkhealth]:[LoadConfig]", "Configuration Path:"+viper.GetViper().ConfigFileUsed())

	log.Println("[Checkhealth]:[ConfigValues]")
	for _, value := range constants.DajeConfigParameters {
		log.Println(value + ": " + viper.GetString(value))
	}

	return nil
}
