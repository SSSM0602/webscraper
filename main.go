package main

import (
	"fmt"
	"os"
	"net/url"
	"sync"
)
func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}  
	rawBaseURL := args[0]
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
		concurrencyControl: make(chan struct{}, 1), // limit concurrent crawls
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(parsedBaseURL.String())

	cfg.wg.Wait()

	for url, count := range cfg.pages {
		fmt.Printf("%s %d\n", url, count)
	}
} 
