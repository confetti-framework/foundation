package encoder

import (
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/view_helper"
	"github.com/lanvard/support/str"
)

type ErrorToHtml struct {
	View func(app inter.App, err error) inter.View
}

func (e ErrorToHtml) IsAble(object interface{}) bool {
	_, ok := e.getErrors(object)
	return ok
}

func (e ErrorToHtml) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	errs, ok := e.getErrors(object)
	if !ok {
		return "", EncodeError.Wrap("can't convert object to html in error format")
	}

	err := errs[0]
	if e.View != nil {
		builder := app.Make("template_builder").(inter.TemplateBuilder)
		return view_helper.ContentByView(e.View(app, err), builder)
	}

	return str.UpperFirst(fmt.Sprintf("%v", err)), nil
}

func (e ErrorToHtml) getErrors(object interface{}) ([]error, bool) {
	err, ok := object.(error)
	if ok {
		return []error{err}, ok
	}

	errs, ok := object.([]error)
	if len(errs) == 0 {
		return []error{}, false
	}
	return errs, ok
}
