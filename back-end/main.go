package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type customAttributes struct {
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type Person struct {
	ID         int              `json:"id"`
	NodeType   string           `json:"nodeType"`
	Label      string           `json:"label"`
	FirstName  string           `json:"firstName"`
	MiddleName string           `json:"middleName"`
	LastName   string           `json:"lastName"`
	Parents    []int            `json:"parents"`
	Spouses    []int            `json:"spouses"`
	Children   []int            `json:"children"`
	Attributes customAttributes `json:"attributes"`
}

func main() {
	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {

		msg := Person{
			ID:         1,
			NodeType:   "Person",
			Label:      "Ancestor",
			FirstName:  "John",
			MiddleName: "d",
			LastName:   "Doe",
			Parents:    []int{2, 3},
			Spouses:    []int{4},
			Children:   []int{5, 6},
			Attributes: customAttributes{
				Age:    45,
				Gender: "Male",
			},
		}

		// Set response header to allow CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Set response header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Log incoming request
		fmt.Printf("Received request: %s %s \n", r.Method, r.URL.Path)

		jsonResponse, err := json.Marshal(msg)

		if err != nil {
			http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)
	})

	// Start the go server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
