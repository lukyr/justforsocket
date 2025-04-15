package main

import (
	"encoding/json"
	"net/http"

	"github.com/continue-team/riot/socket"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	newSocket := socket.GetInstance()

	// Define a simple handler for the root path
	r.HandleFunc("/", newSocket.HandleWebSocket)
	r.HandleFunc("/broadcast", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			StreamKey string      `json:"streamKey"`
			Message   interface{} `json:"message"`
		}

		// Decode JSON request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.StreamKey == "" {
			http.Error(w, "StreamKey is required", http.StatusBadRequest)
			return
		}

		// Broadcast message
		newSocket.BroadcastMessage(req.StreamKey, req.Message)

		// Send response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Broadcast sent successfully",
		})
	}).Methods("POST")

	newSocket.StartPingRoutine()

	// Start the server
	http.ListenAndServe(":9999", r)
}
