package request

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_file_no_content_type_in_request(t *testing.T) {
	request := requestByFilesWithoutContentType()

	_, err := request.FileE("photo")
	require.EqualError(t, err, "request Content-Type isn't multipart/form-data")
}

func Test_file_not_found_in_request(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	_, err := request.FileE("book")
	require.EqualError(t, err, "file not found by key: book")
}

func Test_one_file_found_with_body(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	file, err := request.FileE("photo")
	require.NoError(t, err)
	require.Equal(t, "content_of_file", file.Body())
}

func Test_two_files_given_receive_one_file(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	file, err := request.FileE("photo")
	require.NoError(t, err)
	require.Equal(t, "content_of_first_file", file.Body())
}

func Test_two_files_different_key_given(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo1\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo2\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	file, err := request.FileE("photo2")
	require.NoError(t, err)
	require.Equal(t, "content_of_second_file", file.Body())

	file, err = request.FileE("photo1")
	require.NoError(t, err)
	require.Equal(t, "content_of_first_file", file.Body())
}

func Test_file_not_found_should_panic(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	require.Panics(t, func() {
		request.File("book")
	})
}

func Test_file_get_body(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	require.Equal(t, "content_of_file", request.File("photo").Body())
}

func Test_files_no_content_type_in_request(t *testing.T) {
	request := requestByFilesWithoutContentType()

	_, err := request.FilesE("photo")
	require.EqualError(t, err, "request Content-Type isn't multipart/form-data")
}

func Test_files_not_found_in_request(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	_, err := request.FilesE("book")
	require.EqualError(t, err, "file not found by key: book")
}

func Test_files_one_found_with_body(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	files, err := request.FilesE("photo")
	require.NoError(t, err)
	require.Equal(t, "content_of_file", files[0].Body())
}

func Test_files_multiple_found_with_body(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	files, err := request.FilesE("photo")
	require.NoError(t, err)
	require.Equal(t, "content_of_first_file", files[0].Body())
	require.Equal(t, "content_of_second_file", files[1].Body())
}

func Test_get_file_header_from_second_file(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	files, err := request.FilesE("photo")
	require.NoError(t, err)
	require.Equal(t, "file2.txt", files[1].Header().Filename)
}

func Test_files_not_found_should_panic(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	require.Panics(t, func() {
		request.Files("book")
	})
}

func Test_files_without_error(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	require.Equal(t, "content_of_second_file", request.Files("photo")[1].Body())
}

func Test_file_name_not_present(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	require.Panics(t, func() {
		request.File("book").Name()
	})
}

func Test_file_name_present(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	require.Equal(t, "file1.txt", request.File("photo").Name())
}

func Test_file_extension_not_found(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: invalid\n\ncontent_of_file\n--xxx--")

	require.Equal(t, "", request.File("photo").Extension())
}

func Test_file_extension_txt(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	require.Equal(t, ".txt", request.File("photo").Extension())
}

func Test_file_extension_with_charset(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain; charset=utf-32le\n\ncontent_of_file\n--xxx--")

	require.Equal(t, ".txt", request.File("photo").Extension())
}

func Test_file_extension_jpeg(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: image/jpeg\n\ncontent_of_file\n--xxx--")

	require.Equal(t, ".jpg", request.File("photo").Extension())
}

func requestByFiles(content string) inter.Request {
	body := ioutil.NopCloser(strings.NewReader(content))
	options := http.Options{
		Body:   body,
		Header: map[string][]string{"Content-Type": {"multipart/form-data; boundary=xxx"}},
	}
	return http.NewRequest(options)
}

func requestByFilesWithoutContentType() inter.Request {
	var options http.Options
	return http.NewRequest(options)
}
