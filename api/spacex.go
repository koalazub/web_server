package api

import (
	"encoding/json"
	"fmt"
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

	ptyJson, err := json.MarshalIndent(launches, "", "") // 4 space indent
	if err != nil {
		slog.Error("Failed to prettify this", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s\n", ptyJson)
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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(launches)
	servErr(err)

	pj, err := json.MarshalIndent(launches, "", "    ")
	servErr(err)

	fmt.Printf("%s\n", pj)
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

	checkErr := func(err error) error {
		if err != nil {
			return err
		}
		return nil
	}

	res, err := http.Get("https://api.spacexdata.com/v5/launches/upcoming")
	checkErr(err)

	defer res.Body.Close()

	var launches []CustomLaunchData
	err = json.NewDecoder(res.Body).Decode(&launches)
	checkErr(err)

	return launches, err
}

// ############
