package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type RawToJson struct{}

func (v RawToJson) IsAble(object interface{}) bool {
	value, ok := object.(interface{ Raw() interface{} })
	return ok && InterfaceToJson{}.IsAble(value.Raw())
}

func (v RawToJson) EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	value, ok := object.(interface{ Raw() interface{} })
	if !ok {
		return "", errors.New("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	return EncodeThrough(value.Raw(), encoders)
}
