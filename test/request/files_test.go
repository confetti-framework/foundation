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

func TestOneFileFoundWithBody(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	file, err := request.FileE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_file", file.Body())
}

func TestTwoFilesGivenReceiveOneFile(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	file, err := request.FileE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_first_file", file.Body())
}

func TestTwoFilesDifferentKeyGiven(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo1\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo2\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	file, err := request.FileE("photo2")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_second_file", file.Body())

	file, err = request.FileE("photo1")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_first_file", file.Body())
}

func TestFileNotFoundShouldPanic(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Panics(t, func() {
		request.File("book")
	})
}

func TestFileGetBody(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, "content_of_file", request.File("photo").Body())
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

func TestFilesOneFoundWithBody(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	files, err := request.FilesE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_file", files[0].Body())
}

func TestFilesMultipleFoundWithBody(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_first_file\n--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file2.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_second_file\n--xxx--")

	files, err := request.FilesE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_first_file", files[0].Body())
	assert.Equal(t, "content_of_second_file", files[1].Body())
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

	assert.Equal(t, "content_of_second_file", request.Files("photo")[1].Body())
}

func TestFileNameNotPresent(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Panics(t, func() {
		request.File("book").Name()
	})
}

func TestFileNamePresent(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, "file1.txt", request.File("photo").Name())
}

func TestFileExtensionNotFound(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: invalid\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, "", request.File("photo").Extension())
}

func TestFileExtensionTxt(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, ".txt", request.File("photo").Extension())
}

func TestFileExtensionWithCharset(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: text/plain; charset=utf-32le\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, ".txt", request.File("photo").Extension())
}

func TestFileExtensionJpeg(t *testing.T) {
	request := requestByFiles("--xxx\n" +
		"Content-Disposition: form-data; name=\"photo\"; filename=\"file1.txt\"\n" +
		"Content-Type: image/jpeg\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, ".jpg", request.File("photo").Extension())
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
