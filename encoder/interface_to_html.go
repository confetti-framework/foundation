package encoder

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/spf13/cast"
)

type InterfaceToHtml struct{}

func (j InterfaceToHtml) IsAble(object interface{}) bool {
	_, isError := object.(error)
	if isError {
		return false
	}
	_, err := cast.ToStringE(object)
	return err == nil
}

func (j InterfaceToHtml) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	if object == nil {
		return "", nil
	}

	result, err := cast.ToStringE(object)
	if err != nil {
		return "", errors.Wrap(EncodeError, "can not encode to html: "+err.Error())
	}

	return result, nil
}
