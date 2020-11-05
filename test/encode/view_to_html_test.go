package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestViewIsAbleNil(t *testing.T) {
	require.False(t, encoder.ViewToHtml{}.IsAble(nil))
}

func TestViewIsAbleWithView(t *testing.T) {
	require.True(t, encoder.ViewToHtml{}.IsAble(mock.SimpleViewMock{}))
}

func TestCanNotEncodeNonView(t *testing.T) {
	app := setUp()
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, nil, nil)
	require.Error(t, err)
	require.Equal(t, "", result)
}

func TestConvertViewToHtml(t *testing.T) {
	app := setUp()
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, mock.SimpleViewMock{}, nil)
	require.Nil(t, err)
	require.Equal(t, "ok", result)
}

//goland:noinspection GoNilness
func TestViewWithUnknownTemplate(t *testing.T) {
	app := setUp()
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, mock.PageWithFooter{}, nil)
	require.Contains(t, err.Error(), `no such template "footer"`)
	require.Equal(t, "", result)
}

func TestViewWithFooterTemplate(t *testing.T) {
	app := setUp()
	app.Singleton("defined_templates", []string{mock.TemplateByName("footer_template.gohtml")})
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, mock.PageWithFooter{}, nil)
	require.Nil(t, err)
	require.Equal(t, "<p>body</p>\n<footer>contact information</footer>", result)
}
