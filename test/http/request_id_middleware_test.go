package http

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/stretchr/testify/require"
	net "net/http"
	"testing"
)

func Test_request_id_middleware(t *testing.T) {
	t.Run("check without request id given from original request", func(t *testing.T) {
		req := http.NewRequest(http.Options{})

		resp := middleware.RequestID{}.Handle(req, dummyMiddlewareResponder)
		requestID := req.Headers().Get("x-request-id")
		require.NotEmpty(t, requestID)
		require.Equal(t, requestID, resp.GetHeaders().Get("x-request-id"))
	})

	t.Run("check when request id is already set in original request", func(t *testing.T) {
		givenRequestID := "an-example-uuid"
		reqHeaders := net.Header{}
		reqHeaders.Add("x-request-id", givenRequestID)
		req := http.NewRequest(http.Options{Header: reqHeaders})

		resp := middleware.RequestID{}.Handle(req, dummyMiddlewareResponder)
		require.Equal(t, givenRequestID, resp.GetHeaders().Get("x-request-id"))
	})
}

func dummyMiddlewareResponder(_ inter.Request) inter.Response {
	return outcome.Html("hello world")
}
