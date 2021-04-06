package mock

import _ "embed"

//go:embed simple_template.gohtml
var simpleTemplate string

type SimpleViewMock struct{}

func (_ SimpleViewMock) Template() string {
	return simpleTemplate
}
