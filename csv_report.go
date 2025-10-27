package main

import (
	"encoding/csv"
	"os"
	"log"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string {"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}); err != nil {
		return err
	}
	
	for _, page := range pages {
		record := []string{
			page.URL,
			page.H1,
			page.FirstParagraph,
			strings.Join(page.OutgoingLinks, ";"),
			strings.Join(page.ImageURLs, ";"),
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
