package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/lytics/logrus"
	"github.com/spf13/cobra"
)

var (
	loadArtistsCmd *cobra.Command = &cobra.Command{
		Use:   "loadArtists",
		Short: "load artists to database",
		RunE:  LoadArtistsCommandRunner,
	}
	RunCount int
)

func init() {
	loadArtistsCmd.Flags().IntVar(&RunCount, "count", 0, "number of artists to load")
	rootCmd.AddCommand(loadArtistsCmd)

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logrus.Fatal("Error connecting to DB")
	}
}

func LoadArtistsCommandRunner(cmd *cobra.Command, args []string) error {
	for i := 0; i < RunCount; i++ {
		err := loadArtist(args[i])
		if err != nil {
			logrus.Error(fmt.Sprintf("Error loading email: %s, error: %s", args[i], err))
		}
	}
	return nil
}

func loadArtist(email string) error {
	_, err := DB.Exec("INSERT INTO artist (email) VALUES ($1)", email)
	return err
}
