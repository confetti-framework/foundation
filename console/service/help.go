package service

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
)

func helpFormat(writer io.Writer, command inter.Command, options []Field) func() {
	return func() {
		t := table.NewWriter()
		t.SetOutputMirror(writer)

		_, _ = fmt.Fprintf(writer, "\n  %s\n  \u001B[30;1m%s\n", command.Name(), command.Description())
		for _, option := range options {
			t.AppendRow(helpFormatFlag(option))
		}

		t.SetStyle(style)
		t.Render()
	}
}

func helpFormatFlag(f Field) []interface{} {
	short := f.Tag.Get(inter.Short)
	long := f.Tag.Get(inter.Flag)

	var flags string
	if short != "" {
		flags += fmt.Sprintf("-%s ", short)
	}
	if long != "" {
		flags += fmt.Sprintf("--%s ", long)
	}

	return []interface{}{
		"\u001B[0m" + flags + "\u001b[0m",
		"\t\u001b[30;1m" + support.Name(f.Value) + "\u001B[0m",
		f.Tag.Get(inter.Description),
	}
}
