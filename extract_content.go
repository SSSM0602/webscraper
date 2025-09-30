package main
import (
	"log"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/url"
	"fmt"
)

func getH1FromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	h1 := doc.Find("h1")
	if h1.Length() == 0 {
		return ""
	}
	return strings.TrimSpace(h1.Text())
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	first_par := doc.Find("main").Find("p").First()
	if first_par.Length() == 0 {
		first_par = doc.Find("p").First()
	}

	if first_par.Length() == 0 {
		return ""
	}
	return strings.TrimSpace(first_par.Text())
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	var urls []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
        // For each '<a href>' it finds, it will run this function.
		href, exists := s.Attr("href")
		if exists {
			var full_url string
			if href == baseURL.String() {
				full_url = href
			} else {
				full_url = fmt.Sprintf("%s%s", baseURL.String(), href)
			}
			urls = append(urls, full_url)
		}
    })
	return urls, err
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	var imgs []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			var full_url string
			if src == baseURL.String() {
				full_url = src
			} else {
				full_url = fmt.Sprintf("%s%s", baseURL.String(), src)
			}
			imgs = append(imgs, full_url)
		}
	})
	return imgs, err
}

type PageData struct {
	URL string
	H1 string
	FirstParagraph string
	OutgoingLinks []string
	ImageURLs []string
}

func extractPageData(html, pageURL string) PageData {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}
	}
	images, err := getImagesFromHTML(html, baseURL)
	links, err := getURLsFromHTML(html, baseURL)
	res_data := PageData {
		URL: pageURL,
		H1: getH1FromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks: links,
		ImageURLs: images,
	}
	return res_data
}
