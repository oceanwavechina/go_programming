package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

// SitemapIndexST Location的数组
type SitemapIndexST struct {
	Locations []string `xml:"sitemap>loc"`
}

// NewsST comment
type NewsST struct {
	Titles    []string `xml:"url>news>title"`
	KeyWords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMapST fda
type NewsMapST struct {
	Keyword  string
	Location string
}

// NewsAggPageSt comment
type NewsAggPageSt struct {
	Title string
	News  map[string]NewsMapST
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
	var s SitemapIndexST
	var n NewsST
	NewsMap := make(map[string]NewsMapST)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(strings.Trim(Location, "\n "))
		if resp != nil {
			bytes, _ := ioutil.ReadAll(resp.Body)
			xml.Unmarshal(bytes, &n)
			//fmt.Println(n.Locations, n.Titles, Location)

			for idx := range n.Titles {
				NewsMap[n.Titles[idx]] = NewsMapST{n.KeyWords[idx], n.Locations[idx]}
			}
		}
		break
	}

	p := NewsAggPageSt{Title: "Amazing News Aggragator", News: NewsMap}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	err := t.Execute(w, p)
	fmt.Println(err)
}

// func main() {
// 	http.HandleFunc("/", IndexHandler)
// 	http.HandleFunc("/agg/", NewsAggHander)
// 	http.ListenAndServe(":8000", nil)
// }
