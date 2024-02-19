package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// PageData holds the data to be rendered in the HTML template
type PageData struct {
	Output string
}

// indexHandler handles requests to the root URL ("/")
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// Serve the HTML form
		http.ServeFile(w, r, "index.html")

		// Check if the request method is POST
	} else if r.Method == "POST" {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		// Get the value of the "input" field from the form
		userInput := r.Form.Get("input")

		// Process the user input (e.g., perform some computation)
		// For this example, let's simply echo back the input
		output := userInput + " understood"

		// Create a PageData struct with the output
		data := PageData{
			Output: output,
		}

		// Render the HTML template with the PageData
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Register the indexHandler function to handle requests to the root URL ("/")
	http.HandleFunc("/", indexHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
