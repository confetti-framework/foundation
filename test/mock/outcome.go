package mock

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/encoder"
)

var JsonEncoders = []inter.Encoder{
	encoder.JsonReaderToJson{},
	encoder.RawToJson{},
	encoder.JsonToJson{},
	encoder.ErrorsToJson{},
	encoder.InterfaceToJson{}, // todo: interface is now the default, can't override in ResponseServiceProvider
}

var HtmlEncoders = []inter.Encoder{
	encoder.StringerToHtml{},
	encoder.RawToHtml{},
	encoder.InterfaceToHtml{},
}
