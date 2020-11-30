package mock

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
)

var JsonEncoders = []inter.Encoder{
	encoder.JsonReaderToJson{},
	encoder.RawToJson{},
	encoder.JsonToJson{},
	encoder.ErrorSToJson{},
	encoder.InterfaceToJson{}, // todo: interface is now the default, can't override in ResponseServiceProvider
}

var HtmlEncoders = []inter.Encoder{
	encoder.StringerToHtml{},
	encoder.RawToHtml{},
	encoder.InterfaceToHtml{},
}
