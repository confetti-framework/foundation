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
	code := console.Kernel{App: app, Output: &output}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, output.String(), "Confetti (testing)")
}

func Test_index_with_one_command(t *testing.T) {
	output, app := setUp()
	code := console.Kernel{
		App: app,
		Output:   &output,
		Commands: []inter.Command{console.LogClear{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t,
		TrimDoubleSpaces(output.String()),`
 COMMAND DESCRIPTION

 log:clear Clear the log files as indicated in the configuration.

`)
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
