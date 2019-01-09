package main

import (
	"fmt"
	"net/http"
)

// IndexHandler fasda
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Who, Go is neat!")
}

// AboutHandler fasda
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is about page!")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about/", AboutHandler)
	http.ListenAndServe(":8000", nil)
}
