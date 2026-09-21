package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler parses and displays form data submitted via POST request.
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler serves the GET request for the /hello endpoint.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure strict path matching for the route
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// Restrict endpoint to GET requests only
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	// Serve static files from the ./static directory at root
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Register HTTP route handlers
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the web server on port 8080
	fmt.Println("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
