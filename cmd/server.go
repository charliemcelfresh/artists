package cmd

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "API server for artist data",
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
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logrus.Fatal("Error connecting to DB")
	}
}

func artists(w http.ResponseWriter, req *http.Request) {
	var a artist
	var artists []artist
	rows, err := DB.Query("SELECT * FROM artist;")
	defer rows.Close()
	if err != nil {
		logrus.Error("Cannot perform query")
	}
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
	http.HandleFunc("/artists", artists)
	http.ListenAndServe(":8080", nil)
}
