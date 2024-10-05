package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"splunk-cf-logdrain/handlers"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatalf("could not create test health handler request: %v", err)
	}
	rec := httptest.NewRecorder()
	handlers.HealthHandler(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("health handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	var gotResponse handlers.HealthResponse
	if err := json.Unmarshal(b, &gotResponse); err != nil {
		t.Fatalf("could not marshal Health response: %v", err)
	}

	expectedResponse := handlers.HealthResponse{Status: "UP"}

	if expectedResponse.Status != gotResponse.Status {
		t.Errorf("health handler returned wrong status: got %v want %v", gotResponse.Status, expectedResponse.Status)
	}
}
