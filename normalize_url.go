package main
import (
	"net/url"
	"log"
	"strings"
)

func normalizeURL(url_in string) (string, error) {
	parsed, err := url.Parse(url_in)
	if err != nil {
		log.Fatal(err)
	}
	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}
	parsed.Host = strings.TrimSuffix(parsed.Host, ".")
	return parsed.String(), err
}
