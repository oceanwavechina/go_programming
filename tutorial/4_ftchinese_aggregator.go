package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// ItemST fadf
type ItemST struct {
	ItemList []RssInfoST `xml:"channel>item"`
}

// RssInfoST comment
type RssInfoST struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

// FTNewsAggPageSt comment
type FTNewsAggPageSt struct {
	Title string
	News  []RssInfoST
}

// FTAggregator FT中文网的新闻聚合
func FTAggregator(w http.ResponseWriter, r *http.Request) {
	var items ItemST
	var newsList []RssInfoST

	RssList := [...]string{
		"http://www.ftchinese.com/rss/feed",
		"http://www.ftchinese.com/rss/news",
		"http://www.ftchinese.com/rss/diglossia",
		"http://www.ftchinese.com/rss/column/007000005",
		"http://www.ftchinese.com/rss/column/007000004",
		"http://www.ftchinese.com/rss/column/007000007",
		"http://www.ftchinese.com/rss/column/007000002",
		"http://www.ftchinese.com/rss/lifestyle",
		"http://www.ftchinese.com/rss/letter",
		"http://www.ftchinese.com/rss/column/007000012 ",
	}

	for _, rssLink := range RssList {
		resp, _ := http.Get(rssLink)
		if resp != nil {
			bytes, _ := ioutil.ReadAll(resp.Body)
			xml.Unmarshal(bytes, &items)
			for _, item := range items.ItemList {
				newsList = append(newsList, item)
			}
		}
	}

	p := FTNewsAggPageSt{Title: "Amazing News Aggragator", News: newsList}
	t, _ := template.ParseFiles("ftchinese_template.html")
	err := t.Execute(w, p)
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", FTAggregator)
	http.ListenAndServe(":8000", nil)
}
