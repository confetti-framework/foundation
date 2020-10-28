package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
	"html/template"
	"io/ioutil"
	"os"
)

type ErrorToHtml struct {
	TemplatePath string
}

func (e ErrorToHtml) IsAble(object interface{}) bool {
	_, ok := object.(error)
	return ok
}

func (e ErrorToHtml) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	err, ok := object.(error)
	if !ok {
		return "", errors.New("can't convert object to html in error format")
	}

	return str.UpperFirst(err.Error()), nil
}
