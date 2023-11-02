package api

import (
	"io"
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

type LaunchServer struct {
	launches map[string]LaunchSpecs
}

// just the basics for now and then we'll see if we can elaborate more.
// Specs need to be updated from map[string]interface{} at some point
type LaunchSpecs struct {
	Fairings string                 `json:"fairings"`
	Links    map[string]interface{} `json:"links"`
	Reddit   map[string]interface{} `json:"reddit"`
}

// New creates a new server for Launches
func New() *LaunchServer {
	return &LaunchServer{launches: make(map[string]LaunchSpecs)}
}

func (s *LaunchServer) HandleGetLaunches(w http.ResponseWriter, r *http.Request) {

	err := GetSpaceXLaunches()
	if err != nil {
		slog.Error("couldn't get space launches", err)
		return
	}
	file, err := os.Open("writtendata.json")
	if err != nil {
		slog.Error("Couldn't read the file. You sure it's there? ", err)
		return
	}
	defer file.Close()
	_, err = io.Copy(w, file)
	if err != nil {
		slog.Error("Issue with copying file", err)
		return
	}

}

// All the crud operations
func GetSpaceXLaunches() error {
	res, err := http.Get("https://api.spacexdata.com/v5/launches/upcoming")
	if err != nil {
		slog.Error("Something happened when requesting the launches", err)
		return err
	}
	slog.Info("couldn't read", res)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Couldn't read body", err)
		return err
	}

	err = os.WriteFile("writtendata.json", body, 0644)
	if err != nil {
		slog.Error("Error writing to file", err)
		return err
	}
	return nil
}
