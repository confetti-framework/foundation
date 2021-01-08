package http_helper

import (
	"github.com/confetti-framework/contract/inter"
	"strings"
)

func HasJson(headerHolder inter.Request) bool {
	header := headerHolder.Header("Content-Type")
	for _, jsonHeader := range []string{"/json", "+json"} {
		if strings.Contains(header, jsonHeader) {
			return true
		}
	}

	return false
}

func HasMultiPartFormData(headerHolder inter.Request) bool {
	header := headerHolder.Header("Content-Type")
	for _, jsonHeader := range []string{"multipart/form-data", "application/x-www-form-urlencoded"} {
		if strings.Contains(header, jsonHeader) {
			return true
		}
	}

	return false
}
