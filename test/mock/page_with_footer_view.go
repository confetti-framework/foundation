package mock

type PageWithFooter struct{}

func (_ PageWithFooter) Template() string {
	return TemplateByName("/page_with_footer_template.gohtml")
}
