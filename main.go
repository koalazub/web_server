package main

import (
	d "github.com/koalazub/web-server/database"
	s "github.com/koalazub/web-server/server"
)

func main() {
	db := d.StartDatabase()
	defer db.Close()

	s.RunServer(db)
}
