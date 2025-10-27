# Webscraper

A Go-based web scraping tool that crawls web pages, extracts content and images, normalizes URLs, and generates CSV reports.

## Features
- Crawl multiple pages concurrently.
- Extract text and image URLs from HTML.
- Normalize and deduplicate URLs.
- Export extracted data to a CSV file.
- Unit tests included for core functions.

## Project Structure
```
.
├── crawling.go           # Core crawler logic
├── extract_content.go    # HTML content extraction
├── normalize_url.go      # URL normalization utilities
├── csv_report.go         # CSV export functionality
├── main.go               # Entry point
├── go.mod                # Go module configuration
└── *_test.go             # Unit tests
```

## Installation
```bash
git clone https://github.com/SSSM0602/webscraper.git
cd webscraper
go mod tidy
```

## Usage
```bash
go run main.go <base_url> <max_concurrency> <page_limit>
```
- **base_url** – Starting URL for crawling  
- **max_concurrency** – Number of concurrent workers  
- **page_limit** – Maximum number of pages to crawl  

Example:
```bash
go run main.go https://example.com 5 100
```

## Output
- CSV file containing extracted URLs, titles, and image links.

## Testing
```bash
go test ./...
```

## Requirements
- Go 1.20+
- Internet connection for crawling

## License
MIT
