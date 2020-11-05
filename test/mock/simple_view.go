package mock

type SimpleViewMock struct{}

func (_ SimpleViewMock) Template() string {
	return TemplateByName("/simple_template.gohtml")
}
