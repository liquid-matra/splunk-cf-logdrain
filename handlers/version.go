package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func VersionHandler(version string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &VersionResponse{
			Version: version,
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
}
