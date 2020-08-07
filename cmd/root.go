package cmd

import (
	"database/sql"
	"errors"
	"os"

	"artists/internal/translations"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   translations.StringValues.Commands.Root.Use,
	Short: translations.StringValues.Commands.Root.Short,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New(translations.StringValues.Errors.ErrorSubcommandRequired)
	},
}

var DB *sql.DB

func init() {
	err := godotenv.Overload()
	if err != nil {
		logrus.Fatal(translations.StringValues.Errors.ErrorLoadingEnv)
	}
	DB, err = sql.Open(translations.StringValues.SqlDrivers.Postgres, os.Getenv(translations.StringValues.EnvironmentVariables.DatabaseUrl))
	if err != nil {
		logrus.Fatal(translations.StringValues.Errors.ErrorConnectingToDb)
	}

}

// Execute launches the CLI
func Execute() error {
	return rootCmd.Execute()
}
