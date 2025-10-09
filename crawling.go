package main
import (
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io"
	"time"
	"strings"
)

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

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	baseHost := baseURL.Hostname()
	currHost := currURL.Hostname()

	if !strings.EqualFold(baseHost, currHost) {
		return 
	}

	normRaw, err := normalizeUrl(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing URL: %s\n", rawCurrentURL)
		return
	}
	if val, ok := pages[normRaw]; ok {
		pages[normRaw] += 1
		return
	} else {
		pages[normRaw] = 1
		fmt.Printf("Fetching html from: %s\n", normRaw)

		currHTML, err := getHTML(normRaw)
		if err != nil {
			fmt.Printf("Error fetching HTML from URL: %s\n", normRaw)
			return
		}

		currPageURLs, err := getURLsFromHTML(currHTML)
		if err != nil {
			fmt.Printf("Error fetching links from current page: %s\n", normRaw)
			return
		}
		for _, url := range currPageURLs {
			crawlPage(rawBaseURL, url, pages)
		}
	}
}


