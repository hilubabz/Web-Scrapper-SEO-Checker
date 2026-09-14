package main

import (
	"log"
	"net/http"

	"github.com/hilubabz/web-scraper-seo/audit"
	"github.com/hilubabz/web-scraper-seo/crawler"
)

func main() {
	c := crawler.New()

	service := audit.NewService(c)
	handler := audit.NewHandler(service)

	http.HandleFunc("/api/audits", handler.CreateAudit)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}