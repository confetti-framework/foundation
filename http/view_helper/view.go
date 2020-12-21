package view_helper

import (
	"bytes"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"html/template"
	"path"
)

func ContentByView(
	view inter.View,
	builder func(template *template.Template) (*template.Template, error),
) (string, error) {
	buf := bytes.NewBufferString("")
	t := template.New(path.Base(view.Template()))

	t, err := builder(t)
	if err != nil {
		return "", err
	}

	err = t.Execute(buf, view)
	if templateErr, ok := err.(*template.Error); ok && templateErr.ErrorCode == template.ErrNoSuchTemplate {
		err = errors.Wrap(NoSuchTemplateError, err.Error())
	}

	return buf.String(), err
}
