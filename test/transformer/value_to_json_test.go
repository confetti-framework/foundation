package transformer

import (
	"github.com/lanvard/foundation/transformer"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringCanNotTransformValueExpected(t *testing.T) {
	assert.False(t, transformer.ValueToJson{}.Transformable("foo"))
}

func TestCanTransformValue(t *testing.T) {
	assert.True(t, transformer.ValueToJson{}.Transformable(support.NewValue(foo{})))
}

func TestTransformStringValueExpected(t *testing.T) {
	assert.PanicsWithValue(t, "can not transform to json with an unsupported type", func() {
		transformer.ValueToJson{}.Transform("foo")
	})
}
