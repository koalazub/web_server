package api

import (
	"encoding/json"
	"html/template"
	"net/http"

	"golang.org/x/exp/slog"
)

type LaunchServer struct {
	launches map[string]CustomLaunchData
}

type Reddit struct {
	Campaign interface{} `json:"campaign"`
	Launch   interface{} `json:"launch"`
	Media    interface{} `json:"media"`
}
type Links struct {
	Reddit Reddit
}

// just the basics for now and then we'll see if we can elaborate more.
// Specs need to be updated from map[string]interface{} at some point
type CustomLaunchData struct {
	Name         string `json:"name"`
	Rocket       string `json:"rocket"`
	Links        Reddit `json:"reddit"`
	Success      bool   `json:"success"`
	FlightNumber int    `json:"flight_number"`
}

// New creates a new server for Launches
func New() *LaunchServer {
	return &LaunchServer{launches: make(map[string]CustomLaunchData)}
}

func (s *LaunchServer) HandleGetLaunches(w http.ResponseWriter, r *http.Request) {
	launches, err := GetSpaceXLaunches()
	if err != nil {
		slog.Error("Could't get launches", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(launches); err != nil {
		slog.Error("Failed to send launches", err)
		http.Error(w, "internal server error", http.StatusInternalServerError) // 500
		return
	}
}

func (s *LaunchServer) HandleGetCustomLaunchData(w http.ResponseWriter, r *http.Request) {
	servErr := func(err error) error {
		if err != nil {
			slog.Error("couldn't read body", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return nil
	}

	launches, err := GetSpaceXLaunches()
	servErr(err)

	tmpl, err := template.ParseFiles("templates/launches.templ")
	servErr(err)

	err = tmpl.Execute(w, map[string]interface{}{
		"Launches": launches,
	})
	servErr(err)
}

// ###############
type Launch struct {
	date_utc string
	// Fairings string                 `json:"fairings"`
	Links  map[string]interface{} `json:"links"`
	Reddit map[string]interface{} `json:"reddit"`
}

// All the crud operations
func GetSpaceXLaunches() ([]CustomLaunchData, error) {

	res, err := http.Get("https://api.spacexdata.com/v5/launches/upcoming")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var launches []CustomLaunchData
	err = json.NewDecoder(res.Body).Decode(&launches)
	if err != nil {
		return nil, err
	}

	return launches, nil
}

// ############
