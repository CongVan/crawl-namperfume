package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func NewCrawlBrand() {
	// Create file
	fileName := "brands.csv"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Thumbnail", "Slug"})

	// Crawler page
	c := colly.NewCollector()

	c.OnHTML(".desktop-brands-section .all-brands .page-template__brand__description .brand-name", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("div"))
		e.ForEach("div.brand-item", func(i int, h *colly.HTMLElement) {
			fmt.Println("section ", i, h.ChildText("div"))

		})

	})
	c.Visit("https://namperfume.net/pages/thuong-hieu-a-z")
}
