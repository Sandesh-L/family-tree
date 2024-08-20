package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {

		// Log incoming request
		fmt.Printf("Received request: %s %s \n", r.Method, r.URL.Path)

		// Set response header to allow CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Set response header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Send a JSON message
		fmt.Fprint(w, `{"message": "Hello from Go backend!"}`)
	})

	// Start the go server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
