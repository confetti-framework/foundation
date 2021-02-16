package commands

import (
	"bytes"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"regexp"
	"strings"
	"testing"
)

func Test_index_show_title(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{App: app, Output: &output}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, output.String(), "Confetti")
}

func Test_index_with_one_command(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{
		App: app,
		Output:   &output,
		Commands: []inter.Command{console.AppServe{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t,
		trimDoubleSpaces(output.String()),`
 COMMAND DESCRIPTION

 app:serve Start the http server to handle requests

`)
}

func setUp() (bytes.Buffer, inter.App) {
	var output bytes.Buffer

	app := foundation.NewTestApp(func(container inter.Container) inter.Container {
		var osArgs []interface{}
		container.Bind("config.App.OsArgs", osArgs)
		container.Bind("config.App.Name", "test_app")
		return container
	})

	return output, app
}

func trimDoubleSpaces(value string) string {
	// Replace double spaces
	regex := regexp.MustCompile(` {2,}`)
	value = regex.ReplaceAllString(strings.Trim(value, " "), " ")

	// replace newline with only one space
	regex = regexp.MustCompile(` \n`)
	value = regex.ReplaceAllString(value, "\n")

	return value
}
