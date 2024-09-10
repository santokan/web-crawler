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

	const maxConcurrency = 3
	cfg, err := configure(rawURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure %v", err)
		return
	}

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%s: %d\n", normalizedURL, count)
	}
}
