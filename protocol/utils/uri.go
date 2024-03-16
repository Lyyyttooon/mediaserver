package utils

import "net/url"

func ParseURI(uri string) (*url.URL, error) {
	return url.Parse(uri)
}
