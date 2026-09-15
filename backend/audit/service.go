package audit

import (
	"fmt"
	"net/url"

	"github.com/hilubabz/web-scraper-seo/crawler"
)

type Service struct {
	crawler *crawler.Crawler
}

func NewService(c *crawler.Crawler) *Service {
	return &Service{
		crawler: c,
	}
}

func (s *Service) Run(startURL string, maxDepth int, maxPages int) (crawler.AuditResult, error) {
	parsedURL, err := url.Parse(startURL)
	if err != nil {
		return crawler.AuditResult{}, fmt.Errorf("invalid URL")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return crawler.AuditResult{}, fmt.Errorf(
			"URL must use http or https",
		)
	}

	if parsedURL.Host == "" {
		return crawler.AuditResult{}, fmt.Errorf(
			"URL must contain a host",
		)
	}

	pages, err := s.crawler.Crawl(startURL, maxDepth, maxPages)
	if err != nil {
		return crawler.AuditResult{}, err
	}

	result := crawler.BuildAuditResult(startURL, pages)

	return result, nil
}