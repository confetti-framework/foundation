package service

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
)

var Style = table.Style{
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

func RenderIndex(output io.Writer, commands []inter.Command) inter.ExitCode {
	t := table.NewWriter()
	t.SetOutputMirror(output)
	t.SetTitle("Confetti Baker")
	t.AppendHeader(table.Row{"Command", "Description"})

	for _, command := range commands {
		t.AppendRow([]interface{}{command.Name(), command.Description()})
	}

	t.SetStyle(Style)
	t.Render()

	return inter.Success
}
