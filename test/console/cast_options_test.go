package console

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/console"
	"github.com/confetti-framework/foundation/console/facade"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
)

type structWithOptionBool struct {
	DryRun bool `flag:"dry-run"`
}

func (s structWithOptionBool) Name() string        { return "the_name" }
func (s structWithOptionBool) Description() string { return "the_description" }

func (s structWithOptionBool) Handle(c inter.Cli) inter.ExitCode {
	if s.DryRun {
		_, _ = fmt.Fprintln(c.Writer(), "true")
	} else {
		_, _ = fmt.Fprintln(c.Writer(), "false")
	}

	return inter.Success
}

func Test_cast_multiple_fields_true(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run", "--send-mail"})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Contains(t, output.String(), "start\nDryRun\nSendMail\nend")
}

func Test_cast_multiple_fields_one_true(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--dry-run"})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFields{}},
	}.Handle()

	require.Contains(t, output.String(), "start\nDryRun\nend")
}

func Test_cast_option_bool_true(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "the_name", "--dry-run"})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithOptionBool{}},
	}.Handle()

	require.Contains(t, output.String(), `true`)
}

type structWithOptionString struct {
	Username string `flag:"username"`
}

func (s structWithOptionString) Name() string        { return "test" }
func (s structWithOptionString) Description() string { return "test" }

func (s structWithOptionString) Handle(c inter.Cli) inter.ExitCode {
	_, _ = fmt.Fprintln(c.Writer(), "name:"+s.Username)

	return inter.Success
}

func Test_cast_option_string(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--username", "viktor"})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithOptionString{}},
	}.Handle()

	require.Contains(t, output.String(), `name:viktor`)
}

type structWithOptionDuration struct {
	Duration time.Duration `flag:"duration"`
}

func (s structWithOptionDuration) Name() string        { return "test" }
func (s structWithOptionDuration) Description() string { return "test" }

func (s structWithOptionDuration) Handle(c inter.Cli) inter.ExitCode {
	_, _ = fmt.Fprintln(c.Writer(), "seconds:"+cast.ToString(s.Duration.Seconds()))

	return inter.Success
}

func Test_cast_option_duration(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "--duration", "2m"})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithOptionDuration{}},
	}.Handle()

	require.Contains(t, output.String(), `seconds:120`)
}

type structWithOptionInt struct {
	Int    int    `flag:"int"`
	Int63  int64  `flag:"int64"`
	Uint   uint   `flag:"uint"`
	Uint64 uint64 `flag:"uint64"`
}

func (s structWithOptionInt) Name() string        { return "test" }
func (s structWithOptionInt) Description() string { return "test" }

func (s structWithOptionInt) Handle(c inter.Cli) inter.ExitCode {
	_, _ = fmt.Fprintln(c.Writer(), "int:"+strconv.Itoa(s.Int))
	_, _ = fmt.Fprintln(c.Writer(), "int64:"+cast.ToString(s.Int63))
	_, _ = fmt.Fprintln(c.Writer(), "uint:"+cast.ToString(s.Uint))
	_, _ = fmt.Fprintln(c.Writer(), "uint64:"+cast.ToString(s.Uint64))

	return inter.Success
}

func Test_cast_option_int(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{
		"/exe/main", "test",
		"--int", "12",
		"--int64", "9223372036854775806",
		"--uint", "18446744073709551615",
		"--uint64", "18446744073709551615",
	})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithOptionInt{}},
	}.Handle()

	require.Contains(t, output.String(), `int:12`)
	require.Contains(t, output.String(), `int64:9223372036854775806`)
	require.Contains(t, output.String(), `uint:18446744073709551615`)
	require.Contains(t, output.String(), `uint64:18446744073709551615`)
}

type structWithDoubleOptionSameType struct {
	Int1 int `flag:"int1"`
	Int2 int `flag:"int2"`
}

func (s structWithDoubleOptionSameType) Name() string        { return "test" }
func (s structWithDoubleOptionSameType) Description() string { return "test" }

func (s structWithDoubleOptionSameType) Handle(c inter.Cli) inter.ExitCode {
	_, _ = fmt.Fprintln(c.Writer(), "int1:"+strconv.Itoa(s.Int1))
	_, _ = fmt.Fprintln(c.Writer(), "int2:"+strconv.Itoa(s.Int2))

	return inter.Success
}

func Test_cast_double_option_same_type(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{
		"/exe/main", "test",
		"--int1", "111",
		"--int2", "222",
	})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithDoubleOptionSameType{}},
	}.Handle()

	require.Contains(t, output.String(), `int1:111`)
	require.Contains(t, output.String(), `int2:222`)
}

type structWithMultipleFlags struct {
	DryRun bool `short:"dr" flag:"dry-run"`
}

func (s structWithMultipleFlags) Name() string        { return "test" }
func (s structWithMultipleFlags) Description() string { return "test" }
func (s structWithMultipleFlags) Handle(c inter.Cli) inter.ExitCode {
	if s.DryRun {
		_, _ = fmt.Fprintln(c.Writer(), "true")
	} else {
		_, _ = fmt.Fprintln(c.Writer(), "false")
	}

	return inter.Success
}

func Test_cast_short_flag(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test", "-dr"})

	console.Kernel{
		App:      app,
		Writer:   &output,
		Commands: []inter.Command{structWithMultipleFlags{}},
	}.Handle()

	require.Contains(t, output.String(), `true`)
}

type structWithDash struct {
	DryRun bool `flag:"-dry-run"`
}

func (s structWithDash) Name() string        { return "test:with_dash" }
func (s structWithDash) Description() string { return "test" }
func (s structWithDash) Handle(_ inter.Cli) inter.ExitCode {
	return inter.Success
}

func Test_cast_flag_with_dash(t *testing.T) {
	output, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/exe/main", "test:with_dash", "-dr"})

	require.Panics(t, func() {
		console.Kernel{
			App:      app,
			Writer:   &output,
			Commands: []inter.Command{structWithDash{}},
		}.Handle()
	})
}

type structWithDescription struct {
	DryRun      bool `short:"dr" flag:"dry-run" description:"The flag description"`
	CommandName string
}

func (s structWithDescription) Name() string {
	if s.CommandName == "" {
		return "test"
	}
	return s.CommandName
}
func (s structWithDescription) Description() string { return "test" }
func (s structWithDescription) Handle(_ inter.Cli) inter.ExitCode {
	return inter.Success
}

func setUp() (bytes.Buffer, inter.App) {
	var writer bytes.Buffer

	app := foundation.NewTestApp(func(container inter.Container) inter.Container {
		container.Bind("config.App.OsArgs", []interface{}{})
		container.Bind("config.App.Name", "Confetti")
		container.Bind("config.App.Env", "testing")
		return container
	})

	return writer, app
}

func Test_programmatically_executing_commands(t *testing.T) {
	_, app := setUp()
	exitCode := structWithDescription{}.Handle(facade.NewCli(app))

	require.Equal(t, inter.Success, exitCode)
}
