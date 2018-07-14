package main

import (
	"fmt"
	"log"

	"github.com/puerkitoBio/goquery"
)

func postScrape() {
	doc, err := goquery.NewDocument("http://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("#main article .entry-title").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})
}

func main() {
	postScrape()
}
