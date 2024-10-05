package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"splunk-cf-logdrain/handlers"
	"testing"
)

func TestVersionHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/version", nil)
	if err != nil {
		t.Fatalf("could not create test version handler request: %v", err)
	}
	rec := httptest.NewRecorder()
	versionHandlerFunc := handlers.VersionHandler("0.0.0")
	versionHandlerFunc(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("version handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read /version response body: %v", err)
	}

	var gotVersionResponse handlers.VersionResponse

	if err := json.Unmarshal(b, &gotVersionResponse); err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	expectedVersionResponse := handlers.VersionResponse{Version: "0.0.0"}

	if gotVersionResponse.Version != expectedVersionResponse.Version {
		t.Fatalf("version response did not match: got %v want %v", gotVersionResponse.Version, expectedVersionResponse.Version)
	}

}
