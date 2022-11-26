package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	title       string
	brand_name  string
	imgs        []string
	raw_sku     string
	origins     string
	year        string
	fragrants   string
	styles      string
	description string
	gender      string
	is_new      bool
}

func NewCrawlProductDetail() {
	jsonFile, err := os.Open("product_links.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var productLinks []string = []string{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &productLinks)

	var size = len(productLinks)
	fmt.Println(size)
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("#detail-product .product-sex-type", func(h *colly.HTMLElement) {
		a, _ := h.DOM.Html()
		fmt.Println("OKOKO", a)
	})

	c.OnHTML("#product", func(h *colly.HTMLElement) {
		var product Product
		product.title = h.ChildText(".product-title h1")
		product.imgs = h.ChildAttrs(".product-gallery .product-gallery__large__slider__slide:not(.slick-cloned) img", "src")
		product.brand_name = h.ChildText("#product-brand a")

		// Attributes
		h.ForEach("#tab-detail .product-attribute-list .product-attribute-list__item", func(i int, h *colly.HTMLElement) {
			dd := h.ChildText("dd")
			dt := h.ChildText("dt")
			switch dt {
			case "Mã hàng":
				product.raw_sku = dd
			case "Xuất xứ":
				product.origins = dd
			case "Năm phát hành":
				product.year = dd
			case "Nhóm hương":
				product.fragrants = dd
			case "Phong cách":
				product.styles = dd
			}
		})

		product.description, _ = h.DOM.Find("#tab-detail .description-productdetail").Html()
		product.gender = h.ChildText(".product-sex-type")
		g := h.ChildAttr("#detail-product span.product-sex-type", "data-gender")
		fmt.Println("gender ", g, product.gender)
		product.is_new = h.ChildText(".product-meta .product-badge") == "New"

		h.ForEach(".product-variant-select ul li", func(i int, h *colly.HTMLElement) {
			// variantTitle := h.ChildText(".product-variant-item__title")
			// fmt.Print("variant", variantTitle)
		})

		// fmt.Printf("%+v\n", product)
	})

	for i := 0; i < size; i++ {
		// linkItem := productLinks[i]
		if i > 0 {
			return
		}
		c.Visit("https://namperfume.net/products/nuoc-hoa-nam-versace-eros")
	}
	c.Wait()
}
