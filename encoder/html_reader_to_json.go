package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type HtmlReaderToHtml struct{}

func (h HtmlReaderToHtml) IsAble(object interface{}) bool {
	_, ok := object.(inter.HtmlReader)
	return ok
}

func (h HtmlReaderToHtml) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	result, ok := object.(inter.HtmlReader)
	if !ok {
		return "", errors.New("can not encode to html with an unsupported type " + reflect.TypeOf(object).String())
	}
	return result.String(), nil
}
