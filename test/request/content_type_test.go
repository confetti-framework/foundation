package request

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/request_helper"
	"github.com/stretchr/testify/require"
	net "net/http"
	"testing"
)

func Test_header_do_not_contain_json(t *testing.T) {
	headers := net.Header{}
	headers.Set("content-type", "text/xml")

	require.False(t, request_helper.HasJson(newRequestWithHeader(headers)))
}

func Test_header_do_not_contain_content_type_json(t *testing.T) {
	headers := net.Header{}

	require.False(t, request_helper.HasJson(newRequestWithHeader(headers)))
}

func Test_header_does_contain_json(t *testing.T) {
	headers := net.Header{}
	headers.Set("content-type", "application/json")

	require.True(t, request_helper.HasJson(newRequestWithHeader(headers)))
}

func Test_header_do_not_contain_form_data(t *testing.T) {
	headers := net.Header{}
	headers.Set("content-type", "text/xml")

	require.False(t, request_helper.HasMultiPartFormData(newRequestWithHeader(headers)))
}

func Test_header_do_not_contain_content_type_form_data(t *testing.T) {
	headers := net.Header{}

	require.False(t, request_helper.HasMultiPartFormData(newRequestWithHeader(headers)))
}

func Test_header_does_contain_form_data(t *testing.T) {
	headers := net.Header{}
	headers.Set("content-type", "multipart/form-data")

	require.True(t, request_helper.HasMultiPartFormData(newRequestWithHeader(headers)))
}

func newRequestWithHeader(headers net.Header) inter.Request {
	return http.NewRequest(http.Options{
		Header: headers,
	})
}
