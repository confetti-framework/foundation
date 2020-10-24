package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"reflect"
)

func EncodeThrough(app inter.App, content interface{}, encoders []inter.Encoder) (string, error) {
	for _, encoder := range encoders {
		if encoder.IsAble(content) {
			return encoder.EncodeThrough(app, content, encoders)
		}
	}

	if err, ok := content.(error); ok {
		err := errors.New("no encoder found to handle error: " + err.Error())
		return err.Error(), err
	}

	err := errors.New("no encoder found to encode response body with type " + reflect.TypeOf(content).String())
	return EncodeThrough(app, err, encoders)
}
