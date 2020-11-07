package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/view_helper"
	"github.com/lanvard/foundation/report"
)

type ViewToHtml struct{}

func (v ViewToHtml) IsAble(object interface{}) bool {
	_, ok := object.(inter.View)
	return ok
}

func (v ViewToHtml) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	view, ok := object.(inter.View)
	if !ok {
		return "", report.EncodeError.Wrap("can't convert object to html in view format")
	}

	builder := app.Make("template_builder").(inter.TemplateBuilder)
	return view_helper.ContentByView(view, builder)
}
