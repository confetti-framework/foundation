package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
)

type PanicToResponse struct{}

// Change an error from a temporary response to a complete response
func (p PanicToResponse) Handle(request inter.Request, next inter.Next) (response inter.Response) {

	defer func() {
		if r := recover(); r != nil {
			instance, err := request.App().MakeE("default_response_outcome")
			if err != nil {
				panic(errors.Wrap(err, "please bind a encoder"))
			}
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("can't convert panic to response. IsError or string required")
			}

			response = instance.(func(interface{}) inter.Response)(err)
		}
	}()

	return next(request)
}
