package encoder

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/http/view_helper"
)

type ViewToHtml struct{}

func (v ViewToHtml) IsAble(object interface{}) bool {
	_, ok := object.(inter.View)
	return ok
}

func (v ViewToHtml) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	view, ok := object.(inter.View)
	if !ok {
		return "", errors.Wrap(EncodeError, "can't convert object to html in view format")
	}

	builder := app.Make("template_builder").(inter.TemplateBuilder)
	return view_helper.ContentByView(view, builder)
}
