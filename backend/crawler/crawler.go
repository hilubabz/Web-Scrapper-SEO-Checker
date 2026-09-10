package crawler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/hilubabz/web-scraper-seo/seo"
)

type Crawler struct {
	Client *http.Client
	MaxDepth int
	MaxPages int
	
	mu sync.Mutex
	visited map[string]bool
}

func New() *Crawler {
	return &Crawler{
		Client: &http.Client{},
		MaxDepth: 3,
		MaxPages: 100,
		visited: make(map[string]bool),
	}
}

func (c *Crawler) markAsVisited(URL string) bool{
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.visited[URL]{
		return false
	}
	c.visited[URL]=true
	return true
}

func (c *Crawler) Crawl(rawURL string) error{
	baseURL, err := seo.NormalizeURL(rawURL)
	if err!=nil{
		return err
	}
	jobs := make(chan crawlData)
	results := make(chan crawlResult)
	var wg sync.WaitGroup
	workerCount:=4
	wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go worker(i, c.Client, jobs, results, &wg)
	}
	go func(){
		wg.Wait()
		close(results)
	}()
	if !c.markAsVisited(baseURL){
		return nil
	}
	go func(){
		jobs <- crawlData{
			Url: baseURL,
			Depth: 0,
		}
		close(jobs)
	}()
	for result := range results{
		fmt.Println("Finished crawling:",result.Url)
		for _, link := range result.Links{
			if c.markAsVisited(link){
				fmt.Println("Discovered",link)
			}
		}
	}
	fmt.Println("Pages crawled:",len(c.visited))
	return nil
}
