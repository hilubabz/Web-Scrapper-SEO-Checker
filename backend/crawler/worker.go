package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hilubabz/web-scraper-seo/seo"
)

type crawlData struct {
	Url   string
	Depth int
}

type crawlResult struct{
	Page seo.AnalysisResult
	Depth int
	Links []string
	Issues []seo.SEOIssue
	Score int
	Err error
}

func worker(id int, client *http.Client, jobs <-chan crawlData, results chan<- crawlResult, wg *sync.WaitGroup, rateLimiter <-chan time.Time) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d scraping %s\n", id, job.Url)
		<-rateLimiter
		res, err := client.Get(job.Url)
		if err!=nil{
			results<-crawlResult{
				Depth: job.Depth,
				Err: err,
			}
			continue
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		if err != nil{
			results<-crawlResult{
				Depth: job.Depth,
				Err: err,
			}
			continue
		}
		baseUrl, err := url.Parse(job.Url)
		if err!=nil{
			results<-crawlResult{
				Depth: job.Depth,
				Err: err,
			}
			continue
		}
		data, err := seo.Analyze(baseUrl, doc, res.StatusCode)
		issues := seo.CheckRules(data)
		score := seo.CalculateScore(issues)
		if err!=nil{
			results<-crawlResult{
				Depth: job.Depth,
				Err: err,
			}
			continue
		}
		links := seo.ExtractInternalLinks(doc, baseUrl)
		results <- crawlResult{
			Page: *data,
			Depth: job.Depth,
			Links: links,
			Issues: issues,
			Score: score,
		}
	}
}