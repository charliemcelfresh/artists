package main

import (
	"database/sql"
	"encoding/json"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	DB *sql.DB
)

type artist struct {
	ArtistID int64 `json:"artist_id"`
	Email string `json:"email"`
}

func init() {
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

func main() {
	http.HandleFunc("/artists", artists)
	http.ListenAndServe(":8080", nil)
}
