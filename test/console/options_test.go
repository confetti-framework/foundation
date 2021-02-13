package console

import (
	"github.com/confetti-framework/foundation/console/service"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type mockCommandWithoutOptions struct{}

func Test_get_option_from_command_without_options(t *testing.T) {
	options := service.GetOptions(mockCommandWithoutOptions{})
	require.Len(t, options, 0)
}

type mockCommandOption struct {
	DryRun bool `flag:"--dry-run"`
}

func Test_get_parsed_option(t *testing.T) {
	options := service.GetOptions(mockCommandOption{})

	require.Equal(t, "--dry-run", options[0].Tag.Get("flag"))
	require.True(t, functionsAreEqual(service.CastToBool, options[0].Caster))
}

type mockCommandMultipleOptions struct {
	DryRun bool `flag:"--dry-run"`
	Vvv    bool `flag:"--vvv"`
}

func Test_get_parsed_option_multiple_fields(t *testing.T) {
	options := service.GetOptions(mockCommandMultipleOptions{})

	require.Equal(t, "--dry-run", options[0].Tag.Get("flag"))
	require.True(t, functionsAreEqual(service.CastToBool, options[0].Caster))
	require.Equal(t, "--vvv", options[1].Tag.Get("flag"))
	require.True(t, functionsAreEqual(service.CastToBool, options[1].Caster))
}

type mockCommandOptions struct {
	DryRun bool `flag:"-dr" flag2:"--dry-run"`
}

func Test_get_parsed_multiple_option(t *testing.T) {
	options := service.GetOptions(mockCommandOptions{})

	require.Equal(t, "-dr", options[0].Tag.Get("flag"))
	require.Equal(t, "--dry-run", options[0].Tag.Get("flag2"))
}

type mockCommandOptionBool struct {
	DryRun bool `flag:"--dry-run"`
}

func Test_get_parsed_option_bool(t *testing.T) {
	options := service.GetOptions(mockCommandOptionBool{})

	require.Equal(t, "--dry-run", options[0].Tag.Get("flag"))
	require.True(t, functionsAreEqual(service.CastToBool, options[0].Caster))
}

type mockCommandOptionString struct {
	Username string `flag:"--username"`
}

func Test_get_parsed_option_string(t *testing.T) {
	options := service.GetOptions(mockCommandOptionString{})

	require.Equal(t, "--username", options[0].Tag.Get("flag"))
	require.True(t, functionsAreEqual(service.CastToString, options[0].Caster))
}

type mockCommandOptionInt struct {
	Amount int `flag:"--amount"`
}

func Test_get_parsed_option_int(t *testing.T) {
	options := service.GetOptions(mockCommandOptionInt{})

	require.Equal(t, "--amount", options[0].Tag.Get("flag"))
	require.True(t, functionsAreEqual(service.CastToInt, options[0].Caster))
}

type mockCommandOptionFloat struct {
	Number float64 `flag:"--number"`
}

func Test_get_parsed_option_float(t *testing.T) {
	options := service.GetOptions(mockCommandOptionFloat{})

	require.Equal(t, "--number", options[0].Tag.Get("flag"))
	require.True(t, functionsAreEqual(service.CastToFloat, options[0].Caster))
}

type mockCommandOptionsWithDescription struct {
	DryRun bool `flag:"-dr" flag2:"--dry-run" description:"Execute the command as a dry run"`
}

func Test_get_parsed_option_with_description(t *testing.T) {
	options := service.GetOptions(mockCommandOptionsWithDescription{})

	require.Equal(t, "-dr", options[0].Tag.Get("flag"))
	require.Equal(t, "--dry-run", options[0].Tag.Get("flag2"))
	require.Equal(t, "Execute the command as a dry run", options[0].Tag.Get("description"))
}

func functionsAreEqual(function1 interface{}, function2 interface{}) bool {
	return reflect.ValueOf(function1).Pointer() == reflect.ValueOf(function2).Pointer()
}
