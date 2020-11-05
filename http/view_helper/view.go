package view_helper

import (
	"bytes"
	"github.com/lanvard/contract/inter"
	"html/template"
	"path"
)

func ContentByView(view inter.View, templates []string) (string, error) {
	main := path.Base(view.Template())

	buf := bytes.NewBufferString("")
	t, err := template.New(main).
		ParseFiles(append(templates, view.Template())...)
	if err != nil {
		return "", err
	}
	err = t.Execute(buf, view)

	return buf.String(), err
}
