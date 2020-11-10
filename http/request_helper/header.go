package request_helper

import (
	"github.com/lanvard/contract/inter"
	"strings"
)

func IsJson(headerHolder inter.Request) bool {
	header := headerHolder.Header("Content-Type")
	for _, jsonHeader := range []string{"/json", "+json"} {
		if strings.Contains(header, jsonHeader) {
			return true
		}
	}

	return false
}
