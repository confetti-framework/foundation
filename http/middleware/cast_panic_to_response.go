package middleware

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
)

type PanicToResponse struct {
	Outcome func(content interface{}) inter.Response
}

// Change an error from a temporary response to a complete response
func (p PanicToResponse) Handle(request inter.Request, next inter.Next) (response inter.Response) {
	defer func() {
		if r := recover(); r != nil {
			var err error
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("can't convert panic to response. Error or string required")
			}

			response = outcome.Json(err)
		}
	}()

	return next(request)
}
