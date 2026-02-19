package main

import (
	"log/slog"
	"net/http"

	"modulacms.com/components"

	modulacms "github.com/hegner123/modulacms/sdks/go"
)

// homeHandler returns an HTTP handler that fetches the home page from the CMS
// and renders it with templ components.
func homeHandler(client *modulacms.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := fetchHomePage(r.Context(), client)
		if err != nil {
			slog.Error("failed to fetch home page", "error", err)
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = components.HomePage(data).Render(r.Context(), w)
		if err != nil {
			slog.Error("failed to render home page", "error", err)
		}
	}
}
