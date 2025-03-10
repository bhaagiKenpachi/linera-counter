package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var (
	currentValue int    = 0 // Initialize the counter value
	lineraPath   string     // Variable to hold the Linera execution path
)

func main() {
	port := flag.String("port", "3002", "Port to run the server on")
	lineraPath = flag.String("linera-path", "", "Path to the Linera executable") // New flag for Linera path
	flag.Parse()

	if *lineraPath == "" {
		log.Fatal("Linera path must be provided")
	}

	http.HandleFunc("/increment", handleIncrement)        // New endpoint for increment
	http.HandleFunc("/counter_value", handleCounterValue) // New endpoint for getting counter value

	log.Printf("Server starting on :%s", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// New endpoint to handle increment requests
func handleIncrement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the increment value from the request body
	var requestBody struct {
		Value int `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Increment the counter
	currentValue += requestBody.Value

	// Return success response
	response := map[string]interface{}{
		"status": "success",
		"data":   currentValue,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// New endpoint to return the current counter value
func handleCounterValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Return the current counter value
	response := map[string]interface{}{
		"status": "success",
		"value":  currentValue,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
