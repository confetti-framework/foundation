package service

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/flag_type"
	"github.com/confetti-framework/support"
)

func ActualCommandName(appArgs []string) string {
	if len(appArgs) > 1 {
		return appArgs[1]
	}
	return ""
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

func registerFlags(set *flag.FlagSet, options []Field, flagProviders []func() []flag.Getter) {
	for _, option := range options {
		fg := flagGettersByProviders(flagProviders)
		registerFlag(set, option, option.Tag.Get("short"), "", fg)
		registerFlag(set, option, option.Tag.Get("flag"), option.Tag.Get("description"), fg)
	}
	set.Var(new(flag_type.String), "env-file", "Can be used with any command to run the command with a defined environment file.")
}

func registerFlag(set *flag.FlagSet, option Field, flag string, description string, getters []flag.Getter) {
	// Check if the flag is already present.
	if set.Lookup(flag) != nil || flag == "" {
		return
	}

	// There is no need to add a - to a tag. This will be added automatically.
	if flag[0] == '-' {
		panic("field with tag `" + flag + "` starts with a -. That's not allowed.")
	}

	for _, getter := range getters {
		getterName := support.Name(getter.Get())
		if getterName == support.Name(option.Value) {
			set.Var(getter, flag, description)
		}
	}
}
