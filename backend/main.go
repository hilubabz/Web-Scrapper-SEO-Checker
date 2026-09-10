package main

import (
	"fmt"

	"github.com/hilubabz/web-scraper-seo/crawler"
)

func main() {
	url := "https://utsargamanandhar.com.np"

	c := crawler.New()

	crawlData, err := c.Crawl(url)
	if err!=nil{
		fmt.Println("Error:",err.Error())
	}

	for _, page := range crawlData {
		fmt.Println("URL:", page.Page.URL)
		fmt.Println("Status:", page.Page.StatusCode)
		fmt.Println("Title:", page.Page.Title)
		fmt.Println("H1:", page.Page.H1Count)
		fmt.Println("Images:", page.Page.ImageCount)
		fmt.Println("---")
	}
}
