package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	response := &HealthResponse{
		Status: "UP",
	}
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Failed to marshal health response", "reason", err.Error())
		w.Write([]byte{})
	}
	w.Write(jsonResponse)
}
