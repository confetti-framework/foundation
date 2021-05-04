package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/http/http_helper"
)

type PanicToResponse struct{}

// Change an error from a temporary response to a complete response
func (p PanicToResponse) Handle(request inter.Request, next inter.Next) (response inter.Response) {

	defer func() {
		rec := recover()
		if rec != nil {
			instance := getDefaultResponseEncoder(request)
			if err, ok := rec.(error); ok {
				rec = errors.WithStack(err)
			}
			response = instance(http_helper.GetErrorFromPanic(rec))
		}
	}()

	return next(request)
}

func getDefaultResponseEncoder(request inter.Request) func(interface{}) inter.Response {
	instance, err := request.App().MakeE("default_response_outcome")
	if err != nil {
		panic(errors.Wrap(err, "please bind a encoder"))
	}
	return instance.(func(interface{}) inter.Response)
}
