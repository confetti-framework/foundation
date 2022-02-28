package outcome

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support"
	"io/ioutil"
	"os"
)

type ContentResponse struct {
	*Response
}

func Content(content interface{}) inter.Response {
	return &ContentResponse{
		Response: NewResponse(Options{
			Content:  content,
			Encoders: "outcome_content_encoders",
		}),
	}
}

func Download(filename string) inter.Response {
	response, err := DownloadE(filename)
	if err != nil {
		panic(err)
	}
	return response
}

func DownloadE(filename string) (inter.Response, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, errors.Wrap(FileNotFoundError, "can't download file %s", filename)
	}
	if info.IsDir() {
		return nil, errors.Wrap(FileNotFoundError, "can't download a directory %s", filename)
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	response := Content(string(content)).Filename(info.Name())
	mime, ok := support.MimeByExtension(info.Name())
	if ok {
		response.Header("Content-Type", mime)
	}
	return response, nil
}
