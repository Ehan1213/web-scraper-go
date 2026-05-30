package main

import (
	"errors"
)

func normalizeURL(url string) (_ string, err error) {
	if url == "" {
		return "", errors.New("empty url")
	}

	return "nil", err
}
