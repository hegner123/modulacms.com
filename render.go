package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"

	"modulacms.com/components"

	modulacms "github.com/hegner123/modulacms/sdks/go"
)

// pageHandler returns an HTTP handler that fetches a page by slug from the CMS
// and renders it with templ components.
func pageHandler(client *modulacms.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.URL.Path
		data, err := fetchPage(r.Context(), client, slug)
		if err != nil {
			slog.Error("failed to fetch page", "slug", slug, "error", err)
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		var page templ.Component
		switch data.Page.Type {
		case "documentation":
			page = components.DocsPage(data)
		default:
			page = components.HomePage(data)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := page.Render(r.Context(), w); err != nil {
			slog.Error("failed to render page", "slug", slug, "error", err)
		}
	}
}
