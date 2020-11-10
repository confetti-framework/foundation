package request

import (
	"github.com/lanvard/foundation/http/request_helper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHeaderDoNotContainJson(t *testing.T) {
	headers := http.Header{}
	headers.Set("content-type", "text/xml")

	assert.False(t, request_helper.IsJson(newRequestWithHeader(headers)))
}

func TestHeaderDoNotContainContentType(t *testing.T) {
	headers := http.Header{}

	assert.False(t, request_helper.IsJson(newRequestWithHeader(headers)))
}

func TestHeaderDoesContainJson(t *testing.T) {
	headers := http.Header{}
	headers.Set("content-type", "application/json")

	assert.True(t, request_helper.IsJson(newRequestWithHeader(headers)))
}

type headerContainer struct {
	headers http.Header
}

func newRequestWithHeader(headers http.Header) headerContainer {
	return headerContainer{headers}
}

func (h headerContainer) Header(key string) string {
	return h.headers.Get(key)
}

func (h headerContainer) GetHeaders() http.Header {
	return h.headers
}
