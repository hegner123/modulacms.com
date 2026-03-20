package main

import (
	"log/slog"
	"net/http"

	"modulacms.com/components"
)

func playgroundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err := components.PlaygroundPage().Render(r.Context(), w)
		if err != nil {
			slog.Error("failed to render playground page", "error", err)
		}
	}
}
