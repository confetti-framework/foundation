package console

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_validate_invalid_option(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--invalid"})

	kernel := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithOptionBool{}},
	}
	code := kernel.Handle()

	require.Equal(t, inter.Failure, code)
}

type StructWithNoFlag struct{}

func (s StructWithNoFlag) Name() string        { return "test" }
func (s StructWithNoFlag) Description() string { return "test" }

func (s StructWithNoFlag) Handle(c inter.Cli) inter.ExitCode {
	return inter.Success
}

func Test_validate_invalid_option_on_empty_struct(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--invalid"})

	code := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{StructWithNoFlag{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
}

func Test_validate_one_option(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run"})

	kernel := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithOptionBool{}},
	}

	require.NotPanics(t, func() {
		kernel.Handle()
	})
}

type structWithMultipleFields struct {
	DryRun   bool `short:"dr" flag:"dry-run"`
	SendMail bool `short:"sm" flag:"send-mail"`
}
func (s structWithMultipleFields) Name() string        { return "test" }
func (s structWithMultipleFields) Description() string { return "test" }
func (s structWithMultipleFields) Handle(c inter.Cli) inter.ExitCode {
	_, _ = fmt.Fprintln(c.Writer(), "start")
	if s.DryRun {
		_, _ = fmt.Fprintln(c.Writer(), "DryRun")
	}
	if s.SendMail {
		_, _ = fmt.Fprintln(c.Writer(), "SendMail")
	}
	_, _ = fmt.Fprintln(c.Writer(), "end")

	return inter.Success
}

func Test_validate_multiple_fields(t *testing.T) {
	output, app := setUp()

	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run", "--send-mail"})

	kernel := console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}

	require.NotPanics(t, func() {
		kernel.Handle()
	})
}
