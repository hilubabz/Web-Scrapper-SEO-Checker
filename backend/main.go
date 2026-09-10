package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hilubabz/web-scraper-seo/crawler"
	"github.com/hilubabz/web-scraper-seo/sql"
	"github.com/hilubabz/web-scraper-seo/sql/generated"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DATABASE_URL")

	ctx := context.Background()

	pool, err := sql.Connect(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	
	queries := sql.NewStore(pool)

	project, err := queries.CreateProject(ctx,generated.CreateProjectParams{
		Name: "Example Website",
		BaseUrl: "https://example.com",
	})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Project created",project.ID)

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
		fmt.Println()
		fmt.Println("Issues:")
		for _, issue := range page.Issues{
			fmt.Printf("[%s] %s\n", issue.Severity, issue.Message)
		}
		fmt.Println("Score:",page.Score)
		fmt.Println("---")
	}
}
