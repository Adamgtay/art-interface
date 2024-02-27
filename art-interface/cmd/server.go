package main

import (
	art_interf "art/art-interface/pkg/interface"
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Output string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("HTTP/1.1 200 OK")

		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		// Get the value of the "input" field from the form
		userInput := r.Form.Get("input")
		isMalformed := false
		var output string

		mode := r.Form.Get("mode") // read radio button selection

		// Encode or decode based on the selected mode
		if mode == "decode" {
			output, isMalformed = art_interf.DecodeInput(userInput)
		} else {
			output = art_interf.EncodeInput(userInput)
		}

		if isMalformed {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		} else {
			fmt.Println("HTTP/1.1 202 Accepted")

			// Create a PageData struct with the decoded value
			processedData := PageData{
				Output: output,
			}
			// HTML template with new PageData
			newHtmlTemplate, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			newHtmlTemplate.Execute(w, processedData)
		}

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	// Register the indexHandler function to handle requests to the root URL ("/")
	http.HandleFunc("/", indexHandler)
	fmt.Println("Copy and paste into web browser -> \033[38;5;205mhttp://localhost:8080/\033[0m")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
