package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	modulacms "github.com/hegner123/modulacms/sdks/go"
)

func main() {
	baseURL := os.Getenv("CMS_BASE_URL")
	if baseURL == "" {
		slog.Error("CMS_BASE_URL environment variable is required")
		os.Exit(1)
	}

	apiKey := os.Getenv("CMS_API_KEY")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	client, err := modulacms.NewClient(modulacms.ClientConfig{
		BaseURL: baseURL,
		APIKey:  apiKey,
	})
	if err != nil {
		slog.Error("failed to create CMS client", "error", err)
		os.Exit(1)
	}

	mux := newMux(client)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		slog.Info("server starting", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
