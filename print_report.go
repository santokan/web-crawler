package main

import "fmt"

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("Report for %s\n", baseURL)
	fmt.Println("=============================")
	sortPages()
	for page, count := range pages {
		fmt.Printf("Found %d internal links to %s", count, page)
	}
}

func (cfg *confg) sortPages() {
	// This is a implementation of sorting the pages
	// in descending order by count
  sortedPages := make([]string, len(pages))
  maxCount := 0
  for page, count := range cfg.pages {
    if count > maxCount {
      maxCount = count
    }

  }
