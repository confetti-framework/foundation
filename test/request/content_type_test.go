package request

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/request_helper"
	"github.com/stretchr/testify/assert"
	net "net/http"
	"testing"
)

func Test_header_do_not_contain_json(t *testing.T) {
	headers := net.Header{}
	headers.Set("content-type", "text/xml")

	assert.False(t, request_helper.IsJson(newRequestWithHeader(headers)))
}

func Test_header_do_not_contain_content_type(t *testing.T) {
	headers := net.Header{}

	assert.False(t, request_helper.IsJson(newRequestWithHeader(headers)))
}

func Test_header_does_contain_json(t *testing.T) {
	headers := net.Header{}
	headers.Set("content-type", "application/json")

	assert.True(t, request_helper.IsJson(newRequestWithHeader(headers)))
}

func newRequestWithHeader(headers net.Header) inter.Request {
	return http.NewRequest(http.Options{
		Header: headers,
	})
}
