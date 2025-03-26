package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Response represents a generic API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for assetlinks.json
	if strings.HasSuffix(r.URL.Path, "/.well-known/assetlinks.json") {
		// Set the content type
		w.Header().Set("Content-Type", "application/json")

		// Path to the assetlinks.json file
		assetlinksPath := filepath.Join("api", "public", ".well-known", "assetlinks.json")

		// Read the file
		content, err := os.ReadFile(assetlinksPath)
		if err != nil {
			http.Error(w, "Could not read assetlinks.json", http.StatusInternalServerError)
			return
		}

		// Write the content
		w.Write(content)
		return
	}

	// API route handling
	if strings.HasPrefix(r.URL.Path, "/api/") {
		handleAPI(w, r)
		return
	}

	// Default response for other paths
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// handleAPI processes API requests
func handleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the endpoint from the URL path
	path := strings.TrimPrefix(r.URL.Path, "/api")

	switch {
	case path == "" || path == "/":
		// API root endpoint
		response := Response{
			Status:  "success",
			Message: "Go DeepLink API is running",
			Data: map[string]string{
				"version": "1.0.0",
				"name":    "Go DeepLink API",
			},
		}
		json.NewEncoder(w).Encode(response)

	case path == "/health":
		// Health check endpoint
		response := Response{
			Status:  "success",
			Message: "API is healthy",
		}
		json.NewEncoder(w).Encode(response)

	default:
		// Handle unknown endpoints
		w.WriteHeader(http.StatusNotFound)
		response := Response{
			Status:  "error",
			Message: "Endpoint not found",
		}
		json.NewEncoder(w).Encode(response)
	}
}
