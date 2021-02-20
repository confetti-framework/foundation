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

type structWithOptionBool struct {
	DryRun bool `flag:"dry-run"`
}

func (s structWithOptionBool) Name() string        { return "test" }
func (s structWithOptionBool) Description() string { return "test" }

func (s structWithOptionBool) Handle(app inter.App, writer io.Writer) inter.ExitCode {
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
		Commands: []inter.Command{structWithOptionBool{}},
	}.Handle()

	require.Contains(t, output.String(), `true`)
}

func Test_cast_multiple_fields_true(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run", "--send-mail"})

	console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Contains(t, output.String(), "start\nDryRun\nSendMail\nend")
}

func Test_cast_multiple_fields_one_true(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run"})

	console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Contains(t, output.String(), "start\nDryRun\nend")
}

type structWithMultipleFlags struct {
	DryRun bool `short:"dr" flag:"dry-run"`
}

func (s structWithMultipleFlags) Name() string        { return "test" }
func (s structWithMultipleFlags) Description() string { return "test" }
func (s structWithMultipleFlags) Handle(app inter.App, writer io.Writer) inter.ExitCode {
	if s.DryRun {
		_, _ = fmt.Fprintln(writer, "true")
	} else {
		_, _ = fmt.Fprintln(writer, "false")
	}

	return inter.Success
}

func Test_cast_short_flag(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "-dr"})

	console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithMultipleFlags{}},
	}.Handle()

	require.Contains(t, output.String(), `true`)
}

type structWithDescription struct {
	DryRun bool `short:"dr" flag:"dry-run" description:"The flag description"`
}

func setUp() (bytes.Buffer, inter.App) {
	var output bytes.Buffer

	app := foundation.NewTestApp(func(container inter.Container) inter.Container {
		var osArgs []interface{}
		container.Bind("config.App.OsArgs", osArgs)
		container.Bind("config.App.Name", "Confetti")
		container.Bind("config.App.Env", "testing")
		return container
	})

	return output, app
}
