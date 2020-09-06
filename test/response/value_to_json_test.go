package response

import (
	"github.com/lanvard/support"
	"github.com/lanvard/support/transformer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringCanNotTransformValueExpected(t *testing.T) {
	assert.False(t, transformer.ValueToJson{}.IsValid("foo"))
}

func TestCanTransformValue(t *testing.T) {
	assert.True(t, transformer.ValueToJson{}.IsValid(support.NewValue(Foo{})))
}

func TestTransformStringValueExpected(t *testing.T) {
	assert.PanicsWithValue(t, "can not transform to json with an unsupported type", func() {
		transformer.ValueToJson{}.Transform("foo")
	})
}
