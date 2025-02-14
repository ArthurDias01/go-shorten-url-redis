package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		url, ok := db[code]
		if !ok {
			sendJSON(w, Response{Error: "Not found"}, http.StatusNotFound)
			return
		}
		sendJSON(w, Response{Data: getShortenedURLResponse{FullURL: url}}, http.StatusOK)
	}
}
