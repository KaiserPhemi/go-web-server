package main

// Imports
import (
	"fmt"
	"log"
	"net/http"
)

// Main function
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/home", homeHandler)

	fmt.Printf("Starting server on port 7000\n")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}

// Handles form route
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error: %v", err)
		return
	}

	fmt.Fprintf(w, "Request Successfully parsed")
	name := r.FormValue("name")

	fmt.Print(name)
}

// handles main route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// checks for route
	if r.URL.Path != "/home" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// checks if method is not GET
	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Welcome Home")
}
