package seo

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type AnalysisResult struct {
	URL                *url.URL
	StatusCode         int
	Title              string
	TitleLength        int
	MetaDescription    string
	HasMetaDescription bool
	H1Count            int
	H2Count            int
	ImageCount         int
	ImagesWithoutAlt   int
	LinkCount          int
}

func (anRes *AnalysisResult) PrintResult(){
	fmt.Println("Status Code:", anRes.StatusCode)
	fmt.Println("Title:", anRes.Title)
	fmt.Println("Title Length:", anRes.TitleLength)
	if anRes.HasMetaDescription{
		fmt.Println("Meta Description:", anRes.MetaDescription)
	} else{
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
	metaDesc, exists := doc.Find(`meta["description"]`).Attr("content")
	images := doc.Find("img")
	imagesWithoutAlt:=0
	images.Each(func(index int, accessor *goquery.Selection){
		alt, exist:=accessor.Attr("alt")
		if !exist || alt!=""{
			imagesWithoutAlt++
		}
	})
	result := &AnalysisResult{
		URL: rawUrl,
		StatusCode: statusCode,
		Title: title.Text(),
		TitleLength: len(title.Text()),
		MetaDescription: metaDesc,
		HasMetaDescription: exists,
		H1Count: doc.Find("h1").Length(),
		H2Count: doc.Find("h2").Length(),
		ImageCount: images.Length(),
		ImagesWithoutAlt: imagesWithoutAlt,
		LinkCount: doc.Find("a").Length(),
	}
	return result, nil
}