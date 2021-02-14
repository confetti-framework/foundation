package console

import (
	"github.com/confetti-framework/foundation/console/service"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockCommandWithoutOptions struct{}

func Test_get_option_from_command_without_options(t *testing.T) {
	options := service.GetOptions(mockCommandWithoutOptions{})
	require.Len(t, options, 0)
}

type mockCommandOption struct {
	DryRun bool `flag:"dry-run"`
}

func Test_get_parsed_option(t *testing.T) {
	options := service.GetOptions(mockCommandOption{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "bool", options[0].TypeName)
}

type mockCommandMultipleOptions struct {
	DryRun bool `flag:"dry-run"`
	Vvv    bool `flag:"vvv"`
}

func Test_get_parsed_option_multiple_fields(t *testing.T) {
	options := service.GetOptions(mockCommandMultipleOptions{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "bool", options[0].TypeName)
	require.Equal(t, "vvv", options[1].Tag.Get("flag"))
	require.Equal(t, "bool", options[1].TypeName)
}

type mockCommandOptions struct {
	DryRun bool `flag:"dr" flag2:"dry-run"`
}

func Test_get_parsed_multiple_option(t *testing.T) {
	options := service.GetOptions(mockCommandOptions{})

	require.Equal(t, "dr", options[0].Tag.Get("flag"))
	require.Equal(t, "dry-run", options[0].Tag.Get("flag2"))
}

type mockCommandOptionBool struct {
	DryRun bool `flag:"dry-run"`
}

func Test_get_parsed_option_bool(t *testing.T) {
	options := service.GetOptions(mockCommandOptionBool{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "bool", options[0].TypeName)
}

type mockCommandOptionString struct {
	Username string `flag:"username"`
}

func Test_get_parsed_option_string(t *testing.T) {
	options := service.GetOptions(mockCommandOptionString{})

	require.Equal(t, "username", options[0].Tag.Get("flag"))
	require.Equal(t, "string", options[0].TypeName)
}

type mockCommandOptionInt struct {
	Amount int `flag:"amount"`
}

func Test_get_parsed_option_int(t *testing.T) {
	options := service.GetOptions(mockCommandOptionInt{})

	require.Equal(t, "amount", options[0].Tag.Get("flag"))
	require.Equal(t, "int", options[0].TypeName)
}

type mockCommandOptionFloat struct {
	Number float64 `flag:"number"`
}

func Test_get_parsed_option_float(t *testing.T) {
	options := service.GetOptions(mockCommandOptionFloat{})

	require.Equal(t, "number", options[0].Tag.Get("flag"))
	require.Equal(t, "float64", options[0].TypeName)
}

type mockCommandOptionsWithDescription struct {
	DryRun bool `flag:"dr" flag2:"dry-run" description:"Execute the command as a dry run"`
}

func Test_get_parsed_option_with_description(t *testing.T) {
	options := service.GetOptions(mockCommandOptionsWithDescription{})

	require.Equal(t, "dr", options[0].Tag.Get("flag"))
	require.Equal(t, "dry-run", options[0].Tag.Get("flag2"))
	require.Equal(t, "Execute the command as a dry run", options[0].Tag.Get("description"))
}
