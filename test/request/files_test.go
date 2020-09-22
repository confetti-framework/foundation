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
	request := requestByFiles("--xxx\nContent-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	_, err := request.FileE("book")
	assert.EqualError(t, err, "file not found by key: book")
}

func TestOneFileFoundWithContent(t *testing.T) {
	request := requestByFiles("--xxx\nContent-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	file, err := request.FileE("photo")
	assert.Nil(t, err)
	assert.Equal(t, "content_of_file", file.Content())
}

func TestFileNotFoundShouldPanic(t *testing.T) {
	request := requestByFiles("--xxx\nContent-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Panics(t, func() {
		request.File("book")
	})
}

func TestFileGetContent(t *testing.T) {
	request := requestByFiles("--xxx\nContent-Disposition: form-data; name=\"photo\"; filename=\"file.txt\"\n" +
		"Content-Type: text/plain\n\ncontent_of_file\n--xxx--")

	assert.Equal(t, "content_of_file", request.File("photo").Content())
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
