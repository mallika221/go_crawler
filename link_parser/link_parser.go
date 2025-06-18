package linkparser

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(htmlBody string, rawBaseUrl string) (string, error) {
	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		return "", fmt.Errorf("Invalid URL", err)
	}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return "", fmt.Errorf("error parsing HTML", err)
	}

	fmt.Println(baseUrl, doc)

	return "", nil
}
