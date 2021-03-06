package console

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"regexp"
	"strings"
	"testing"
)

func Test_index_show_title(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{App: app, Writer: &output}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, output.String(), "Confetti (testing)")
}

func Test_index_with_one_command(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{
		App:    app,
		Writer: &output,
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(
		t,
		TrimDoubleSpaces(output.String()),
		"Confetti (testing)\x1b[39m"+
			"\n\n"+
			" -h --help Can be used with any command to show\n"+
			" the command's available arguments and options.\n\n"+
			" baker Interact with your application.\n"+
			" log:clear Clear the log files as indicated in the configuration.",
	)
}

// todo fix test
func Test_command_suggestions_on_failed_command(t *testing.T) {
	output, app := setUp()
	kernel := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{aCommand{}, bCommand{}},
	}
	app.Bind("config.App.OsArgs", []interface{}{"/main", "com"})

	code := kernel.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(
		t,
		TrimDoubleSpaces(output.String()),
		"command provided but not defined: com\x1b[39m\n\x1b[31m\n"+
			"Do you mean one of these?\x1b[39m\n\x1b[31m"+
			"\ta_command\x1b[39m\n\x1b[31m"+
			"\tb_command\x1b[39m\n\x1b[32m\x1b[39m",
	)

	output.Reset()
	app.Bind("config.App.OsArgs", []interface{}{"/main", "a_com", "\ny\n"})
	//fmt.Fprint(kernel.Writer, "y\n")

	kernel.Handle()
	require.Contains(t, output.String(), "command a done")
	//stdin
}

type aCommand struct {
	DryRun bool `flag:"-dry-run"`
}

func (s aCommand) Name() string        { return "a_command" }
func (s aCommand) Description() string { return "" }
func (s aCommand) Handle(c inter.Cli) inter.ExitCode {
	c.Info("command a done")
	return inter.Success
}

type bCommand struct{}

func (s bCommand) Name() string        { return "b_command" }
func (s bCommand) Description() string { return "" }
func (s bCommand) Handle(_ inter.Cli) inter.ExitCode {
	return inter.Success
}

func Test_index_in_correct_order(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{console.LogClear{}, aCommand{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Regexp(t, "(?s)a_command.*log", output.String())
}

func TrimDoubleSpaces(value string) string {
	// Replace double spaces
	regex := regexp.MustCompile(` {2,}`)
	value = regex.ReplaceAllString(strings.Trim(value, " "), " ")

	// replace newline with only one space
	regex = regexp.MustCompile(` \n`)
	value = regex.ReplaceAllString(value, "\n")

	return value
}
