package durable

import (
	"encoding/base64"
	"net/url"
)

func URLDecode(input string) (string, error) {
	decodedURL, err := url.QueryUnescape(input)
	if err != nil {
		return "", err
	}

	base64DecodedText, err := base64.StdEncoding.DecodeString(decodedURL)
	if err != nil {
		return "", err
	}

	return string(base64DecodedText), nil
}
