package cmd

import (
	"artists/internal/translations"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	loadArtistsCmd *cobra.Command = &cobra.Command{
		Use:   translations.StringValues.Commands.LoadArtists.Use,
		Short: translations.StringValues.Commands.LoadArtists.Short,
		RunE:  LoadArtistsCommandRunner,
	}
	RunCount int
)

func init() {
	loadArtistsCmd.Flags().IntVar(&RunCount, translations.StringValues.Flags.LoadArtistsCountFlag.Name, 0, translations.StringValues.Flags.LoadArtistsCountFlag.Description)
	rootCmd.AddCommand(loadArtistsCmd)
}

func LoadArtistsCommandRunner(cmd *cobra.Command, args []string) error {
	for i := 0; i < RunCount; i++ {
		err := loadArtist(args[i])
		if err != nil {
			logrus.Error(translations.StringValues.Errors.ErrorPerformingQuery)
		}
	}
	return nil
}

func loadArtist(email string) error {
	_, err := DB.Exec("INSERT INTO artist (email) VALUES ($1)", email)
	return err
}
