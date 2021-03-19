package service

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
	"github.com/jedib0t/go-pretty/v6/table"
)

func helpFormat(c inter.Cli, command inter.Command, options []Field) func() {
	return func() {
		t := c.Table()

		_, _ = fmt.Fprintf(c.Writer(), "\n  %s  \u001B[30;1m%s\n", command.Name(), command.Description())
		t.AppendRow(helpRow("--env-file ", "string", "Run the command with a environment file."))
		for _, option := range options {
			t.AppendRow(helpRowByFlag(option))
		}

		t.Render()
	}
}

func helpRowByFlag(f Field) table.Row {
	var flags string
	short := f.Tag.Get(inter.Short)
	long := f.Tag.Get(inter.Flag)

	if short != "" {
		flags += fmt.Sprintf("-%s ", short)
	}
	if long != "" {
		flags += fmt.Sprintf("--%s ", long)
	}

	return helpRow(flags, typeFormat(f), f.Tag.Get(inter.Description))
}

func helpRow(flags string, typeFormat string, desc string) table.Row {
	return []interface{}{
		"\u001B[0m" + flags + "\u001b[0m",
		"\t\u001b[30;1m" + typeFormat + "\u001B[0m",
		desc,
	}
}

func typeFormat(f Field) string {
	t := support.Name(f.Value)
	if f.Tag.Get("required") == "true" {
		t = t + " required"
	}
	return t
}
