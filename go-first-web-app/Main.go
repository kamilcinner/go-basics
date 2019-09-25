package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert wed designer Kamil Cinner c:")
}

func main() {
	http.HandleFunc("/", indexHandler)      // Home Page
	http.HandleFunc("/about", aboutHandler) // About Page
	http.ListenAndServe(":8000", nil)
}
