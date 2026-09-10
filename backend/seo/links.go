package seo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractInternalLinks(doc *goquery.Document, baseURL *url.URL) []string {
	links := make([]string, 0)

	doc.Find("a[href]").Each(func(_ int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")

		if !exists || href == "" {
			return
		}

		if strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "javascript:") ||
			strings.HasPrefix(href, "mailto:") {
			return
		}

		parsedURL, err := url.Parse(href)
		if err != nil {
			return
		}

		absoluteURL := baseURL.ResolveReference(parsedURL)

		if absoluteURL.Host != baseURL.Host {
			return
		}

		absoluteURL.Fragment = ""

		links = append(links, absoluteURL.String())
	})

	return links
}