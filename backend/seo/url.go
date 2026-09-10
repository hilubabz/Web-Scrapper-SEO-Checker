package seo

import (
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", nil
	}

	parsedURL.Fragment = ""

	parsedURL.Scheme = strings.ToLower(parsedURL.Scheme)
	parsedURL.Host = strings.ToLower(parsedURL.Host)

	if parsedURL.Scheme == "http" && parsedURL.Port() == "80" {
		parsedURL.Host = parsedURL.Hostname()
	}

	if parsedURL.Scheme == "https" && parsedURL.Port() == "443" {
		parsedURL.Host = parsedURL.Hostname()
	}

	if parsedURL.Path == "" {
		parsedURL.Path = "/"
	}

	return parsedURL.String(), nil
}