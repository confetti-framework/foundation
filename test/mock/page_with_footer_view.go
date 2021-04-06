package mock

import _ "embed"

//go:embed page_with_footer_template.gohtml
var pageWithFooterTemplate string

type PageWithFooter struct{}

func (_ PageWithFooter) Template() string {
	return pageWithFooterTemplate
}
