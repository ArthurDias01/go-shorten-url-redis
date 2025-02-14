package main

import (
	"go-db-project/internal/api"
	"go-db-project/internal/store"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	store := store.NewStore(rdb)
	handler := api.NewHandler(store)
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
