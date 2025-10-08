package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"time"
	"strings"
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
		fmt.Println(getHTML(args[0]))
	}
} 
