package encode

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_string_to_string_no_match_with_zero(t *testing.T) {
	var val int
	require.False(t, encoder.StringToString{}.IsAble(val))
}

func Test_string_to_string_no_match_with_string(t *testing.T) {
	require.True(t, encoder.StringToString{}.IsAble("val"))
}

func Test_string_to_string_encode_valid_string(t *testing.T) {
	var app inter.App
	result, err := encoder.StringToString{}.EncodeThrough(app, "val", []inter.Encoder{})
	assert.Nil(t, err)
	require.Equal(t, "val", result)
}
