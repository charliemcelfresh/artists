package cmd

import (
	"artists/internal/translations"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   translations.StringValues.Commands.Server.Use,
	Short: translations.StringValues.Commands.Server.Short,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverCmdRunner()
		return nil
	},
}

type artist struct {
	ArtistID int64  `json:"artist_id"`
	Email    string `json:"email"`
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func artists(w http.ResponseWriter, req *http.Request) {
	var a artist
	var artists []artist
	rows, err := DB.Query("SELECT * FROM artist;")
	if err != nil {
		logrus.Error(translations.StringValues.Errors.ErrorPerformingQuery)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&a.ArtistID, &a.Email)
		if err != nil {
			logrus.Error(err)
		}
		artists = append(artists, a)
		logrus.Println(a)
	}
	json.NewEncoder(w).Encode(artists)
}

func serverCmdRunner() {
	http.HandleFunc(translations.StringValues.Paths.Artists, artists)
	http.ListenAndServe(translations.StringValues.Ports.EightyEighty, nil)
}
