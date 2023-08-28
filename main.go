package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/web-server/server"
	"golang.org/x/exp/slog"
)

func main() {
	r := mux.NewRouter()
	srv := server.New()
	addr := ":8008"
	r.HandleFunc("/", srv.HandleIndex)
	r.HandleFunc("/users/{name}", srv.HandleReadUsers)
	r.HandleFunc("/users/create", srv.HandleCreateUser)
	s := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server: %v", addr)
	log.Fatal(s.ListenAndServe())
	slog.Error("Error launching serve: ", s.ListenAndServe())
}
