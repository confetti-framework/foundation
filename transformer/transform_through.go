package transformer

import (
	"errors"
	"github.com/lanvard/contract/inter"
)

func TransformThrough(object interface{}, encoders []inter.ResponseEncoder) (string, error) {
	for _, encoder := range encoders {
		if encoder.Transformable(object) {
			return encoder.TransformThrough(object, encoders)
		}
	}

	return "", errors.New("no transformer found to encode response body")
}
