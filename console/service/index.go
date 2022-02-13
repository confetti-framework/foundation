package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/confetti-framework/contract/inter"
	"github.com/jedib0t/go-pretty/v6/table"
)

var globalOptions = map[string]string{
	"-h --help":  "Show the command's available arguments.",
	"--env-file": "Run the command with a defined environment file.",
}

func RenderIndex(c inter.Cli, commands []inter.Command) inter.ExitCode {
	// Add title and instruction for global usage
	name := c.App().Make("config.App.Name").(string)
	c.Line("\n  %s", name)

	// Sort all commands
	sort.SliceStable(commands, func(i, c int) bool {
		return commands[i].Name() < commands[c].Name()
	})

	t := c.Table()

	t.AppendRow([]interface{}{"\u001B[32mGlobal options:\u001B[0m"})
	printGlobalOptions(t)

	t.AppendRow([]interface{}{"\n\u001B[32mAvailable commands:\u001B[0m"})
	printCommandsByGroup(t, commands)

	t.Render()

	return inter.Success
}

func printGlobalOptions(t table.Writer) {
	for flag, description := range globalOptions {
		t.AppendRow([]interface{}{" " + flag, description})
	}
}

func printCommandsByGroup(t table.Writer, commands []inter.Command) {
	// Commands without a group
	for _, command := range commands {
		if !strings.ContainsRune(command.Name(), ':') {
			t.AppendRow([]interface{}{fmt.Sprintf("  \u001B[33m%s\u001B[0m", command.Name()), command.Description()})
		}
	}

	// Commands with a group (e.g. app:serve)
	for _, groupName := range getGroupNames(commands) {
		t.AppendRow([]interface{}{fmt.Sprintf(" \u001B[32m%s\u001B[0m", groupName)})

		for _, command := range commands {
			if strings.HasPrefix(command.Name(), groupName+":") {
				t.AppendRow([]interface{}{fmt.Sprintf("  \u001B[33m%s\u001B[0m", command.Name()), command.Description()})
			}
		}
	}
}

func getGroupNames(commands []inter.Command) (names []string) {
	groups := map[string]bool{}

	// Find all group names (unique!)
	for _, command := range commands {
		parts := strings.Split(command.Name(), ":")

		if len(parts) <= 1 {
			continue
		}

		groupName := parts[0]

		groups[groupName] = true
	}

	// Reduce to keys
	for groupName := range groups {
		names = append(names, groupName)
	}

	// Ensure groups are sorted, just like commands
	sort.SliceStable(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	return
}
