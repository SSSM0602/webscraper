package main
import (
	"net/url"
	"fmt"
	"log"
	"path"
)

func normalizeURL(url_in string) (string, error) {
	parsed, err := url.Parse(url_in)
	if err != nil {
		log.Fatal(err)
	}
	final := fmt.Sprintf("%s%s", parsed.Host, path.Clean(parsed.Path))
	return final, err
}
