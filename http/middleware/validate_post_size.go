package middleware

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/http"
	"lanvard/app/http/middleware"
)

type ValidatePostSize struct {
	App foundation.Application
}

func (v ValidatePostSize) Handle(data http.Request, next middleware.Destination) http.Response {
	// todo validate
	return next(data)
}
