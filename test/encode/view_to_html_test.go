package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/stretchr/testify/require"
	"html/template"
	"testing"
)

func Test_view_is_able_nil(t *testing.T) {
	require.False(t, encoder.ViewToHtml{}.IsAble(nil))
}

func Test_view_is_able_with_view(t *testing.T) {
	require.True(t, encoder.ViewToHtml{}.IsAble(mock.SimpleViewMock{}))
}

func Test_can_not_encode_non_view(t *testing.T) {
	app := setUp()
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, nil, nil)
	require.Error(t, err)
	require.Equal(t, "", result)
}

func Test_convert_view_to_html(t *testing.T) {
	app := setUp()
	app.Bind("template_builder", func(template *template.Template) (*template.Template, error) {
		return template.Parse("ok")
	})
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, mock.SimpleViewMock{}, nil)
	require.NoError(t, err)
	require.Equal(t, "ok", result)
}

func Test_view_with_footer_template(t *testing.T) {
	app := setUp()
	app.Bind("template_builder", func(template *template.Template) (*template.Template, error) {
		return template.ParseFiles(
			mock.TemplateByName("page_with_footer_template.gohtml"),
			mock.TemplateByName("sub/footer_template.gohtml"),
		)
	})
	app.Singleton("defined_templates", []string{mock.TemplateByName("footer_template.gohtml")})
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, mock.PageWithFooter{}, nil)
	require.NoError(t, err)
	require.Equal(t, "<p>body</p>\n\n    <footer>contact information</footer>", result)
}
