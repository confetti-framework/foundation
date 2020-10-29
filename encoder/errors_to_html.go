package encoder

import (
	"bytes"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
	"html/template"
	"io/ioutil"
)

type ErrorToHtml struct {
	TemplateFile string
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

	errorMessage := str.UpperFirst(err.Error())
	if e.TemplateFile != "" {
		buf := bytes.NewBufferString("")
		status, _ := errors.FindStatus(err)
		view := ErrorView{
			Message: errorMessage,
			Status:  status,
		}

		message, err := ioutil.ReadFile(e.TemplateFile)
		if err != nil {
			return "", err
		}
		t, err := template.New("error").Parse(string(message))
		if err != nil {
			return "", err
		}
		err = t.Execute(buf, view)

		return buf.String(), nil
	}
	return errorMessage, nil
}

type ErrorView struct {
	Message string
	Status  int
}
