package console

import (
	"github.com/confetti-framework/foundation/console/service"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_cast_to_bool(t *testing.T) {
	require.True(t, service.CastToBool("true").(bool))
	require.False(t, service.CastToBool("false").(bool))
}

func Test_cast_to_float(t *testing.T) {
	require.Equal(t, 1., service.CastToFloat("1").(float64))
	require.Equal(t, 1.5, service.CastToFloat("1.5").(float64))
}

func Test_cast_to_int(t *testing.T) {
	require.Equal(t, 1, service.CastToInt("1").(int))
}

func Test_cast_to_string(t *testing.T) {
	require.Equal(t, "piet", service.CastToString("piet").(string))
}
