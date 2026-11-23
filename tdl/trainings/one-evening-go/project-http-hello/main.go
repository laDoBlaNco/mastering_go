package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", yourFunction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func yourFunction(w http.ResponseWriter, r *http.Request) {
	// handler function body
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hello, ", name)
}
