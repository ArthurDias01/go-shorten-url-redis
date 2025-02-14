package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, resp Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Error marshalling response:", "error", err)
		sendJSON(w, Response{Error: "Internal server error"}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	if _, err := w.Write(data); err != nil {
		// fmt.Println("Error writing response:", err)
		slog.Error("Error writing response:", "error", err)
		return
	}
}
