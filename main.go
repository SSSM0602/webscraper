package main

import (
	"fmt"
	"os"
)
func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else {
		fmt.Println("starting crawl")
		pages := make(map[string]int, 10)
		crawlPage(args[0], args[0], pages)
		for url, count := range pages {
			fmt.Printf("%s %d\n", url, count)
		}
		fmt.Println()
	}
} 
