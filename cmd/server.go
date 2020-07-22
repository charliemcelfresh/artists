package cmd

import (
	"artists/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the artist server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

