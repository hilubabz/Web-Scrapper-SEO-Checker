package main

import (
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	err := http.ListenAndServe(port, server)
	if err != nil {
		log.Fatal(err)
	}
}
