// +build !race

package console

import (
	"github.com/confetti-framework/foundation/console/facade"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strings"
	"testing"
)

const navigateDown = "\u001B[B"

func Test_output_info_with_empty_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Info("")
	require.Equal(t, "\033[32m\033[39m\n", writer.String())
}

func Test_output_info_with_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Info("val")
	require.Equal(t, "\033[32mval\033[39m\n", writer.String())
}

func Test_output_info_with_format_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Info("first %s", "second")
	require.Equal(t, "\033[32mfirst second\033[39m\n", writer.String())
}

func Test_output_error_with_empty_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, nil, &writer)
	cli.Error("")
	require.Equal(t, "\033[31m\033[39m\n", writer.String())
}

func Test_output_error_with_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, nil, &writer)
	cli.Error("val")
	require.Equal(t, "\033[31mval\033[39m\n", writer.String())
}

func Test_output_error_with_format_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, nil, &writer)
	cli.Error("first %s", "second")
	require.Equal(t, "\033[31mfirst second\033[39m\n", writer.String())
}

func Test_output_line_with_empty_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Line("")
	require.Equal(t, "\033[39m\033[39m\n", writer.String())
}

func Test_output_line_with_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Line("val")
	require.Equal(t, "\033[39mval\033[39m\n", writer.String())
}

func Test_output_line_with_format_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Line("first %s", "second")
	require.Equal(t, "\033[39mfirst second\033[39m\n", writer.String())
}

func Test_output_comment_with_empty_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Comment("")
	require.Equal(t, "\u001B[30;1m\033[39m\n", writer.String())
}

func Test_output_comment_with_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Comment("val")
	require.Equal(t, "\u001B[30;1mval\033[39m\n", writer.String())
}

func Test_output_comment_with_format_string(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	cli.Comment("first %s", "second")
	require.Equal(t, "\u001B[30;1mfirst second\033[39m\n", writer.String())
}

func Test_output_table_with_title(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	table := cli.Table()
	table.AppendRow([]interface{}{"first", "second"})
	table.Render()
	require.Equal(t, "\n first second\n\n", TrimDoubleSpaces(writer.String()))
}

func Test_progress_bar(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	bar := cli.ProgressBar(4, "Sending emails")
	_ = bar.Add(4)

	require.Regexp(t, `100% |█████████████████████████████████████████████| (4/4, .* it/s)`, writer.String())
}

func Test_progress_with_description(t *testing.T) {
	writer, app := setUp()
	cli := facade.NewCli(app, &writer, nil)
	bar := cli.ProgressBar(100, "Sending emails")
	_ = bar.Add(100)

	require.Contains(t, writer.String(), "Sending emails 100%")
}

func Test_ask_valid(t *testing.T) {
	writer, app := setUp()
	reader := strings.NewReader("horse\n")
	rc := ioutil.NopCloser(reader)
	cli := facade.NewCliByReadersAndWriter(app, rc, &writer, nil)

	answer := cli.Ask("Animal:")

	require.Contains(t, writer.String(), "Animal:")
	require.Contains(t, writer.String(), "horse")
	require.Equal(t, "horse", answer)
}

func Test_ask_empty_answer(t *testing.T) {
	writer, app := setUp()
	reader := strings.NewReader("\n")
	rc := ioutil.NopCloser(reader)
	cli := facade.NewCliByReadersAndWriter(app, rc, &writer, nil)

	answer := cli.Ask("Animal:")

	require.Contains(t, writer.String(), "Animal:")
	require.Equal(t, "", answer)
}

func Test_secret(t *testing.T) {
	writer, app := setUp()
	reader := strings.NewReader("a67is@fAs!\n")
	rc := ioutil.NopCloser(reader)
	cli := facade.NewCliByReadersAndWriter(app, rc, &writer, nil)

	answer := cli.Secret("Animal:")

	require.Contains(t, writer.String(), "Animal:")
	require.NotContains(t, writer.String(), "a67is@fAs!")
	require.Equal(t, "a67is@fAs!", answer)
}

func Test_choice_valid_answer(t *testing.T) {
	reader := strings.NewReader(navigateDown + navigateDown + "\n")
	writer, app := setUp()
	rc := ioutil.NopCloser(reader)
	cli := facade.NewCliByReadersAndWriter(app, rc, &writer, nil)

	answer := cli.Choice("Animal:", "dog", "bear", "horse")

	require.Contains(t, writer.String(), "Use the arrow keys to navigate:")
	require.Contains(t, writer.String(), "Animal:")
	require.Contains(t, writer.String(), "horse")
	require.Equal(t, "horse", answer)
}

func Test_confirm_true(t *testing.T) {
	writer, app := setUp()
	reader := strings.NewReader("y\n")
	rc := ioutil.NopCloser(reader)
	cli := facade.NewCliByReadersAndWriter(app, rc, &writer, nil)

	con := cli.Confirm("Sure?", false)

	require.Contains(t, writer.String(), "Sure?")
	require.True(t, con)
}

func Test_confirm_false(t *testing.T) {
	writer, app := setUp()
	reader := strings.NewReader("n\n")
	rc := ioutil.NopCloser(reader)
	cli := facade.NewCliByReadersAndWriter(app, rc, &writer, nil)

	con := cli.Confirm("Sure?", false)

	require.Contains(t, writer.String(), "Sure?")
	require.False(t, con)
}

func Test_confirm_default_true(t *testing.T) {
	writer, app := setUp()
	reader := strings.NewReader("\n")
	rc := ioutil.NopCloser(reader)
	cli := facade.NewCliByReadersAndWriter(app, rc, &writer, nil)

	con := cli.Confirm("Sure?", true)

	require.Contains(t, writer.String(), "Sure?")
	require.True(t, con)
}
