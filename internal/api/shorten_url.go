package api

import (
	"encoding/json"
	"go-db-project/internal/store"
	"log/slog"
	"net/http"
	"net/url"
)

type shortenURLResponse struct {
	Code string `json:"code"`
}

type shortenURLRequest struct {
	URL string `json:"url"`
}

func handlePost(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shortenURLRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "Invalid request"}, http.StatusUnprocessableEntity)
			return
		}
		_, err := url.Parse(body.URL)
		if err != nil {
			sendJSON(w, Response{Error: "Invalid URL"}, http.StatusBadRequest)
			return
		}
		code, err := store.SaveShortenedURL(r.Context(), body.URL)
		if err != nil {
			slog.Error("Failed to save shortened url", "error", err)
			sendJSON(w, Response{Error: "Something went wrong"}, http.StatusInternalServerError)
			return
		}
		sendJSON(w, Response{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
