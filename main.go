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
func formHandler(w http.ResponseWriter, r *http.Request){

}


// handles main route
func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/home" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
}
