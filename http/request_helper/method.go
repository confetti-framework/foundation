package request_helper

import (
	"github.com/lanvard/contract/inter"
	"strings"
)

func IsMethod(request inter.Request, method string) bool {
	return request.Method() == strings.ToUpper(method)
}
