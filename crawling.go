package main
import (
	"fmt"
	"net/http"
	"net/url"
	"io"
	"time"
	"strings"
	"log"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages		   int
}

func getHTML(rawURL string) (string, error) {
	cli := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if len(cfg.pages) >= cfg.maxPages {
		return false
	}

	// Already visited
	if _, ok := cfg.pages[normalizedURL]; ok {
		return false
	}

	// First visit — initialize a new PageData entry
	cfg.pages[normalizedURL] = PageData{
		URL: normalizedURL,
	}
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()
	if cfg.concurrencyControl != nil {
		cfg.concurrencyControl <- struct{}{}
		defer func() { <-cfg.concurrencyControl }()
	}

	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %s\n", rawCurrentURL)
		return
	}

	// Skip URLs outside the base domain
	if !strings.EqualFold(cfg.baseURL.Hostname(), currURL.Hostname()) {
		return
	}

	normRaw, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing URL: %s\n", rawCurrentURL)
		return
	}

	isFirst := cfg.addPageVisit(normRaw)
	if !isFirst {
		return // Already crawled
	}

	fmt.Printf("Fetching HTML from: %s\n", normRaw)

	currHTML, err := getHTML(normRaw)
	if err != nil {
		fmt.Printf("Error fetching HTML from URL: %s\n", normRaw)
		log.Fatal(err)
		return
	}

	currPageURLs, err := getURLsFromHTML(currHTML, currURL)
	if err != nil {
		fmt.Printf("Error fetching links from current page: %s\n", normRaw)
		log.Fatal(err)
		return
	}

	for _, url := range currPageURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}




