package cmd

import (
	"artists/internal/tasks"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(populateArtistCommand)
}

var populateArtistCommand = &cobra.Command{
	Use:   "populate_artist",
	Short: "Populate artist table with fake data",
	Run: func(cmd *cobra.Command, args []string) {
		count, err := strconv.Atoi(args[0])
		if err != nil {
			logrus.Fatal(err)
		}
		tasks.PopulateArtist(count)
	},
}
