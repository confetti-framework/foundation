package http

import (
	"github.com/gorilla/mux"
	"github.com/lanvard/contract/inter"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

type Request struct {
	app       inter.App
	source    http.Request
	urlValues map[string]string
}

type Options struct {
	App     inter.App
	Source  http.Request
	Method  string
	Host    string
	Uri     string
	Headers http.Header
	Body    string
	Route   *mux.Route
}

func NewRequest(options Options) inter.Request {

	var body io.Reader

	if options.Body != "" {
		body = strings.NewReader(options.Body)
	}

	source := options.Source

	if options.Uri == "" {
		options.Uri = "/"
	}

	source = *httptest.NewRequest(options.Method, options.Uri, body)

	if options.Host != "" {
		source.Host = options.Host
	}

	if options.Headers != nil {
		source.Header = options.Headers
	}

	request := Request{source: source}

	if options.App != nil {
		request.app = options.App
	}

	// If a route has been identified (usually by a test), add route values to request.
	if options.Route != nil {
		var match mux.RouteMatch
		ok := options.Route.Match(&source, &match)
		if !ok {
			panic("Route don't match with url")
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

func (r Request) App() inter.App {
	return r.app
}

func (r *Request) SetApp(app inter.App) inter.Request {
	r.app = app

	return r
}

func (r *Request) Make(abstract interface{}) interface{} {
	return r.App().Make(abstract)
}

func (r Request) Source() http.Request {
	return r.source
}

func (r Request) Method() string {
	return r.source.Method
}

func (r Request) UrlValue(key string) inter.Value {
	return NewUrlByValues(r.urlValues).Get(key)
}

func (r *Request) SetUrlValues(vars map[string]string) inter.Request {
	r.urlValues = vars
	return r
}

func (r Request) QueryValue(key string) inter.Value {
	return NewUrlByMultiValues(r.Source().URL.Query()).Get(key)
}

func (r Request) QueryValues(key string) []inter.Value {
	return NewUrlByMultiValues(r.Source().URL.Query()).GetMulti(key)
}

func (r Request) Header(key string) string {
	return r.source.Header.Get(key)
}

func (r Request) Headers() http.Header {
	return r.source.Header
}

func (r Request) Route() inter.Route {
	return r.app.Make("route").(inter.Route)
}
