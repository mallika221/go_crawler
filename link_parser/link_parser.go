package link_parser

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(htmlBody string, rawBaseUrl string) ([]string, error) {
	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL %w", err)
	}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML %w", err)
	}

	fmt.Println(baseUrl, doc)

	var urls = []string{}
	var traverseNodes func(*html.Node)

	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("could not parse %v", err)
						continue
					}
					resolvedUrl := baseUrl.ResolveReference(href)
					urls = append(urls, resolvedUrl.String())

				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}
	traverseNodes(doc)

	return urls, nil

}
