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
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			if renderErr := components.NotFoundPage().Render(r.Context(), w); renderErr != nil {
				slog.Error("failed to render 404", "error", renderErr)
			}
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
		w.Header().Set("Cache-Control", "no-cache")
		if err := page.Render(r.Context(), w); err != nil {
			slog.Error("failed to render page", "slug", slug, "error", err)
		}
	}
}
