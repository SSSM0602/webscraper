package main

import (
	"fmt"
	"os"
	"net/url"
	"sync"
	"strconv"
)
func main() {
	args := os.Args[1:]
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else if len(args) < 3 {
		fmt.Println("usage: go run . [URL] [maxConcurrency] [maxPages]")
		os.Exit(1)
	}  
	rawBaseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid max concurrency")
	}
	pageCap, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid max pages")
	}
	fmt.Println("starting crawl")

	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Invalid base URL: %v\n", err)
		os.Exit(1)
	}

	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency), // limit concurrent crawls
		wg:                 &sync.WaitGroup{},
		maxPages:			pageCap,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(parsedBaseURL.String())

	cfg.wg.Wait()

	for url, count := range cfg.pages {
		fmt.Printf("%s %d\n", url, count)
	}
} 
