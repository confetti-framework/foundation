package encoder

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support/str"
)

type ErrorSToJson struct {
	Jsonapi map[string]string `json:"jsonapi"`
	Errors  []Error           `json:"errors"`
}

type Error struct {
	Title string `json:"title"`
}

func (e ErrorSToJson) IsAble(object interface{}) bool {
	_, ok := e.getErrors(object)
	return ok
}

func (e ErrorSToJson) EncodeThrough(app inter.App, object interface{}, encoders []inter.Encoder) (string, error) {
	e.Errors = []Error{}
	errs, ok := e.getErrors(object)
	if !ok {
		return "", EncodeError.Wrap("can't convert object to json in error format")
	}

	e.Jsonapi = map[string]string{"version": "1.0"}

	for _, err := range errs {
		e.Errors = append(e.Errors, Error{
			Title: str.UpperFirst(err.Error()),
		})
	}

	return EncodeThrough(app, e, encoders)
}

func (e ErrorSToJson) getErrors(object interface{}) ([]error, bool) {
	err, ok := object.(error)
	if ok {
		return []error{err}, ok
	}

	errs, ok := object.([]error)
	return errs, ok
}
