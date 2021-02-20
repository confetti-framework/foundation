package service

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"io"
	"reflect"
	"strings"
)

func DispatchCommands(app inter.App, output io.Writer, commands []inter.Command) inter.ExitCode {
	for _, command := range commands {
		code := handleCommand(app, output, command)
		if code != inter.Continue {
			return code
		}
	}

	return inter.Help
}

func handleCommand(app inter.App, output io.Writer, command inter.Command) inter.ExitCode {
	flagSet := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
	flagSet.SetOutput(output)
	actualArgs := actualArgs(app)

	if actualCommandName(actualArgs) == command.Name() {
		options := GetOptions(command)

		registerOptions(flagSet, options)
		code := validate(flagSet, actualArgs)
		if code != inter.Continue {
			return code
		}
		setValuesInCommand(&command, options, actualArgs)

		return command.Handle(app, output)
	}

	return inter.Continue
}

func setValuesInCommand(command *inter.Command, options []ParsedOption, actualArgs []string) {
	for _, option := range options {
		for _, actual := range actualArgs[2:] {
			actual := strings.TrimLeft(actual, "-")
			if option.Tag.Get("short") == actual {
				setValue(command, option.Number)
			}
			if option.Tag.Get("flag") == actual {
				setValue(command, option.Number)
			}
		}
	}
}

func actualCommandName(appArgs []string) string {
	if len(appArgs) > 1 {
		return appArgs[1]
	}
	return ""
}

func validate(flagSet *flag.FlagSet, appArgs []string) inter.ExitCode {
	var err error = nil
	if len(appArgs) > 1 {
		err = flagSet.Parse(appArgs[2:])
	}

	switch {
	case errors.Is(err, flag.ErrHelp):
		return inter.Help
	case err != nil:
		return inter.Failure
	default:
		return inter.Continue
	}
}

func registerOptions(set *flag.FlagSet, options []ParsedOption) {
	for _, option := range options {
		shortFlag := option.Tag.Get("short")
		if set.Lookup(shortFlag) == nil {
			set.Bool(shortFlag, false, "")
		}
		longFlag := option.Tag.Get("flag")
		if set.Lookup(longFlag) == nil {
			set.Bool(longFlag, false, option.Tag.Get("description"))
		}
	}
}

func setValue(command *inter.Command, i int) {
	// v is the interface{}
	v := reflect.ValueOf(command).Elem()

	// Allocate a temporary variable with type of the struct.
	// v.Elem() is the value contained in the interface.
	tmp := reflect.New(v.Elem().Type()).Elem()

	// Copy the struct value contained in interface to
	// the temporary variable.
	tmp.Set(v.Elem())

	// Set the field.
	tmp.Field(i).SetBool(true)

	// Set the interface to the modified struct value.
	v.Set(tmp)
}

// actualArgs converts arguments from config []interface{} to []string
func actualArgs(app inter.App) []string {
	var result []string
	argsRaw := app.Make("config.App.OsArgs")
	for _, argRaw := range argsRaw.([]interface{}) {
		result = append(result, argRaw.(string))
	}

	return result
}
