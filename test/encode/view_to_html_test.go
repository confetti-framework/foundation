package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewIsAbleNil(t *testing.T) {
	assert.False(t, encoder.ViewToHtml{}.IsAble(nil))
}

func TestViewIsAbleWithView(t *testing.T) {
	assert.True(t, encoder.ViewToHtml{}.IsAble(test.ViewMock{}))
}

func TestCanNotEncodeNonView(t *testing.T) {
	app := setUp()
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, nil, nil)
	assert.Error(t, err)
	assert.Equal(t, "", result)
}

func TestConvertViewToHtml(t *testing.T) {
	app := setUp()
	result, err := encoder.ViewToHtml{}.EncodeThrough(app, test.ViewMock{}, nil)
	assert.Nil(t, err)
	assert.Equal(t, "ok", result)
}
