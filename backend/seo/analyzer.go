package seo

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type AnalysisResult struct {
	URL string `json:"url"`

	StatusCode         int    `json:"statusCode"`
	Title              string `json:"title"`
	TitleLength        int    `json:"titleLength"`
	MetaDescription    string `json:"metaDescription"`
	HasMetaDescription bool   `json:"hasMetaDescription"`
	H1Count            int    `json:"h1Count"`
	H2Count            int    `json:"h2Count"`
	ImageCount         int    `json:"imageCount"`
	ImagesWithoutAlt   int    `json:"imagesWithoutAlt"`
	LinkCount          int    `json:"linkCount"`

	CanonicalURL       string `json:"canonicalUrl"`
	HasCanonical       bool   `json:"hasCanonical"`
	RobotsMeta         string `json:"robotsMeta"`
	HasRobotsMeta      bool   `json:"hasRobotsMeta"`
	InternalLinkCount  int    `json:"internalLinkCount"`
	ExternalLinkCount  int    `json:"externalLinkCount"`
	MetaDescriptionLen int    `json:"metaDescriptionLen"`
}

func (anRes *AnalysisResult) PrintResult() {
	fmt.Println("Status Code:", anRes.StatusCode)
	fmt.Println("Title:", anRes.Title)
	fmt.Println("Title Length:", anRes.TitleLength)
	if anRes.HasMetaDescription {
		fmt.Println("Meta Description:", anRes.MetaDescription)
	} else {
		fmt.Println("No meta description found")
	}
	fmt.Println("H1 Count:", anRes.H1Count)
	fmt.Println("H2 Count:", anRes.H2Count)
	fmt.Println("Images:", anRes.ImageCount)
	fmt.Println("Images Without Alt:", anRes.ImagesWithoutAlt)
	fmt.Println("Total Links:", anRes.LinkCount)
}

func Analyze(rawUrl *url.URL, doc *goquery.Document, statusCode int) (*AnalysisResult, error) {
	title := doc.Find("title")
	metaDesc, exists := doc.Find(`meta[name="description"]`).Attr("content")
	images := doc.Find("img")
	imagesWithoutAlt := 0
	images.Each(func(index int, accessor *goquery.Selection) {
		alt, exist := accessor.Attr("alt")
		if !exist || alt == "" {
			imagesWithoutAlt++
		}
	})
	canonicalURL, hasCanonical := doc.Find(`link[rel="canonical"]`).Attr("href")
	robotsMeta, hasRobotsMeta := doc.Find(`meta[name="robots"]`).Attr("content")
	descriptionLength := len(metaDesc)
	internalLinks := 0
	externalLinks := 0

	doc.Find("a[href]").Each(func(_ int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")

		if !exists || href == "" {
			return
		}

		parsedLink, err := rawUrl.Parse(href)
		if err != nil {
			return
		}

		if parsedLink.Hostname() == rawUrl.Hostname() {
			internalLinks++
		} else {
			externalLinks++
		}
	})
	result := &AnalysisResult{
		URL:                rawUrl.String(),
		StatusCode:         statusCode,
		Title:              title.Text(),
		TitleLength:        len(title.Text()),
		MetaDescription:    metaDesc,
		HasMetaDescription: exists,
		H1Count:            doc.Find("h1").Length(),
		H2Count:            doc.Find("h2").Length(),
		ImageCount:         images.Length(),
		ImagesWithoutAlt:   imagesWithoutAlt,
		LinkCount:          doc.Find("a").Length(),
		CanonicalURL:       canonicalURL,
		HasCanonical:       hasCanonical,
		RobotsMeta:         robotsMeta,
		HasRobotsMeta:      hasRobotsMeta,
		MetaDescriptionLen: descriptionLength,
		InternalLinkCount:  internalLinks,
		ExternalLinkCount:  externalLinks,
	}
	return result, nil
}
