package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPageSt struct {
	Title string
	News  string
}

// IndexHandler fasda
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1> Hey, there !<H1>")
	fmt.Fprintf(w, "<H1> go is fast !<H1>")
	fmt.Fprintf(w, "<H1> ... and simple !<H1>")

	fmt.Fprintf(w, `You can use multi-line
with emphasize-quot
	and caution the space`)
}

// NewsAggHander web page with template
func NewsAggHander(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPageSt{Title: "Amazing News Aggragator", News: "some news"}
	t, _ := template.ParseFiles("basic_templating.html")
	err := t.Execute(w, p)
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/agg/", NewsAggHander)
	http.ListenAndServe(":8000", nil)
}
