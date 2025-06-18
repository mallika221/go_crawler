package main

import (
	"fmt"
	"go_crawler/normalize_url"
)

func main() {
	fmt.Println("hello world")
	urlString, err := normalize_url.NormalizeUrl("https://blog.boot.dev/path/")
	if err != nil {
		fmt.Println("error %w ", err)
		return

	}
	fmt.Println("Normalized URL", urlString)
}
