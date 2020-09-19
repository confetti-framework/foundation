package http

import (
	"bytes"
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
	body      support.Value
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
		body = bytes.NewBufferString(options.Body)
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
	} else {
		request.urlValues = support.Map{}
	}

	return &request
}

func (r Request) App() inter.App {
	return r.app
}

func (r *Request) SetApp(app inter.App) {
	r.app = app
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

func (r Request) Path() string {
	return r.source.URL.Path
}

func (r Request) Url() string {
	return r.source.URL.Scheme + r.source.Host + r.source.URL.Path
}

func (r Request) FullUrl() string {
	return r.source.URL.Scheme + r.source.Host + r.source.RequestURI
}

func (r Request) Content() string {
	body, err := ioutil.ReadAll(r.source.Body)
	if err == io.EOF {
		return ""
	}

	return string(body)
}

func (r *Request) SetContent(content string) inter.Request {
	// Update source body
	r.source.Body = ioutil.NopCloser(strings.NewReader(content))

	// Invalidate Lanvard body. Rebuild this body when requested.
	r.body = support.NewValue(nil)

	return r
}

func (r *Request) Body(keyInput ...string) support.Value {
	// Let key be a default parameter
	var key string
	if len(keyInput) > 0 {
		key = keyInput[0]
	}

	formMap := support.NewMapByUrlValues(r.source.Form)
	if !formMap.Empty() {
		return formMap.Get(key)
	}

	r.body = r.generateBodyFromRawContent()

	return r.body.Get(key)
}

func (r Request) BodyOr(key string, defaultValue interface{}) support.Value {
	value := r.Body(key)
	if value.Error() == nil {
		return value
	}

	return support.NewValue(defaultValue)
}

func (r Request) Parameter(key string) support.Value {
	return r.parameters().Get(key)
}

func (r Request) ParameterOr(key string, defaultValue interface{}) support.Value {
	value := r.Parameter(key)
	if value.Error() == nil {
		return value
	}

	return support.NewValue(defaultValue)
}

func (r *Request) SetUrlValues(vars map[string]string) inter.Request {
	r.urlValues = support.NewMapByString(vars)
	return r
}

func (r Request) Query(key string) support.Value {
	return support.NewMapByUrlValues(r.Source().URL.Query()).Get(key)
}

func (r Request) QueryOr(key string, defaultValue interface{}) support.Value {
	value := support.NewMapByUrlValues(r.Source().URL.Query()).Get(key)
	if value.Error() == nil {
		return value
	}

	return support.NewValue(defaultValue)
}

func (r Request) Header(key string) string {
	return r.source.Header.Get(key)
}

func (r Request) Headers() http.Header {
	return r.source.Header
}

func (r Request) Cookie(key string) string {
	result, err := r.CookieE(key)
	if err != nil {
		panic(err)
	}
	return result
}

func (r Request) CookieE(key string) (string, error) {
	var result string
	cookie, err := r.source.Cookie(key)
	if cookie != nil {
		result = cookie.Value
	}

	return result, err
}

func (r Request) Route() inter.Route {
	return r.app.Make("route").(inter.Route)
}

func (r Request) parameters() support.Map {
	urlMap := r.urlValues
	queryMap := support.NewMapByUrlValues(r.Source().URL.Query())

	return support.NewMap().Merge(urlMap, queryMap)
}

func (r Request) generateBodyFromRawContent() support.Value {
	if r.body.Filled() {
		return r.body
	}

	rawBody, err := ioutil.ReadAll(r.source.Body)
	if err != nil {
		return support.NewValueE(rawBody, err)
	}

	rawDecoder := r.Make(inter.RequestBodyDecoder)
	if rawDecoder == nil {
		return support.NewValueE(nil, errors.New("no request body decoder found"))
	}

	decoder := rawDecoder.(func(string) support.Value)
	body := decoder(string(rawBody))

	return body
}
