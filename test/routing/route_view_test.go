package routing

import (
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	routing2 "github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_route_with_view(t *testing.T) {
	// Given
	routes := routing2.Group(
		routing2.View("/comments", SimpleViewMock{}),
	)

	// When
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/comments",
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, SimpleViewMock{}, response.GetContent())
}

type SimpleViewMock struct{}

func (_ SimpleViewMock) Template() string {
	return ""
}
