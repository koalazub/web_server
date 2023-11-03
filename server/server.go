package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/koalazub/web-server/api"
	"golang.org/x/exp/slog"
)

var addr, port string

func init() {
	loadEnv()
}

type Server struct {
	users map[string]UserInfo //key -> username
}

type User struct {
	Name                 string `json:"name"`
	Starsign             string `json:"starsign"`
	Diabolicaltendencies int    `json:"diabolicaltendencies"`
}

type UserInfo struct {
	starsign             string
	diabolicaltendencies int
}

// New generates a new server with empty users
func New() *Server {
	return &Server{users: make(map[string]UserInfo)}
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

func RunServer() {

	server_env := addr + port

	r := mux.NewRouter()
	srv := New()
	lsrv := api.New()

	r.HandleFunc("/", srv.HandleIndex)
	r.HandleFunc("/launches/upcoming", lsrv.HandleGetLaunches)
	r.HandleFunc("/launches/upcoming/all", lsrv.HandleGetCustomLaunchData)
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
