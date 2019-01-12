package main

import (
	"fmt"
	"net/http"
)

// IndexHandler fasda
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1> Hey, there !<H1>")
	fmt.Fprintf(w, "<H1> go is fast !<H1>")
	fmt.Fprintf(w, "<H1> ... and simple !<H1>")

	fmt.Fprintf(w, `You can use multi-line
with emphasize-quot
	and caution the space`)
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
