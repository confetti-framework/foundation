package service

import (
	"flag"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"io"
	"reflect"
	"strings"
)

const flagShort = "short"
const flagLong = "flag"

func DispatchCommands(
	app inter.App,
	w io.Writer,
	commands []inter.Command,
	flagProviders []func() []flag.Getter,
) inter.ExitCode {
	for _, command := range commands {
		code := handleCommand(app, w, command, flagProviders)
		if code != inter.Continue {
			return code
		}
	}

	args := actualArgs(app)
	if len(args) > 1 {
		fmt.Fprintf(w, "command provided but not defined: %s\n", ActualCommandName(args))
		return inter.Failure
	}

	return inter.Index
}

func handleCommand(
	app inter.App,
	w io.Writer,
	command inter.Command,
	flagProviders []func() []flag.Getter,
) inter.ExitCode {
	set := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
	set.SetOutput(w)

	actualArgs := actualArgs(app)

	if ActualCommandName(actualArgs) == command.Name() {
		options := GetOptions(command)
		set.Usage = helpFormat(w, command, options)

		registerFlags(set, options, flagProviders)
		code := parse(set, actualArgs)
		if code != inter.Continue {
			return code
		}
		code = validateRequiredFields(w, set, options)
		if code != inter.Continue {
			return code
		}
		setValuesInCommand(&command, set, options)

		return command.Handle(app, w)
	}

	return inter.Continue
}

func validateRequiredFields(w io.Writer, set *flag.FlagSet, options []Field) inter.ExitCode {
	for _, option := range options {
		if option.Tag.Get("required") != "true" {
			continue
		}

		short := getActualValue(set, option, flagShort)
		long := getActualValue(set, option, flagLong)
		if isEqualOrNil(option.Value, short) && isEqualOrNil(option.Value, long) {
			_, _ = fmt.Fprintf(
				w,
				"\n  flag is not provided but is required:\n\n  %s \u001B[30;1m%T\u001B[0m\n\n",
				flagsFormat(option),
				option.Value,
			)

			return inter.Failure
		}
	}
	return inter.Continue
}

func flagsFormat(option Field) string {
	var result string
	if option.Tag.Get(flagShort) != "" {
		result += "-" + option.Tag.Get(flagShort)
	}
	result += " "
	if option.Tag.Get(flagLong) != "" {
		result += "--" + option.Tag.Get(flagLong)
	}
	return strings.Trim(result, " ")
}

func isEqualOrNil(field interface{}, actual interface{}) bool {
	return actual == nil || field == actual
}

func getActualValue(set *flag.FlagSet, option Field, key string) interface{} {
	lookup := set.Lookup(option.Tag.Get(key))
	if lookup == nil {
		return nil
	}
	return lookup.Value.(flag.Getter).Get()
}

func setValuesInCommand(command *inter.Command, set *flag.FlagSet, options []Field) {
	for _, option := range options {
		short := getActualValue(set, option, flagShort)
		if short != nil {
			setValueByFlag(command, option.Number, short)
		}
		long := getActualValue(set, option, flagLong)
		if long != nil {
			setValueByFlag(command, option.Number, long)
		}
	}
}

func parse(flagSet *flag.FlagSet, appArgs []string) inter.ExitCode {
	var err error = nil
	if len(appArgs) > 1 {
		err = flagSet.Parse(appArgs[2:])
	}
	if err != nil {
		return inter.Failure
	}
	return inter.Continue
}

func flagGettersByProviders(providers []func() []flag.Getter) []flag.Getter {
	var result []flag.Getter
	for _, provider := range providers {
		result = append(result, provider()...)
	}
	return result
}

func setValueByFlag(command *inter.Command, i int, rawValue interface{}) {
	// v is the interface{}
	v := reflect.ValueOf(command).Elem()

	// Allocate a temporary variable with type of the struct.
	// v.Elem() is the value contained in the interface.
	tmp := reflect.New(v.Elem().Type()).Elem()

	// Copy the struct value contained in interface to
	// the temporary variable.
	tmp.Set(v.Elem())

	// Set the field.
	value := reflect.ValueOf(rawValue)
	tmp.Field(i).Set(value)

	// Set the interface to the modified struct value.
	v.Set(tmp)
}
