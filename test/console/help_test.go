package console

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_help_command(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main",  "the_name", "-h"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithOptionBool{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, output.String(), "\n  the_name")
}
