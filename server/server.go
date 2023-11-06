package server

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	a "github.com/koalazub/web-server/api"
	"golang.org/x/exp/slog"
)

var addr, port string

func init() {
	loadEnv()
}

type Server struct {
	DB *sql.DB
}

// Server spins up with an initalised database
func New(db *sql.DB) *Server {
	return &Server{
		DB: db,
	}
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		slog.Error("Internal server error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError) // 500
		return
	}

	err = tmpl.Execute(w, "")
	if err != nil {
		slog.Error("Internal server error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError) // 500
		return
	}
}

func RunServer(db *sql.DB) {

	server_env := addr + port

	r := mux.NewRouter()
	handlers(db, r)
	s := &http.Server{
		Addr:           addr + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server: %v", server_env)
	slog.Error("Error launching serve: ", s.ListenAndServe())
}

func handlers(db *sql.DB, r *mux.Router) {
	srv := New(db)
	lsrv := a.New()
	r.HandleFunc("/", srv.HandleIndex)
	r.HandleFunc("/launches/upcoming", lsrv.HandleGetLaunches)
	r.HandleFunc("/launches/custom", lsrv.HandleGetCustomLaunchData)
	r.HandleFunc("/launches/database", srv.HandleGetDatabaseLaunches)
}

// Reaches into Server to get access to DB
func (s *Server) HandleGetDatabaseLaunches(w http.ResponseWriter, r *http.Request) {
	launches, err := a.GetDatabaseLaunches(s.DB)
	if err != nil {
		slog.Error("Could't get launch info", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}

	t, err := template.ParseFiles("templates/launches.templ")
	if err != nil {

		slog.Error("couldn't read from template", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, map[string]interface{}{"Launches": launches})
	if err != nil {
		slog.Error("couldn't execute file from template", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file. Is it there?", err)
		return
	}

	addr = os.Getenv("host_addr")
	port = os.Getenv("host_port")

	if os.Getenv("host_addr") == "" || os.Getenv("host_port") == "" {
		slog.Error("Couldn't load env variables. Is the file present?")
		return
	}
}
