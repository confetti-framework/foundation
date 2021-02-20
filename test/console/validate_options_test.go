package console

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func Test_validate_invalid_option(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--invalid"})

	kernel := console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithOptionBool{}},
	}
	code := kernel.Handle()

	require.Equal(t, inter.Failure, code)
}

type StructWithNoFlag struct{}

func (s StructWithNoFlag) Name() string        { return "test" }
func (s StructWithNoFlag) Description() string { return "test" }

func (s StructWithNoFlag) Handle(app inter.App, writer io.Writer) inter.ExitCode {
	return inter.Success
}

func Test_validate_invalid_option_on_empty_struct(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--invalid"})

	code := console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{StructWithNoFlag{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
}

func Test_validate_one_option(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run"})

	kernel := console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithOptionBool{}},
	}

	require.NotPanics(t, func() {
		kernel.Handle()
	})
}

type structWithMultipleFields struct {
	DryRun   bool `flag:"dry-run"`
	SendMail bool `flag:"send-mail"`
}
func (s structWithMultipleFields) Name() string        { return "test" }
func (s structWithMultipleFields) Description() string { return "test" }
func (s structWithMultipleFields) Handle(app inter.App, writer io.Writer) inter.ExitCode {
	_, _ = fmt.Fprintln(writer, "start")
	if s.DryRun {
		_, _ = fmt.Fprintln(writer, "DryRun")
	}
	if s.SendMail {
		_, _ = fmt.Fprintln(writer, "SendMail")
	}
	_, _ = fmt.Fprintln(writer, "end")

	return inter.Success
}

func Test_validate_multiple_fields(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run", "--send-mail"})

	kernel := console.Kernel{
		App:      app,
		Output:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}

	require.NotPanics(t, func() {
		kernel.Handle()
	})
}
