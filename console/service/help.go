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
		for _, option := range options {
			t.AppendRow(helpFormatFlag(option))
		}

		t.Render()
	}
}

func helpFormatFlag(f Field) table.Row {
	var flags string
	short := f.Tag.Get(inter.Short)
	long := f.Tag.Get(inter.Flag)

	if short != "" {
		flags += fmt.Sprintf("-%s ", short)
	}
	if long != "" {
		flags += fmt.Sprintf("--%s ", long)
	}


	return []interface{}{
		"\u001B[0m" + flags + "\u001b[0m",
		"\t\u001b[30;1m" + typeFormat(f) + "\u001B[0m",
		f.Tag.Get(inter.Description),
	}
}

func typeFormat(f Field) string {
	t := support.Name(f.Value)
	if f.Tag.Get("required") == "true" {
		t = t + " required"
	}
	return t
}
