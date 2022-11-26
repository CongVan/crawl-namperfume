package main

import "sync"

func main() {
	// NewCrawlBrand()
	var wg sync.WaitGroup

	NewCrawlProductDetail()
	wg.Wait()
}
