//go:build !race

package console

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/confetti-framework/foundation/console/facade"
	"github.com/confetti-framework/foundation/console/service"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_index_show_title(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{App: app, Writer: &output}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, output.String(), "Confetti")
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
		"\x1b[39m\n Confetti\x1b[39m\n\n"+
			" \x1b[32mGlobal options:\u001B[0m\n"+
			" -h --help Show the command's available arguments.\n"+
			" --env-file Run the command with a defined environment file.\n"+
			"\n",
		TrimDoubleSpaces(output.String()),
	)
}

func Test_command_suggestions_on_failed_command(t *testing.T) {
	commands := []inter.Command{aCommand{}, bCommand{}}

	t.Run("check with mutliple suggestions", func(t *testing.T) {
		output, app := setUp()
		app.Bind("config.App.OsArgs", []interface{}{"/main", "com"})
		rc := ioutil.NopCloser(strings.NewReader(""))
		cli := facade.NewCliByReadersAndWriter(app, rc, &output, nil)

		code := service.DispatchCommands(cli, commands, []flag.Getter{})

		require.Equal(t, inter.Failure, code)
		require.Contains(
			t,
			TrimDoubleSpaces(output.String()),
			"Command provided but not defined: com\x1b[39m\n\x1b[31m\n"+
				"Do you mean one of these?\x1b[39m\n\x1b[31m"+
				"\ta_command\x1b[39m\n\x1b[31m"+
				"\tb_command\x1b[39m\n\x1b[32m\x1b[39m",
		)
	})

	t.Run("check with single suggestion which is executed", func(t *testing.T) {
		output, app := setUp()
		app.Bind("config.App.OsArgs", []interface{}{"/main", "a_com"})
		rc := ioutil.NopCloser(strings.NewReader("y\n"))
		cli := facade.NewCliByReadersAndWriter(app, rc, &output, nil)

		code := service.DispatchCommands(cli, commands, []flag.Getter{})

		require.Equal(t, inter.Success, code)
		require.Contains(t, output.String(), "command a done")
	})

	t.Run("check with single suggestion which is not executed", func(t *testing.T) {
		output, app := setUp()
		app.Bind("config.App.OsArgs", []interface{}{"/main", "a_com"})
		rc := ioutil.NopCloser(strings.NewReader("\n"))
		cli := facade.NewCliByReadersAndWriter(app, rc, &output, nil)

		code := service.DispatchCommands(cli, commands, []flag.Getter{})

		require.Equal(t, inter.Failure, code)
		require.NotContains(t, output.String(), "command a done")
	})

	t.Run("check with none suggestion which is not executed", func(t *testing.T) {
		output, app := setUp()
		app.Bind("config.App.OsArgs", []interface{}{"/main", "none"})
		rc := ioutil.NopCloser(strings.NewReader("\n"))
		cli := facade.NewCliByReadersAndWriter(app, rc, &output, nil)

		code := service.DispatchCommands(cli, commands, []flag.Getter{})

		require.Equal(t, inter.Failure, code)
		require.NotContains(t, output.String(), "Do you mean one of these?")
	})
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
