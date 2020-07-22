package cmd

import (
	"database/sql"
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "artist",
	Short: "artist service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("subcommand required")
	},
}

var DB *sql.DB

// Execute launches the CLI
func Execute() error {
	return rootCmd.Execute()
}
