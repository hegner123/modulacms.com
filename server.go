package main

import (
	"net/http"

	modulacms "github.com/hegner123/modulacms/sdks/go"
)

// newMux builds the HTTP mux with all routes wired.
func newMux(client *modulacms.Client) *http.ServeMux {
	mux := http.NewServeMux()

	// Static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Home page
	mux.HandleFunc("GET /{$}", homeHandler(client))

	// Playground (static — no CMS API call)
	mux.HandleFunc("GET /playground", playgroundHandler())

	return mux
}
