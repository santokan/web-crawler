package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	rawURL := os.Args[1]

	fmt.Printf("starting crawl of: %s\n", rawURL)

	pages := make(map[string]int)

	crawlPage(rawURL, rawURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%s: %d\n", normalizedURL, count)
	}
}
