package view_helper

import (
	"bytes"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/foundation/report"
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
		err = errors.Wrap(report.NoSuchTemplate, err.Error())
	}

	return buf.String(), err
}
