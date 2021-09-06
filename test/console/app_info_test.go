//go:build !race
// +build !race

package console

import (
	"bytes"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_app_info_show_golang_version(t *testing.T) {
	writer, app := setUp()
	var writerErr bytes.Buffer

	app.Bind("config.App.OsArgs", []interface{}{"/main", "app:info"})

	code := console.Kernel{
		App:       app,
		Writer:    &writer,
		WriterErr: &writerErr,
		Commands:  []inter.Command{console.AppInfo{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, writer.String(), `Golang Version`)
}
