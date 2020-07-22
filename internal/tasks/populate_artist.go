package tasks

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"syreclabs.com/go/faker"
)

var (
	DB *sql.DB
)

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

func PopulateArtist(count int) {
	tx, err := DB.Begin()
	if err != nil {
		logrus.Error(err)
	}
	for i := 0; i < 10; i++ {
		statement := "INSERT INTO artist (email) VALUES ($1)"
		_, err = tx.Exec(statement, faker.Internet().Email())
		if err != nil {
			logrus.Error(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
	}
}

