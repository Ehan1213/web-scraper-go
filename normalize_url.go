package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string string, err error) {
	// if rawURL == "" {
	// 	return "", err

	// }
	rawURL = strings.TrimSpace(rawURL)
	rawURL = strings.ReplaceAll(rawURL, `\\`, `/`)
	rawURL = strings.ReplaceAll(rawURL, `\`, `/`)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
