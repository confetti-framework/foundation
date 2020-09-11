package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
)

func EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	for _, encoder := range encoders {
		if encoder.IsAble(object) {
			return encoder.EncodeThrough(object, encoders)
		}
	}

	return "", errors.New("no transformer found to encode response body")
}
