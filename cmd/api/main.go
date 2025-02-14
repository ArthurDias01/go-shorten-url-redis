package main

import (
	"go-db-project/internal/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("Starting server...", "version", "1.0.0")
	if err := run(); err != nil {
		slog.Error("failed to run", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)
	s := http.Server{
		Addr:                         ":8080",
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  time.Minute,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
