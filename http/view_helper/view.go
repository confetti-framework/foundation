package view_helper

import (
	"bytes"
	"github.com/lanvard/contract/inter"
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

	return buf.String(), err
}
