package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

// 爬虫框架
func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), ".html") {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		if strings.Contains(fmt.Sprintf("%+v", r.URL), ".html") {
			fmt.Println("Visiting", r.URL)
		}
	})

	c.Visit("https://12seaa.com/vodtypehtml/26.html")
}
