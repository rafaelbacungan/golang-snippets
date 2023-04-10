package webscraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func testWebScrape() {
	// Instantiate default collector
	//var articles []Articles
	//var pagesToScrape []string

	pageToScrape := "https://dev.to/search?q=golang"

	// current iteration
	//i := 1
	// max pages to scrape
	//limit := 5
	//
	// Initialize a Colly instance
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// iterating over the list of pagination links to implement the crawling logic
	c.OnHTML("a[href", func(e *colly.HTMLElement) {
		fmt.Println("I have reached the crawling logic...")
		fmt.Println("This is the text: ", e.Text)
		fmt.Println("This is the link: ", e.Attr("href"))
		// article := Articles{}
		// article.url = e.ChildAttr("a", "href")
		// fmt.Println("article url: ", article.url)
		// articles = append(articles, article)
	})

	c.Visit(pageToScrape)
}
