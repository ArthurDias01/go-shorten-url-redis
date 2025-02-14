package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type shortenURLResponse struct {
	Code string `json:"code"`
}

type shortenURLRequest struct {
	URL string `json:"url"`
}

func handlePost(db map[string]string) http.HandlerFunc {
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
		code := generateCode(db)
		db[code] = body.URL
		sendJSON(w, Response{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
