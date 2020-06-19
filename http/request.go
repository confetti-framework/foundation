package http

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/method"
	"github.com/lanvard/support"
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
	urlValues support.Map
}

type Options struct {
	App     inter.App
	Source  http.Request
	Method  string
	Host    string
	Uri     string
	Headers http.Header
	Form    url.Values
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

	if options.Form != nil {
		source.Header.Set("Content-Type", "multipart/form-data; boundary=xxx")

		source.Form = options.Form
	}

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
	if r.source.Method == "" {
		return method.Get
	}

	return r.source.Method
}

func (r Request) IsMethod(method string) bool {
	return r.Method() == strings.ToUpper(method)
}

func (r Request) Path() string {
	return r.source.URL.Path
}

func (r Request) Url() string {
	return r.source.URL.Scheme + r.source.Host + r.source.URL.Path
}

func (r Request) FullUrl() string {
	return r.source.URL.Scheme + r.source.Host + r.source.RequestURI
}

func (r Request) All() inter.Bag {
	urlBag := support.NewBagByMap(r.urlValues)
	queryBag := support.NewBag(r.Source().URL.Query())
	formBag := support.NewBag(r.source.Form)
	return urlBag.Merge(queryBag, formBag)
}

func (r Request) Value(key string) inter.Value {
	values := r.Values(key)
	if values.Len() == 0 {
		return support.NewValueE("", errors.New(key + " not found"))
	}

	return values.First()
}

func (r Request) ValueOr(key string, value interface{}) inter.Value {
	values := r.Values(key)
	if values.Len() == 0 {
		return support.NewValue(value)
	}

	return values.First()
}

func (r Request) Values(key string) inter.Collection {
	return r.All().GetMany(key)
}

func (r *Request) SetUrlValues(vars map[string]string) inter.Request {
	r.urlValues = support.NewMapByString(vars)
	return r
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
