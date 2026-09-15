# Backend - SEO Web Analyzer

The backend service for the SEO Web Analyzer is built in **Go (Golang)**. It exposes a REST API that triggers a concurrent web crawler and runs an SEO analysis rules engine against the fetched pages.

## Architecture & Implementation

The backend is modularized into several key packages:

- **`main.go`**: The entry point of the application. It initializes the crawler, sets up the HTTP handlers, configures CORS, and starts the server.
- **`audit/`**: Contains the API handlers and service layer for managing audit requests (`/api/audits`).
- **`crawler/`**: Implements a highly concurrent web crawler utilizing a **Worker Pool pattern**. It spawns multiple worker goroutines that concurrently pull and process URLs from shared Go channels. This ensures extremely fast, parallel scraping while respecting rate limits, avoiding race conditions (via Mutexes), and capping the crawl based on `maxDepth` and `maxPages`.
- **`seo/`**: The core SEO analysis engine. 
  - Parses HTML to extract titles, meta tags, headers, links, and images.
  - `rules.go`: Contains the business logic and rules for identifying SEO issues (e.g., missing tags, bad lengths, indexability).
  - `score.go`: Calculates a severity-weighted score out of 100 based on the detected issues.

## Running the Backend

Ensure you have Go installed on your system.

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Run the server:
   ```bash
   go run .
   ```

The server will start on `http://localhost:8080` by default. You can override the port by setting the `PORT` environment variable.
