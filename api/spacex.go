package api

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"

	d "github.com/koalazub/web-server/database"
	"golang.org/x/exp/slog"
)

type CustomLaunch struct {
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

// For Turso db
type CustomLaunchData struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Rocket       string `json:"rocket"`
	Links        Reddit `json:"reddit"`
	Success      bool   `json:"success"`
	FlightNumber int    `json:"flight_number"`
}

// New creates a new server for Launches
func New() *CustomLaunch {
	return &CustomLaunch{launches: make(map[string]CustomLaunchData)}
}

func (s *CustomLaunch) HandleGetLaunches(w http.ResponseWriter, r *http.Request) {
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

func (s *CustomLaunch) HandleGetCustomLaunchData(w http.ResponseWriter, r *http.Request) {
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

type Launch struct {
	date_utc string
	// Fairings string                 `json:"fairings"`
	Links  map[string]interface{} `json:"links"`
	Reddit map[string]interface{} `json:"reddit"`
}

// All the crud operations
func GetSpaceXLaunches() ([]Launch, error) {

	res, err := http.Get("https://api.spacexdata.com/v5/launches/upcoming")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var launches []Launch
	err = json.NewDecoder(res.Body).Decode(&launches)
	if err != nil {
		return nil, err
	}

	return launches, nil
}

func GetDatabaseLaunches(db *sql.DB) ([]CustomLaunchData, error) {
	qry := "SELECT id, name, rocket, success, flight_number from spacex"

	rows, err := d.Read(db, qry)
	if err != nil {
		slog.Error("Couldn't query the database for launches: ", err)
		return nil, err
	}

	defer rows.Close()

	var launches []CustomLaunchData
	for rows.Next() {
		var l CustomLaunchData
		if err := rows.Scan(&l.ID, &l.Name, &l.Rocket, &l.Success, &l.FlightNumber); err != nil {
			slog.Error("Error scanning row. Are you querying correctly?\n", err)
			return nil, err
		}

		launches = append(launches, l)
	}

	if err = rows.Err(); err != nil {
		slog.Error("Error iterating rows: ", err)
		return nil, err
	}

	return launches, nil

}
