package main

import (
	"fmt"

	"github.com/hilubabz/web-scraper-seo/crawler"
)

func main() {
	url := "https://utsargamanandhar.com.np"

	c := crawler.New()

	err := c.Crawl(url)
	if err!=nil{
		fmt.Println("Error:",err.Error())
	}
}
