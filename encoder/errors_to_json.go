package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
)

type ErrorToJson struct {
	Jsonapi map[string]string `json:"jsonapi"`
	Errors  []Error           `json:"errors"`
}

type Error struct {
	Title string `json:"title"`
}

func (e ErrorToJson) IsAble(object interface{}) bool {
	_, ok := object.(error)
	return ok
}

func (e ErrorToJson) EncodeThrough(app inter.App, object interface{}, encoders []inter.Encoder) (string, error) {
	err, ok := object.(error)
	if !ok {
		return "", errors.New("can't convert object to json in error format")
	}

	e.Jsonapi = map[string]string{"version": "1.0"}
	e.Errors = append(e.Errors, Error{
		Title: str.UpperFirst(err.Error()),
	})

	return EncodeThrough(app, e, encoders)
}
