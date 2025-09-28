package main
import (
	"log"
	"github.com/PuerkitoBio/goquery"
	"strings"
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
		return ""	
	}
	return strings.TrimSpace(first_par.Text())
}
