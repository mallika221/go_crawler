package main

import (
	"fmt"
	"go_crawler/html_parser"
	"go_crawler/link_parser"
	"go_crawler/normalize_url"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentUrl string, pages map[string]int) {
	currentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentUrl, err)
		return
	}

	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	// skipping other websites
	if currentUrl.Hostname() != baseUrl.Hostname() {
		return
	}

	normalizedUrl, err := normalize_url.NormalizeUrl(rawCurrentUrl)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	if _, visited := pages[normalizedUrl]; visited {
		pages[normalizedUrl]++
		return
	}

	pages[normalizedUrl] = 1

	fmt.Printf("crawling %s\n", rawCurrentUrl)

	htmlBody, err := html_parser.GetHTML(rawCurrentUrl)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := link_parser.GetURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		crawlPage(rawBaseURL, nextURL, pages)
	}
}
