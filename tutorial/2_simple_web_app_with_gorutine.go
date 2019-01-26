package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var wgHttp sync.WaitGroup

// sitemapIndexST Location的数组
type sitemapIndexST struct {
	Locations []string `xml:"sitemap>loc"`
}

// newsST comment
type newsST struct {
	Titles    []string `xml:"url>news>title"`
	KeyWords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// newsMapST fda
type newsMapST struct {
	Keyword  string
	Location string
}

// newsAggPageSt comment
type newsAggPageSt struct {
	Title string
	News  map[string]newsMapST
}

// indexHandler fasda
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1> Hey, there !<H1>")
	fmt.Fprintf(w, "<H1> go is fast !<H1>")
	fmt.Fprintf(w, "<H1> ... and simple !<H1>")

	fmt.Fprintf(w, `You can use multi-line
with emphasize-quot
	and caution the space`)
}

func newsRoutine(c chan newsST, Location string) {
	defer wgHttp.Done()
	var n newsST

	resp, _ := http.Get(strings.Trim(Location, "\n "))
	if resp != nil {
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		//fmt.Println(n.Locations, n.Titles, Location)
		resp.Body.Close()
		c <- n
	}
}

// newsAggHander web page with template
func newsAggHander(w http.ResponseWriter, r *http.Request) {
	var s sitemapIndexST

	newsMap := make(map[string]newsMapST)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()

	queue := make(chan newsST, 100)

	for _, Location := range s.Locations {
		wgHttp.Add(1)
		go newsRoutine(queue, strings.Trim(Location, "\n "))
	}

	for elem := range queue {
		for idx := range elem.Titles {
			newsMap[elem.Titles[idx]] = newsMapST{elem.KeyWords[idx], elem.Locations[idx]}
		}
		break
	}

	wgHttp.Wait()
	close(queue)

	p := newsAggPageSt{Title: "Amazing News Aggragator", News: newsMap}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	err := t.Execute(w, p)
	fmt.Println("error:", err)
}

// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/agg/", newsAggHander)
// 	http.ListenAndServe(":8000", nil)
// }
