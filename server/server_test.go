package server

import (
	d "github.com/koalazub/web-server/database"
	"testing"
)

func TestNew(t *testing.T) {
	db := d.InitDatabase()
	got := New(db)
	if got.DB == nil {
		t.Errorf("New() did not initialise a new server with a db")
	}
}
