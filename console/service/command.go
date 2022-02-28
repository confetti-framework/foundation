package service

import (
	"flag"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"reflect"
	"strings"
)

const flagShort = "short"
const flagLong = "flag"

func DispatchCommands(
	c inter.Cli,
	commands []inter.Command,
	getters []flag.Getter,
) inter.ExitCode {
	for _, command := range commands {
		code := handleCommand(c, command, getters)
		if code != inter.Continue {
			return code
		}
	}

	args := actualArgs(c.App())
	if len(args) > 1 {
		actualCommand := ActualCommandName(args)
		c.Error("Command provided but not defined: %s", actualCommand)
		return suggestCommands(c, actualCommand, commands)
	}

	return inter.Index
}

func suggestCommands(c inter.Cli, actualCommand string, commands []inter.Command) inter.ExitCode {
	//goland:noinspection GoPreferNilSlice
	suggestions := []inter.Command{}
	for _, command := range commands {
		if strings.Contains(command.Name(), actualCommand) {
			suggestions = append(suggestions, command)
		}
	}

	if len(suggestions) > 1 {
		c.Error("\nDo you mean one of these?")
		for _, suggestion := range suggestions {
			c.Error("\t%s", suggestion.Name())
		}
		c.Info("")
	}

	if len(suggestions) == 1 {
		exec := c.Confirm(fmt.Sprintf("Do you want to execute \"%s\"", suggestions[0].Name()), false)
		if exec {
			return suggestions[0].Handle(c)
		}
	}

	return inter.Failure
}

func handleCommand(
	c inter.Cli,
	command inter.Command,
	getters []flag.Getter,
) inter.ExitCode {
	set := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
	set.SetOutput(c.Writer())

	actualArgs := actualArgs(c.App())

	if ActualCommandName(actualArgs) == command.Name() {
		fields := GetCommandFields(command)
		set.Usage = helpFormat(c, command, fields)

		registerFlags(set, fields, getters)
		code := parse(set, actualArgs)
		if code != inter.Continue {
			return code
		}
		code = validateRequiredFields(c, set, fields)
		if code != inter.Continue {
			return code
		}
		setValuesInCommand(&command, set, fields)

		return command.Handle(c)
	}

	return inter.Continue
}

func validateRequiredFields(c inter.Cli, set *flag.FlagSet, fields []Field) inter.ExitCode {
	for _, field := range fields {
		if field.Tag.Get("required") != "true" {
			continue
		}
		short := getActualValue(set, field, flagShort)
		long := getActualValue(set, field, flagLong)
		if isEqualOrNil(field.Value, short) && isEqualOrNil(field.Value, long) {
			c.Error("  flag is not provided but is required:\n")
			c.Line("  %s \u001B[30;1m%s\u001B[0m\n",
				flagsFormat(field),
				typeFormat(field),
			)

			return inter.Failure
		}
	}
	return inter.Continue
}

func flagsFormat(field Field) string {
	var result string
	if field.Tag.Get(flagShort) != "" {
		result += "-" + field.Tag.Get(flagShort)
	}
	result += " "
	if field.Tag.Get(flagLong) != "" {
		result += "--" + field.Tag.Get(flagLong)
	}
	return strings.Trim(result, " ")
}

func isEqualOrNil(field interface{}, actual interface{}) bool {
	return actual == nil || field == actual
}

func getActualValue(set *flag.FlagSet, f Field, key string) interface{} {
	lookup := set.Lookup(f.Tag.Get(key))
	if lookup == nil {
		return nil
	}
	value := lookup.Value.(flag.Getter).Get()
	return value
}

func setValuesInCommand(command *inter.Command, set *flag.FlagSet, fields []Field) {
	for _, f := range fields {
		short := getActualValue(set, f, flagShort)
		if short != nil {
			setValueByFlag(command, f.Number, short)
		}
		long := getActualValue(set, f, flagLong)
		if long != nil {
			setValueByFlag(command, f.Number, long)
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
