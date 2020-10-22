package middleware

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

type PanicToResponse struct{}

// Change an error from a temporary response to a complete response
func (p PanicToResponse) Handle(request inter.Request, next inter.Next) (response inter.Response) {

	defer func() {
		if r := recover(); r != nil {
			instance, err := request.App().MakeE("response_encoder")
			if err != nil {
				panic(support.Wrap(err, "please bind a encoder"))
			}
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("can't convert panic to response. Error or string required")
			}

			response = instance.(func(interface{}) inter.Response)(err)
		}
	}()

	return next(request)
}
