package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support/str"
)

type ErrorToJson struct {
	Jsonapi map[string]string        `json:"jsonapi"`
	Errors  []map[string]interface{} `json:"errors"`
}

func (e ErrorToJson) IsAble(object interface{}) bool {
	_, ok := object.(error)
	return ok
}

func (e ErrorToJson) EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	err, ok := object.(error)
	if !ok {
		return "", errors.New("can't convert object to json in error format")
	}

	title := str.UpperFirst(err.Error())
	e.Jsonapi = map[string]string{"version": "1.0"}
	e.Errors = []map[string]interface{}{{"title": title}}

	return EncodeThrough(e, encoders)
}
