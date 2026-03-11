package logic

import "net/url"

func IsValidURL(raw string) bool {
	parsedURL, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}
