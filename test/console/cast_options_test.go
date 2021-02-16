package console

import (
	"bytes"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

type StructWithOptionBool struct {
	DryRun bool `flag:"dry-run"`
}

func (s StructWithOptionBool) Name() string        { return "test" }
func (s StructWithOptionBool) Description() string { return "test" }

func (s StructWithOptionBool) Handle(app inter.App, writer io.Writer) inter.ExitCode {
	if s.DryRun {
		fmt.Fprintln(writer, "true")
	} else {
		fmt.Fprintln(writer, "false")
	}

	return inter.Success
}

func Test_cast_option_bool_true(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run"})

	console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{StructWithOptionBool{}},
	}.Handle()

	require.Contains(t, output.String(), `true`)
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
