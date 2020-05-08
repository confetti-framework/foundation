package http

import (
	"github.com/gorilla/mux"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

type Request struct {
	app       inter.App
	source    http.Request
	urlValues map[string]string
}

type Options struct {
	App    inter.App
	Source http.Request
	Method string
	Url    string
	Header http.Header
	Host   string
	Route  *mux.Route
}

func NewRequest(options Options) inter.Request {

	if "" != options.Method {
		options.Source.Method = options.Method
	}

	if "" != options.Host {
		options.Source.Host = options.Host
	}

	if "" != options.Url {
		options.Source.URL = &url.URL{Path: options.Url}
	}

	if nil != options.Header {
		options.Source.Header = options.Header
	}

	// @todo make body fillable for tests
	var body io.Reader

	sourceRequest := httptest.NewRequest(options.Source.Method, options.Source.URL.Path, body)
	options.Source = *sourceRequest

	request := Request{source: options.Source}

	if nil != options.App {
		request.app = options.App
	}

	// add route values to request
	if nil != options.Route {
		var match mux.RouteMatch
		ok := options.Route.Match(&options.Source, &match)
		if !ok {
			panic("test route don't match with url")
		}

		request.SetUrlValues(match.Vars)
	}

	return &request
}

func (r Request) Content() string {
	body, err := ioutil.ReadAll(r.source.Body)
	if err == io.EOF {
		return ""
	}

	return string(body)
}

func (r *Request) SetContent(content string) inter.Request {
	r.source.Body = ioutil.NopCloser(strings.NewReader(content))

	return r
}

func (r Request) Method() string {
	return r.source.Method
}

func (r Request) App() inter.App {
	return r.app
}

func (r *Request) SetApp(app inter.App) inter.Request {
	r.app = app

	return r
}

func (r Request) Source() http.Request {
	return r.source
}

func (r Request) UrlValues() inter.UrlValues {
	return routing.NewUrlByValues(r.urlValues)
}

func (r *Request) SetUrlValues(vars map[string]string) inter.Request {
	r.urlValues = vars
	return r
}

func (r Request) QueryValues() inter.UrlValues {
	return routing.NewUrlByMultiValues(r.Source().URL.Query())
}
