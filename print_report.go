package main

import (
	"fmt"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("Report for %s\n", baseURL)
	fmt.Println("=============================")
	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		url := page.URL
		count := page.Count
		fmt.Printf("Found %d internal links to %s", count, url)
	}
}

func sortPages(pages map[string]int) []Page {
	pagesSlice := []Page{}
	for url, count := range pages {
		pagesSlice = append(pagesSlice, Page{URL: url, Count: count})
	}
	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].Count == pagesSlice[j].Count {
			return pagesSlice[i].URL < pagesSlice[j].URL
		}
		return pagesSlice[i].Count > pagesSlice[j].Count
	})
	return pagesSlice
}
