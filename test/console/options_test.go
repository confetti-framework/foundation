package console

import (
	"testing"

	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/confetti-framework/foundation/console/service"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
)

type mockCommandWithoutOptions struct{}

func Test_show_index_if_no_command(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithDescription{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	result := TrimDoubleSpaces(output.String())
	require.Contains(t, result, "\n Confetti\x1b[39m\n\n")

	require.Contains(t, result, "\n \x1b[32mGlobal options:\x1b[0m")
	require.Contains(t, result, "-h --help Show the command's available arguments.")
	require.Contains(t, result, "--env-file Run the command with a defined environment file.")
}

func Test_show_index_with_groups(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithDescription{CommandName: "make:test"}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	result := TrimDoubleSpaces(output.String())

	require.Contains(t, result, "\x1b[32mAvailable commands:\x1b[0m")
	require.Contains(t, result, "\n \x1b[32mmake\x1b[0m\n \x1b[33mmake:test\x1b[0m test\n")
}

func Test_get_option_from_command_without_options(t *testing.T) {
	options := service.GetCommandFields(mockCommandWithoutOptions{})
	require.Len(t, options, 0)
}

type mockCommandOption struct {
	DryRun bool `flag:"dry-run"`
}

func Test_get_parsed_option(t *testing.T) {
	options := service.GetCommandFields(mockCommandOption{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "bool", support.Name(options[0].Value))
}

type mockCommandMultipleOptions struct {
	DryRun bool `flag:"dry-run"`
	Vvv    bool `flag:"vvv"`
}

func Test_get_parsed_option_multiple_fields(t *testing.T) {
	options := service.GetCommandFields(mockCommandMultipleOptions{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "bool", support.Name(options[0].Value))
	require.Equal(t, "vvv", options[1].Tag.Get("flag"))
	require.Equal(t, "bool", support.Name(options[1].Value))
}

type mockCommandOptions struct {
	DryRun bool `short:"dr" flag:"dry-run"`
}

func Test_get_parsed_multiple_options(t *testing.T) {
	options := service.GetCommandFields(mockCommandOptions{})

	require.Equal(t, "dr", options[0].Tag.Get("short"))
	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
}

type mockCommandOptionBool struct {
	DryRun bool `flag:"dry-run"`
}

func Test_get_parsed_option_bool(t *testing.T) {
	options := service.GetCommandFields(mockCommandOptionBool{})

	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "bool", support.Name(options[0].Value))
}

type mockCommandOptionString struct {
	Username string `flag:"username"`
}

func Test_get_parsed_option_string(t *testing.T) {
	options := service.GetCommandFields(mockCommandOptionString{})

	require.Equal(t, "username", options[0].Tag.Get("flag"))
	require.Equal(t, "string", support.Name(options[0].Value))
}

type mockCommandOptionInt struct {
	Amount int `flag:"amount"`
}

func Test_get_parsed_option_int(t *testing.T) {
	options := service.GetCommandFields(mockCommandOptionInt{})

	require.Equal(t, "amount", options[0].Tag.Get("flag"))
	require.Equal(t, "int", support.Name(options[0].Value))
}

type mockCommandOptionFloat struct {
	Number float64 `flag:"number"`
}

func Test_get_parsed_option_float(t *testing.T) {
	options := service.GetCommandFields(mockCommandOptionFloat{})

	require.Equal(t, "number", options[0].Tag.Get("flag"))
	require.Equal(t, "float64", support.Name(options[0].Value))
}

type mockCommandOptionsWithDescription struct {
	DryRun bool `short:"dr" flag:"dry-run" description:"Execute the command as a dry run"`
}

func Test_get_parsed_option_with_description(t *testing.T) {
	options := service.GetCommandFields(mockCommandOptionsWithDescription{})

	require.Equal(t, "dr", options[0].Tag.Get("short"))
	require.Equal(t, "dry-run", options[0].Tag.Get("flag"))
	require.Equal(t, "Execute the command as a dry run", options[0].Tag.Get("description"))
}

func Test_show_help_description_of_wrong_flag(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--fake_flag"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithDescription{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, TrimDoubleSpaces(output.String()), "flag provided but not defined: -fake_flag")
	require.Regexp(t, `-dr --dry-run .*bool.* The flag description`, output.String())
}

func Test_show_help_description_of_wrong_short(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--fake_short"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithDescription{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, TrimDoubleSpaces(output.String()), `flag provided but not defined: -fake_short`)
	require.Regexp(t, `-dr --dry-run .*bool.* The flag description`, output.String())

}

func Test_show_help_with_short_flags_when_multiple_options_are_given(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--invalid"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, output.String(), `flag provided but not defined: -invalid`)
	require.Contains(t, output.String(), "-dr")
	require.Contains(t, output.String(), "-dry-run")
	require.Contains(t, output.String(), "-send-mail")
	require.Contains(t, output.String(), "-sm")
}

func Test_allow_env_file_flag_without_value(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--env-file"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Equal(t, inter.Failure, code, "Output: "+output.String())
	require.Contains(t, output.String(), `flag needs an argument: -env-file`)
}

func Test_allow_env_file_flag_with_value(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--env-file", ".env.testing"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Equal(t, inter.Success, code, "Output: "+output.String())
	require.NotContains(t, output.String(), `flag provided but not defined: -env-file`)
}

func Test_show_env_file_flag_in_help(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--help"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Equal(t, inter.Failure, code, "Output: "+output.String())
	require.Contains(t, output.String(), `--env-file`, output.String())
	require.Contains(t, output.String(), `Run the command with a environment file.`, output.String())
}

func Test_show_if_invalid_command_is_given(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "fake"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, output.String(), `Command provided but not defined: fake`)
}

type structWithRequiredFlag struct {
	DryRun bool `flag:"dry-run" required:"true"`
}

func (s structWithRequiredFlag) Name() string        { return "test" }
func (s structWithRequiredFlag) Description() string { return "test" }
func (s structWithRequiredFlag) Handle(_ inter.Cli) inter.ExitCode {
	return inter.Success
}

func Test_command_with_required_flag(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithRequiredFlag{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, output.String(), "flag is not provided but is required:")
	require.Contains(t, output.String(), "--dry-run")
}

type structWithRequiredShortFlag struct {
	DryRun bool `short:"d" required:"true"`
}

func (s structWithRequiredShortFlag) Name() string        { return "test" }
func (s structWithRequiredShortFlag) Description() string { return "test" }
func (s structWithRequiredShortFlag) Handle(_ inter.Cli) inter.ExitCode {
	return inter.Success
}

func Test_command_with_required_short_flag(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithRequiredShortFlag{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, output.String(), "flag is not provided but is required:")
	require.Contains(t, output.String(), "  -d")
}

type structWithRequiredShortOrLongFlag struct {
	DryRun bool `short:"d" flag:"dry-run" required:"true"`
}

func (s structWithRequiredShortOrLongFlag) Name() string        { return "test" }
func (s structWithRequiredShortOrLongFlag) Description() string { return "test" }
func (s structWithRequiredShortOrLongFlag) Handle(_ inter.Cli) inter.ExitCode {
	return inter.Success
}

func Test_command_with_required_short_or_long_flag(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithRequiredShortOrLongFlag{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, output.String(), "  flag is not provided but is required:")
	require.Contains(t, output.String(), " -d --dry-run ")
	require.Contains(t, output.String(), "bool required")
}
