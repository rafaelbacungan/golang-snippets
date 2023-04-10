package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	//val := Chaser(2, 3)
	//fmt.Println(val)
	testWebScrape()
}

func Chaser(speed, time int) int {

	var maxDistance int

	if time == 1 {
		return speed * 2
	}

	for i := 0; i < time; i++ {
		if i == time-1 {
			maxDistance += speed * 2
		} else {
			maxDistance += speed
		}
	}

	return maxDistance
}

func testWebScrape() {
	// Instantiate default collector
	pageToScrape := "https://dev.to"
	// pageToScrape := "https://www.technologyreview.com/topic/artificial-intelligence/"

	// Initialize a Colly instance
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// iterating over the list of pagination links to implement the crawling logic
	c.OnHTML("div.crayons-story", func(e *colly.HTMLElement) {
		tags := e.ChildText("a.crayons-tag")
		if containsTag(tags, "#react") {
			fmt.Println("Title: ", e.ChildText("h2.crayons-story__title"))
			fmt.Println("Author: ", e.ChildText("a.crayons-story__secondary"))
			fmt.Println("Link: ", pageToScrape+e.ChildAttr("a", "href"))
		}
	})

	c.Visit(pageToScrape)
}

func containsTag(tags string, tag string) bool {
	if strings.Contains(tags, tag) {
		return true
	} else {
		return false
	}
}
