package http

import (
	"github.com/lanvard/contract/inter"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	App    inter.App
	Source http.Request
}

type Options struct {
	App    inter.App
	Source http.Request
	Method string
	Url    string
	Header http.Header
	Host   string
}

func NewRequest(options Options) inter.Request {

	if options.Method != "" {
		options.Source.Method = options.Method
	}

	if options.Host != "" {
		options.Source.Host = options.Host
	}

	if options.Url != "" {
		options.Source.URL = &url.URL{Path: options.Url}
	}

	if options.Header != nil {
		options.Source.Header = options.Header
	}

	request := Request{}
	request.Source = options.Source
	request.App = options.App

	return request
}

func (r Request) Content() string {
	body, err := ioutil.ReadAll(r.Source.Body)
	if err == io.EOF {
		return ""
	}

	return string(body)
}

func (r Request) SetContent(content string) inter.Request {
	r.Source.Body = ioutil.NopCloser(strings.NewReader(content))

	return r
}

func (r Request) GetMethod() string {
	return r.Source.Method
}

func (r Request) SetApp(app inter.App) inter.Request {
	r.App = app

	return r
}

func (r Request) GetSource() http.Request {
	return r.Source
}