package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type InterfaceToHtml struct{}

func (j InterfaceToHtml) IsAble(object interface{}) bool {
	_, ok := object.(string)
	return ok || object == nil
}

func (j InterfaceToHtml) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	if object == nil {
		return "", nil
	}

	result, ok := object.(string)
	if !ok {
		return "", errors.New("can not encode to html with an unsupported type " + reflect.TypeOf(object).String())
	}

	return result, nil
}
