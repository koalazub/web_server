[![Go](https://github.com/koalazub/web_server/actions/workflows/go.yml/badge.svg)](https://github.com/koalazub/web_server/actions/workflows/go.yml)

> Go version:  1.21

## Installation

Best to use `nix`, but can use `go cli`. 

For `nix`:
	if developing, it's as simple as `nix develop`. Seek further information fro the provided `flake.nix` file.

If using `go` cli:
	Just run `go install ./...`, followed by `go run`. This should allow you to run the server. But you might run into issues if you don't have a `.env` file ready. 

For the database: 
	Ensure you have a local or `Turso` database to connect to. Seek information below

## Goal

The main aim was to get better at stitching services together. I leaned on eliminating the annoyances of setting up a frontend and wanted to focus on the underpinnings of a backend/database relationship.

The renders made in the client are templates, the backend server is a `go` server using `gorilla/mux` which handles routing and updates from the `Turso` database using `libsql`. That's the bulk of the tech stack

# Building a Basic Web Server with Go's net/http Package

### Handling Routes
There is a `handler` function that establishes all of the required routes for the fetch

## Testing
Minimal testing - making sure there's a database connection along with a couple of little checks. This was more of a discovery piece of work.

# Database

Using `Turso` and `libsql` for the creation of the database along with some logic to create the table. This will mean that you will either: 
	- an `.env` file with connections that match the global constant variables for connections to be made
		So:

			- [ ] `host_addr`
			- [ ] `host_port`
			- [ ] `turso_auth_token`
			- [ ] `turso_url`

	- a `libsql/sqlite` database that's created locally and that connects to the correct address(default is usually `127.0.0.1:8080`)


## Building
## Conclusion
This Go project efficiently demonstrates creating a web server using the standard \`net/http\` package. It offers a han
ds-on introduction to web server design using native Go libraries without external dependencies.
