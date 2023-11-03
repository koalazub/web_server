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

var auth, dbUrl string

func init() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file. Is it there?", err)
		return
	}

	auth = os.Getenv("libsql_auth_token")
	dbUrl = os.Getenv("libsql_url")

	if os.Getenv("libsql_auth_token") == "" || os.Getenv("libsql_url") == "" {
		slog.Error("Couldn't load database env variables. Is the file present?")
		return
	}

}

func InitDatabase() *sql.DB {
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %v: %v", dbUrl, err)
		os.Exit(1)
	}

	return db
}
