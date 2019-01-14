package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
	<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<sitemap>
	<loc>
	https://www.washingtonpost.com/news-sitemaps/politics.xml
	</loc>
	</sitemap>
	<sitemap>
	<loc>
	https://www.washingtonpost.com/news-sitemaps/opinions.xml
	</loc>
	</sitemap>
	<sitemap>
	<loc>
*/

/*
// SitemapIndex Location的数组
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location 包含url的最内层节点
type Location struct {
	Loc string `xml:"loc"`
}

// String 使用fmt.Println()时，如果当前结构中有定义String()方法时，会默认调用此方法返回值用于输出
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
*/

// SitemapIndexST Location的数组
type sitemapIndexST struct {
	Locations []string `xml:"sitemap>loc"`
}

type newsST struct {
	Titles    []string `xml:"url>news>title"`
	KeyWords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMapST fda
type newsMapST struct {
	Keyword  string
	Location string
}

func agg() {
	var s sitemapIndexST
	var n newsST
	newsMap := make(map[string]newsMapST)

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
				newsMap[n.Titles[idx]] = newsMapST{n.KeyWords[idx], n.Locations[idx]}
			}
		}
		break
	}

	for idx, data := range newsMap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}

// func main() {
// 	agg()
// }
