package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function for handling requests to the root URL "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // Respond with "Hello, World!" to any request
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
