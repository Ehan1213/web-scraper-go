package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return html, err
	}
	h1 := doc.Find("h1")
	if h1.Length() >= 1 {
		return h1.Text(), nil
	}
	h2 := doc.Find("h2")
	if h2.Length() >= 1 {
		return h2.Text(), nil
	}

	return "", fmt.Errorf("no headings found: %w", err)
}

func getFirstParagraphFromHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return html, err
	}
	main := doc.Find("p")
	if main.Length() >= 1 {
		return main.First().Text(), nil
	}
	p := doc.Find("p")
	if p.Length() >= 1 {
		return p.First().Text(), nil
	}
	return "", fmt.Errorf("no paragraphs found: %w", err)
}

func getURLsFromHTML(html string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return []string{html}, err
	}
	absoluteURLS := []string{}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		parsedHref, err := url.Parse(link)
		if err != nil {
			return
		}
		absolute := baseURL.ResolveReference(parsedHref).String()
		absoluteURLS = append(absoluteURLS, absolute)
	})
	return absoluteURLS, err
}
