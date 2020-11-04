package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/foundation/http/view_helper"
)

type ViewToHtml struct{}

func (v ViewToHtml) IsAble(object interface{}) bool {
	_, ok := object.(inter.View)
	return ok
}

func (v ViewToHtml) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	view, ok := object.(inter.View)
	if !ok {
		return "", errors.New("can't convert object to html in view format")
	}

	return view_helper.ContentByView(view)
}
