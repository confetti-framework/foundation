package view_helper

import (
	"bytes"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"html/template"
)

func ContentByView(
	view inter.View,
	builder func(template *template.Template) (*template.Template, error),
) (string, error) {
	buf := bytes.NewBufferString("")
	t, err := template.New("template.view").Parse(view.Template())
	if err != nil {
		return "", err
	}

	t, err = builder(t)
	if err != nil {
		return "", err
	}

	err = t.Execute(buf, view)
	if templateErr, ok := err.(*template.Error); ok && templateErr.ErrorCode == template.ErrNoSuchTemplate {
		//goland:noinspection ALL
		err = errors.Wrap(NoSuchTemplateError, err.Error())
	}

	return buf.String(), err
}
