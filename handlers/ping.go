package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JustSteveKing/go-ship/config"
)

// Response defines the return Response from the ping handler
type Response struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// PingHandler is the http.Handler for this request
type PingHandler struct {
	logger *log.Logger
}

// NewPingHandler creates a ping handler with the given logger
func NewPingHandler(logger *log.Logger) *PingHandler {
	return &PingHandler{logger}
}

// Handle will handle this request
func (pingHandler *PingHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	pingHandler.logger.Println("Ping Handler Dispatched")

	config := config.New()

	respondWithJSON(
		rw,
		http.StatusOK,
		&Response{
			Name:    config.App.Name,
			Version: config.App.Version,
		},
		"application/vnd.api+json",
	)
}

func respondWithJSON(rw http.ResponseWriter, code int, payload interface{}, contentType string) {
	response, _ := json.Marshal(payload)

	rw.Header().Set("Content-Type", contentType)
	rw.WriteHeader(code)
	rw.Write(response)
}
