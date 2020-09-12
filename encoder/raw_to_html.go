package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type RawToHtml struct{}

func (v RawToHtml) IsAble(object interface{}) bool {
	_, ok := object.(interface{ Raw() interface{} })
	return ok
}

func (v RawToHtml) EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	result, ok := object.(interface{ Raw() interface{} })
	if !ok {
		return "", errors.New("can not encode to html with an unsupported type " + reflect.TypeOf(object).String())
	}

	res := result.Raw().(string)
	return res, nil
}
