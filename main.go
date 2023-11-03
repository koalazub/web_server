package main

import (
	d "github.com/koalazub/web-server/database"
	s "github.com/koalazub/web-server/server"
)

func main() {
	s.RunServer()
	d.StartDatabase()
}
