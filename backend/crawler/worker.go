package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/hilubabz/web-scraper-seo/seo"
)

type crawlData struct {
	Url   string
	Depth int
}

type crawlResult struct{
	Url string
	Depth int
	Links []string
}

func worker(id int, client *http.Client, jobs <-chan crawlData, results chan<- crawlResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		res, err := client.Get(job.Url)
		if err!=nil{
			fmt.Println("Error:",err.Error())
			continue
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		if err != nil{
			fmt.Println("Error:",err.Error())
			continue
		}
		fmt.Println("Parsing",job.Url)
		baseUrl, err := url.Parse(job.Url)
		if err!=nil{
			fmt.Println("Error:",err.Error())
			continue
		}
		data, err := seo.Analyze(baseUrl, doc, res.StatusCode)
		if err!=nil{
			fmt.Println("Error:",err.Error())
			continue
		}
		data.PrintResult()
		links := seo.ExtractInternalLinks(doc, baseUrl)
		results <- crawlResult{
			Url: job.Url,
			Depth: job.Depth,
			Links: links,
		}
	}
}