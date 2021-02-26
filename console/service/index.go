package service

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"sort"
)

var style = table.Style{
	Name: "Custom",
	Box: table.BoxStyle{
		BottomLeft:       " ",
		BottomRight:      " ",
		BottomSeparator:  " ",
		EmptySeparator:   " ",
		Left:             " ",
		LeftSeparator:    " ",
		MiddleHorizontal: " ",
		MiddleSeparator:  " ",
		MiddleVertical:   " ",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            " ",
		RightSeparator:   " ",
		TopLeft:          " ",
		TopRight:         " ",
		TopSeparator:     " ",
		UnfinishedRow:    " ~",
	},
	Color:  table.ColorOptionsDefault,
	Format: table.FormatOptionsDefault,
	HTML:   table.DefaultHTMLOptions,
	Options: table.Options{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    false,
	},
	Title: table.TitleOptionsBright,
}

func RenderIndex(app inter.App, writer io.Writer, commands []inter.Command) inter.ExitCode {
	// Setup the writers
	t := table.NewWriter()
	t.SetOutputMirror(writer)

	// Add title and instruction for global usage
	_, _ = fmt.Fprintf(
		writer,
		"\n  %s (%s)\n",
		app.Make("config.App.Name").(string),
		app.Make("config.App.Env").(string),
	)
	t.AppendRow([]interface{}{"-h --help", "Can be used with any command to show\nthe command's available arguments and options."})
	t.AppendRow([]interface{}{" "})

	// Sort all rows
	sort.SliceStable(commands, func(i, c int) bool {
		return commands[i].Name() < commands[c].Name()
	})

	// Add all rows
	for _, command := range commands {
		t.AppendRow([]interface{}{command.Name(), command.Description()})
	}

	t.SetStyle(style)
	t.Render()

	return inter.Success
}
