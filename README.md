# SEO Web Analyzer

A full-stack web application designed to instantly analyze technical SEO issues across websites. It features a concurrent backend crawler and a modern, interactive frontend dashboard to review scores, page metrics, and actionable recommendations.

## Features

- **Concurrent Web Crawler**: High-performance crawler built in Go that explores websites up to a configurable maximum depth and page count.
- **Robust SEO Rules Engine**: Analyzes critical on-page SEO factors including Title tags, Meta descriptions, H1 usage, Image Alt text, Canonical links, and Robots meta tags.
- **Dynamic Scoring System**: Calculates an overall SEO score weighted by the severity of the issues found (Errors, Warnings, and Info).
- **Interactive UI**: A sleek, glassmorphic dashboard with staggered animations and responsive design.

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: Next.js (React, TypeScript), Tailwind CSS
- **UI Libraries**: Framer Motion (animations), Lucide React (icons)

## Directory Structure

- `/backend`: The Go API server, crawler, and SEO analysis engine.
- `/frontend`: The Next.js React application and UI components.

## Getting Started

Please see the respective README files in the `backend` and `frontend` directories for detailed instructions on how to run each service.
