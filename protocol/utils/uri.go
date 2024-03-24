package utils

import (
	"fmt"
	"net/url"
)

const (
	HttpDefaultPort  = 80
	HttpsDefaultPort = 443
)

func ParseURI(uri string) (*url.URL, error) {
	uriParsed, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if uriParsed.Port() == "" {
		switch uriParsed.Scheme {
		case "http":
			uriParsed.Host += fmt.Sprintf(":%d", HttpDefaultPort)
		case "https":
			uriParsed.Host += fmt.Sprintf(":%d", HttpsDefaultPort)
		}
	}

	return uriParsed, nil
}
