package console

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/facade"
	"github.com/confetti-framework/foundation/console/flag_type"
	"github.com/confetti-framework/foundation/console/service"
	"io"
)

// These values must be in a callback so that they are new
// for each field. If there are multiple fields with the same
// type, they would use the same pointer.
var flagGetters = func() []flag.Getter {
	return []flag.Getter{
		new(flag_type.Bool),
		new(flag_type.String),
		new(flag_type.Int),
		new(flag_type.Int64),
		new(flag_type.Uint),
		new(flag_type.Uint64),
		new(flag_type.Duration),
	}
}

var commands = []inter.Command{
	Baker{},
}

type Kernel struct {
	App           inter.App
	Writer        io.Writer
	WriterErr     io.Writer
	Commands      []inter.Command
	FlagProviders []func() []flag.Getter
}

func (k Kernel) Handle() inter.ExitCode {
	k.Commands = append(k.Commands, commands...)
	k.FlagProviders = append(k.FlagProviders, flagGetters)

	cli := facade.NewCli(k.App, k.Writer, k.WriterErr)
	code := service.DispatchCommands(cli, k.Commands, k.FlagProviders)
	if code != inter.Index {
		return code
	}

	return service.RenderIndex(cli, k.Commands)
}

func (k Kernel) GetCommands() []inter.Command {
	return k.Commands
}

func (k Kernel) GetFlagProviders() []func() []flag.Getter {
	return k.FlagProviders
}
