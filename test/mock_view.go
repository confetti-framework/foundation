package test

type ViewMock struct {
	TemplatePath string
}

func (v ViewMock) Template() string {
	return v.TemplatePath
}
