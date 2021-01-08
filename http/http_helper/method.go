package http_helper

import (
	"github.com/confetti-framework/contract/inter"
	"strings"
)

func IsMethod(request inter.Request, method string) bool {
	return request.Method() == strings.ToUpper(method)
}
