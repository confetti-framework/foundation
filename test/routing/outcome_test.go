package routing

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/confetti-framework/support/caller"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_json_with_default_header(t *testing.T) {
	response := outcome.Json("")
	require.Equal(t, http.Header{"Content-Type": {"application/json", "charset=UTF-8"}}, response.GetHeaders())
}

func Test_html_with_default_header(t *testing.T) {
	response := outcome.Html("")
	require.Equal(t, http.Header{"Content-Type": {"text/html", "charset=UTF-8"}}, response.GetHeaders())
}

func Test_content_without_default_header(t *testing.T) {
	response := outcome.Content("")
	require.Equal(t, http.Header{}, response.GetHeaders())
}

func Test_outcome_with_fluent_headers(t *testing.T) {
	response := outcome.Html("").Headers(http.Header{"Content-Type": {"application/pdf", "charset=UTF-8"}})
	require.Equal(t, http.Header{"Content-Type": {"application/pdf", "charset=UTF-8"}}, response.GetHeaders())
}

func Test_outcome_with_fluent_add_one_header(t *testing.T) {
	response := outcome.Content("").Header("Content-Type", "application/pdf")
	require.Equal(t, http.Header{"Content-Type": {"application/pdf"}}, response.GetHeaders())
}

func Test_outcome_with_fluent_add_multiple_headers(t *testing.T) {
	response := outcome.Content("").
		Header("Content-Type", "application/pdf").
		Header("X-Header-One", "Header Value")
	expected := http.Header{"Content-Type": {"application/pdf"}, "X-Header-One": {"Header Value"}}
	require.Equal(t, expected, response.GetHeaders())
}

func Test_outcome_with_fluent_add_multiple_headers_same_key(t *testing.T) {
	response := outcome.Content("").
		Header("Content-Type", "application/pdf", "charset=UTF-8")
	expected := http.Header{"Content-Type": {"application/pdf", "charset=UTF-8"}}
	require.Equal(t, expected, response.GetHeaders())
}

func Test_outcome_redirect_permanent(t *testing.T) {
	response := outcome.RedirectPermanent("")
	require.Equal(t, 301, response.GetStatus())
}

func Test_outcome_redirect_temporary(t *testing.T) {
	response := outcome.RedirectTemporary("")
	require.Equal(t, 302, response.GetStatus())
}

func Test_outcome_download_content_by_non_existing_file(t *testing.T) {
	result, err := outcome.DownloadE("non_existing_file.txt")
	require.True(t, errors.Is(err, outcome.FileNotFoundError))
	require.Equal(t, "can't download file non_existing_file.txt: file not found", err.Error())
	require.Nil(t, result)
}

func Test_no_file_found_without_returning_error(t *testing.T) {
	require.PanicsWithError(t, "can't download file non_existing_file.txt: file not found", func() {
		outcome.Download("non_existing_file.txt")
	})
}

func Test_existing_file_without_returning_error(t *testing.T) {
	dir := caller.CurrentDir()
	result := outcome.Download(dir + "/mock_file.md")
	require.Equal(t, "# Mock File", result.GetContent())
	require.Equal(t, `attachment; filename="mock_file.md"`, result.GetHeader("Content-Disposition"))
}

func Test_outcome_download_content_by_directory(t *testing.T) {
	dir := caller.CurrentDir()
	result, err := outcome.DownloadE(dir)
	require.True(t, errors.Is(err, outcome.FileNotFoundError))
	require.Contains(t, err.Error(), "test/routing: file not found")
	require.Nil(t, result)
}

func Test_outcome_download_content_by_valid_file(t *testing.T) {
	dir := caller.CurrentDir()
	result, err := outcome.DownloadE(dir + "/mock_file.md")
	require.Nil(t, err)
	require.Equal(t, "# Mock File", result.GetContent())
	require.Equal(t, `attachment; filename="mock_file.md"`, result.GetHeader("Content-Disposition"))
}

func Test_outcome_with_filename(t *testing.T) {
	result := outcome.Content("Só para testar").Filename("filename.jpg")
	require.Equal(t, `attachment; filename="filename.jpg"`, result.GetHeader("Content-Disposition"))
}

func Test_outcome_show_in_browser(t *testing.T) {
	result := outcome.Content("Só para testar").ShowInBrowser()
	require.Equal(t, `inline`, result.GetHeader("Content-Disposition"))
}

func Test_outcome_download_show_in_browser(t *testing.T) {
	dir := caller.CurrentDir()
	result, _ := outcome.DownloadE(dir + "/mock_file.md")
	require.Equal(t, `inline`, result.ShowInBrowser().GetHeader("Content-Disposition"))
}

func Test_outcome_download_with_header_from_file(t *testing.T) {
	dir := caller.CurrentDir()
	result, _ := outcome.DownloadE(dir + "/mock_file.md")
	require.Equal(t, `text/markdown`, result.GetHeader("Content-Type"))
}
