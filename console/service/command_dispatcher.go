package service

import (
	"flag"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/console/flag_type"
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
		setValuesInCommand(&command, flagSet, options, actualArgs)

		return command.Handle(app, output)
	}

	return inter.Continue
}

func setValuesInCommand(command *inter.Command, set *flag.FlagSet, options []ParsedOption, actualArgs []string) {
	for _, option := range options {
		for _, actual := range actualArgs[2:] {
			actual := strings.TrimLeft(actual, "-")
			if actual == option.Tag.Get("short") {
				setValueByFlag(command, option.Number, set.Lookup(option.Tag.Get("short")))
			} else if actual == option.Tag.Get("flag") {
				setValueByFlag(command, option.Number, set.Lookup(option.Tag.Get("flag")))
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
		registerFlag(set, option, option.Tag.Get("short"), "")
		registerFlag(set, option, option.Tag.Get("flag"), option.Tag.Get("description"))
	}
}

func registerFlag(set *flag.FlagSet, option ParsedOption, flag string, description string) interface{} {
	var result interface{}

	// Check if the flag is already present.
	if set.Lookup(flag) != nil || flag == "" {
		return nil
	}

	// special case: bool doesn't need an arg
	if _, ok := option.Value.(bool); ok {
		value := flag_type.BoolValue(false)
		set.Var(&value, flag, description)
		result = value
	} else {
		value := flag_type.String("")
		set.Var(&value, flag, description)
		result = value
	}

	return result
}

func setValueByFlag(command *inter.Command, i int, lookup *flag.Flag) {
	// v is the interface{}
	v := reflect.ValueOf(command).Elem()

	// Allocate a temporary variable with type of the struct.
	// v.Elem() is the value contained in the interface.
	tmp := reflect.New(v.Elem().Type()).Elem()

	// Copy the struct value contained in interface to
	// the temporary variable.
	tmp.Set(v.Elem())

	// Set the field.
	rawValue, ok := lookup.Value.(interface{ Get() interface{} })
	if ok == false {
		panic(fmt.Sprintf("Can't get value from flag type: Method `Get() interface{}` not found in %v", lookup))
	}
	value := reflect.ValueOf(rawValue.Get())
	tmp.Field(i).Set(value)

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
