package crawler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/hilubabz/web-scraper-seo/seo"
)

type Crawler struct {
	Client *http.Client
	MaxDepth int
	MaxPages int
	
	mu sync.Mutex
	visited map[string]bool

	rateLimiter <-chan time.Time
}

func New() *Crawler {
	timer := time.NewTicker(500*time.Millisecond)
	return &Crawler{
		Client: &http.Client{},
		MaxDepth: 3,
		MaxPages: 100,
		visited: make(map[string]bool),
		rateLimiter: timer.C,
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

func (c *Crawler) Crawl(rawURL string) ([]crawlResult, error){
	baseURL, err := seo.NormalizeURL(rawURL)
	if err!=nil{
		return nil, err
	}
	jobs := make(chan crawlData, c.MaxPages)
	results := make(chan crawlResult)
	pages := make([]crawlResult,0)
	var wg sync.WaitGroup
	workerCount:=4
	wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go worker(i, c.Client, jobs, results, &wg, c.rateLimiter)
	}
	go func(){
		wg.Wait()
		close(results)
	}()
	if !c.markAsVisited(baseURL){
		return pages, nil
	}
	
	pendingJobs := 1
	jobs <- crawlData{
		Url: baseURL,
		Depth: 0,
	}

	for{
		result, ok := <-results
		if !ok{
			break
		}
		fmt.Printf("%s scraped successfully\n",result.Page.URL)
		pages = append(pages, result)
		pendingJobs--
		if result.Depth < c.MaxDepth{
			for _, link := range result.Links{
				if len(c.visited) >= c.MaxPages {
					break
				}
				if c.markAsVisited(link){
					jobs<-crawlData{
						Url: link,
						Depth: result.Depth+1,
					}
					pendingJobs++
				}
			}
		}
		if pendingJobs==0{
			close(jobs)
			break
		}
	}
	wg.Wait()
	fmt.Println("Pages crawled:",len(c.visited))
	return pages, nil
}
