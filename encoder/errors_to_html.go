package encoder

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/http/view_helper"
	"github.com/confetti-framework/support/str"
)

type ErrorsToHtml struct {
	View func(app inter.App, err error) inter.View
}

func (e ErrorsToHtml) IsAble(object interface{}) bool {
	_, ok := e.getErrors(object)
	return ok
}

func (e ErrorsToHtml) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	errs, ok := e.getErrors(object)
	if !ok {
		return "", errors.Wrap(EncodeError, "can't convert object to html in error format")
	}

	err := errs[0]
	if e.View != nil {
		builder := app.Make("template_builder").(inter.TemplateBuilder)
		return view_helper.ContentByView(e.View(app, err), builder)
	}

	return str.UpperFirst(fmt.Sprintf("%v", err)), nil
}

func (e ErrorsToHtml) getErrors(object interface{}) ([]error, bool) {
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
