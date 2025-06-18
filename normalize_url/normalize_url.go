package normalize_url

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeUrl(rawURL string) (string, error) {
	parse_url, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parsing error %w", err)
	}
	fullPath := parse_url.Host + parse_url.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
