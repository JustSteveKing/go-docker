package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPintHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req, error := http.NewRequest("GET", "/", nil)

	if error != nil {
		t.Fatal(error)
	}

	// Create a new Logger instance to pass into our Handler
	logger := log.New(os.Stdout, "go-ship-service ", log.LstdFlags)

	// create our handler
	pingHandler := NewPingHandler(logger)

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler.Handle)

	// Call ServeHTTP directly
	handler.ServeHTTP(rr, req)

	// Test status code is correct
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"name":"Go App","version":"v1.0"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
