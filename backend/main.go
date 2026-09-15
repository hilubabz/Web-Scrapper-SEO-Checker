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

	mux := http.NewServeMux()

	mux.HandleFunc("/api/audits", handler.CreateAudit)

	server := audit.EnableCORS(mux)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}