package service

import (
	"github.com/confetti-framework/contract/inter"
	"sort"
)

func RenderIndex(c inter.Cli, commands []inter.Command) inter.ExitCode {
	// Add title and instruction for global usage
	name := c.App().Make("config.App.Name").(string)
	c.Line("\n  %s", name)

	// Sort all commands
	sort.SliceStable(commands, func(i, c int) bool {
		return commands[i].Name() < commands[c].Name()
	})

	t := c.Table()
	t.AppendRow([]interface{}{"\u001b[30;1mGlobal options:"})
	t.AppendRow([]interface{}{"-h --help", "Show the command's available arguments."})
	t.AppendRow([]interface{}{"--env-file", "Run the command with a defined environment file."})
	t.AppendRow([]interface{}{"\u001B[0m "})
	for _, command := range commands {
		t.AppendRow([]interface{}{command.Name(), command.Description()})
	}

	t.Render()

	return inter.Success
}
