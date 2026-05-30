package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func normalizeURL(URL string) (_ string, err error) {
	if URL == "" {
		return "", err

	}
	clean := strings.TrimSpace(URL)
	url, err := url.ParseRequestURI(clean)
	if err != nil {
		log.Fatal(err)
		return url.Scheme, err
	}

	return fmt.Sprintf("%s%s", url.Host, url.Path), err
}
