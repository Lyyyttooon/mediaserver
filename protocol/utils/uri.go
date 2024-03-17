package utils

import (
	"fmt"
	"net/url"

	"github.com/Lyyyttooon/mediaserver/protocol/common"
)

func ParseURI(uri string) (*url.URL, error) {
	uriParsed, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if uriParsed.Port() == "" {
		switch uriParsed.Scheme {
		case "http":
			uriParsed.Host += fmt.Sprintf(":%d", common.HttpDefaultPort)
		case "https":
			uriParsed.Host += fmt.Sprintf(":%d", common.HttpsDefaultPort)
		}
	}

	return uriParsed, nil
}
