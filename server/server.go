package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
)

var indexPage = `<!DOCTYPE html>
	<html>
		<body>
			<h1 style="text-align:center;"> Oooohh look at disssss!</h1>
			<p style="text-align:center;"> Oooohh look at disssss!</p>
		</ body>
	</ html>
	`

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

func New() *Server {
	return &Server{users: make(map[string]UserInfo)}
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexPage))
}

func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	switch r.Method {
	case http.MethodPost, http.MethodPut:
		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Couldn't read req body:", err, 3)
			w.WriteHeader(http.StatusInternalServerError) // 500
			return
		}
		defer r.Body.Close()

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			slog.Error("Couldn't unmarshal reqt body: %v ", err)
			w.WriteHeader(http.StatusBadRequest) // 400
			return
		}

		slog.Info("Created User:", u.Name, 0)
		s.users[u.Name] = UserInfo{
			starsign:             u.Starsign,
			diabolicaltendencies: u.Diabolicaltendencies,
		}
	case http.MethodDelete:
		delete(s.users, name)
		slog.Info("Deleting user:", name)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed) // 415
	}
}

func (s *Server) HandleReadUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Fetch from query strin
		name := r.URL.Query().Get("name")
		u, ok := s.users[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ret := User{
			Name:                 name,
			Starsign:             u.starsign,
			Diabolicaltendencies: u.diabolicaltendencies,
		}
		msg, err := json.Marshal(ret)
		if err != nil {
			slog.Error("couldn't marshall:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(msg)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
