package view_helper

import (
	"bytes"
	"github.com/lanvard/contract/inter"
	"html/template"
	"io/ioutil"
)

func ContentByView(view inter.View) (string, error) {
	content, err := ioutil.ReadFile(view.Template())
	buf := bytes.NewBufferString("")
	if err != nil {
		return "", err
	}
	t, err := template.New("error").Parse(string(content))
	if err != nil {
		return "", err
	}
	err = t.Execute(buf, view)

	return buf.String(), nil
}
