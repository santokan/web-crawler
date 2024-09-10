package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("couldn't parse base URL")
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("couldn't parse current URL")
		return
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalized URL: %v\n", err)
	}

	if _, ok := pages[normalizedURL]; ok {
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1

	fmt.Println("crawling", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - couldn't get HTML - %v\n", err)
		return
	}

	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - get URLs from HTML: %v\n", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
