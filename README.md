[![Go](https://github.com/koalazub/web_server/actions/workflows/go.yml/badge.svg)](https://github.com/koalazub/web_server/actions/workflows/go.yml)

# Building a Basic Web Server with Go's net/http Package
## Introduction
This repository showcases how to create a simple web server using Go's standard \`net/http\` package. It represents an 
understanding of essential web server functionalities using native Go libraries.

## Implementation

### Setting up the Server
Import the \`net/http\` package and use the \`http.HandleFunc\` and \`http.ListenAndServe\` functions to create a basic
server.

\`\`\`go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\\n")
	})

	http.ListenAndServe(":3000", nil)
}
\`\`\`

### Handling Routes
Define different handlers for various routes to provide dynamic responses.

\`\`\`go
http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
})
\`\`\`

## Testing
Run the server and navigate to \`http://localhost:3000\` in your browser to see a 'Hello, World!' response.

## Conclusion
This Go project efficiently demonstrates creating a web server using the standard \`net/http\` package. It offers a han
ds-on introduction to web server design using native Go libraries without external dependencies.
