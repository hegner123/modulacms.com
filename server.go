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

	// Favicon redirect
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/favicon.svg", http.StatusMovedPermanently)
	})

	// Playground (static — no CMS API call)
	mux.HandleFunc("GET /playground", playgroundHandler())

	// CMS pages — catch-all, must be last
	mux.HandleFunc("GET /", pageHandler(client))

	return mux
}
