package middleware

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http/lanvard"
	"lanvard/app/http/middleware"
)

type ValidatePostSize struct {
	App foundation.Application
}

func (v ValidatePostSize) Handle(data lanvard.Request, next middleware.Destination) lanvard.Response {
	// todo validate
	return next(data)
}
