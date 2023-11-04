package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var tursoAuth, tursoUrl, databaseUrl string

func init() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file. Is it there?", err)
		return
	}

	tursoAuth = os.Getenv("turso_auth_token")
	tursoUrl = os.Getenv("turso_url")

	if os.Getenv("turso_auth_token") == "" || os.Getenv("turso_url") == "" {
		slog.Error("Couldn't load database env variables. Is the file present?")
		return
	}

}

func StartDatabase() *sql.DB {
	db := InitDatabase()
	err := InitTable(db)
	if err != nil {
		slog.Error("Couldn't start database")
		return nil
	}

	slog.Info("Database successfyully started")
	return db
}

func InitDatabase() *sql.DB {
	databaseUrl = fmt.Sprintf("libsql://%s?authToken=%s", tursoUrl, tursoAuth)

	db, err := sql.Open("libsql", databaseUrl)
	if err != nil {
		slog.Error("failed to open database", "url", tursoUrl, "error", err)
		os.Exit(1)
	}

	return db
}

func InitTable(db *sql.DB) error {
	_, err := db.Exec("create table if not exists spacex (id INT, name varchar(255), rocket varchar(255), links json, success boolean, flight_number int)")
	if err != nil {
		slog.Error("error initialising table", err)
		return err
	}

	slog.Info("table initialised")
	return nil
}

func Read(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		slog.Error("couldn't make the query ", err)
		return nil, err
	}

	return rows, nil
}
