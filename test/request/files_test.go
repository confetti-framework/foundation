package request

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFileNoContentTypeInRequest(t *testing.T) {
	request := requestByFilesWithoutContentType()

	_, err := request.FileE("photo")
	assert.EqualError(t, err, "request Content-Type isn't multipart/form-data")
}

func TestFileNotFoundInRequest(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	_, err := request.FileE("book")
	assert.EqualError(t, err, "file not found by key: book")
}

func TestOneFileFoundWithContent(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	file, err := request.FileE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_file", file.Content())
}

func TestTwoFilesGivenReceiveOneFile(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	file, err := request.FileE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_first_file", file.Content())
}

func TestTwoFilesDifferentKeyGiven(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo1\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo2\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	file, err := request.FileE("photo2")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_second_file", file.Content())

	file, err = request.FileE("photo1")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_first_file", file.Content())
}

func TestFileNotFoundShouldPanic(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Panics(t, func() {
		request.File("book")
	})
}

func TestFileGetContent(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, "content_of_file", request.File("photo").Content())
}

func TestFilesNoContentTypeInRequest(t *testing.T) {
	request := requestByFilesWithoutContentType()

	_, err := request.FilesE("photo")
	assert.EqualError(t, err, "request Content-Type isn't multipart/form-data")
}

func TestFilesNotFoundInRequest(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	_, err := request.FilesE("book")
	assert.EqualError(t, err, "file not found by key: book")
}

func TestFilesOneFoundWithContent(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	files, err := request.FilesE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_file", files[0].Content())
}

func TestFilesMultipleFoundWithContent(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	files, err := request.FilesE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_first_file", files[0].Content())
	assert.Equal(t, "content_of_second_file", files[1].Content())
}

func TestGetFileHeaderFromSecondFile(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	files, err := request.FilesE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "file2.txt", files[1].Header().Filename)
}

func TestFilesNotFoundShouldPanic(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Panics(t, func() {
		request.Files("book")
	})
}

func TestFilesWithoutError(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	assert.Equal(t, "content_of_second_file", request.Files("photo")[1].Content())
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
