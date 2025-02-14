package api

import (
	"errors"
	"go-db-project/internal/store"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/redis/go-redis/v9"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGet(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		url, err := store.GetFullURL(r.Context(), code)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				sendJSON(w, Response{Error: "Code not found"}, http.StatusNotFound)
			}
			slog.Error("Failed to get full url", "error", err)
			sendJSON(w, Response{Error: "Something went wrong"}, http.StatusInternalServerError)
			return
		}
		sendJSON(w, Response{Data: getShortenedURLResponse{FullURL: url}}, http.StatusOK)
	}
}
