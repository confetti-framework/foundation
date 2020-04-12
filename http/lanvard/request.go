package lanvard

import (
	"github.com/lanvard/foundation"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	App      *foundation.Application
	Original http.Request
}

func NewRequest(app *foundation.Application, request http.Request) Request {
	return Request{App: app, Original: request}
}

func (r Request) Content() string {
	body, err := ioutil.ReadAll(r.Original.Body)
	if err == io.EOF {
		return ""
	}

	return string(body)
}

func (r Request) SetContent(content string) Request {
	r.Original.Body = ioutil.NopCloser(strings.NewReader(content))

	return r
}

func (r Request) GetMethod() string {
	return r.Original.Method
}
