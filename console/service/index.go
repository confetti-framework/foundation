package service

import (
	"github.com/confetti-framework/contract/inter"
	"sort"
)

func RenderIndex(c inter.Cli, commands []inter.Command) inter.ExitCode {
	// Add title and instruction for global usage
	name := c.App().Make("config.App.Name").(string)
	env := c.App().Make("config.App.Env").(string)
	c.Line("\n  %s (%s)", name, env)

	// Sort all commands
	sort.SliceStable(commands, func(i, c int) bool {
		return commands[i].Name() < commands[c].Name()
	})

	t := c.Table()
	t.AppendRow([]interface{}{"-h --help", "Can be used with any command to show\nthe command's available arguments and options."})
	t.AppendRow([]interface{}{" "})
	for _, command := range commands {
		t.AppendRow([]interface{}{command.Name(), command.Description()})
	}

	t.Render()

	return inter.Success
}
