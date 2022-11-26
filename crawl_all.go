package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Data struct {
	Title string
	Link  string
}

func GetLimitPage() int {
	c := colly.NewCollector()
	var limitPage int

	c.OnHTML("#pagination", func(e *colly.HTMLElement) {
		number := e.ChildText("a:nth-last-child(2)")

		// string to int
		i, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		limitPage = i
	})

	c.Visit("https://namperfume.net/collections/all")

	return limitPage
}

func NewCrawlProduct() {
	// Create file
	fileName := "result.csv"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Title", "Link"})

	// Crawler page
	c := colly.NewCollector()

	limitPage := GetLimitPage()

	c.OnHTML(".pro-loop .product-block", func(e *colly.HTMLElement) {
		title := e.ChildText(".box-pro-detail > .pro-name")
		link := e.ChildAttr(".box-pro-detail > .pro-name a", "href")

		//fmt.Printf("Title: %s \nLink: %s \n", title, link)
		writer.Write([]string{
			title,
			fmt.Sprintf("https://namperfume.net%s", link),
		})
	})

	for i := 1; i <= limitPage; i++ {
		fullURL := fmt.Sprintf("https://namperfume.net/collections/all?page=%d", i)
		c.Visit(fullURL)
	}
}
