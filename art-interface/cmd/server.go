package main

import (
	art_interf "art/art-interface/pkg/interface"
	"fmt"
	"html/template"
	"net/http"
)

// newlines are spaces - maybe adjust my decode code to convert?

type PageData struct {
	Output string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("HTTP/1.1 200 OK")
		http.ServeFile(w, r, "index.html")

	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		// Get the value of the "input" field from the form
		userInput := r.Form.Get("input")
		isMalformed := false
		var decodedOutput string

		// Decode the input
		decodedOutput, isMalformed = art_interf.DecodeInput(userInput)

		if isMalformed {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		} else {
			fmt.Println("HTTP/1.1 202 Accepted")

			// Create a PageData struct with the decoded value
			decodedData := PageData{
				Output: decodedOutput,
			}

			// Render the HTML template with the PageData
			newHtmlTemplate := template.Must(template.ParseFiles("index.html"))
			newHtmlTemplate.Execute(w, decodedData)

		}

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	// Register the indexHandler function to handle requests to the root URL ("/")
	http.HandleFunc("/", indexHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
