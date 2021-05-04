package http

import (
	"bytes"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/support"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
)

type Request struct {
	app       inter.App
	source    http.Request
	urlValues support.Map
	content   support.Value
}

type Options struct {
	App     inter.App
	Source  http.Request
	Method  string
	Host    string
	Url     string
	Header  http.Header
	Form    url.Values
	Content string
	Route   *mux.Route
	Body    io.ReadCloser
}

func NewRequest(options Options) inter.Request {
	var body io.Reader

	if options.Content != "" {
		body = bytes.NewBufferString(options.Content)
	}

	if options.Url == "" {
		options.Url = "/"
	}

	source := options.Source
	if source.Method == "" {
		// For testing purpose
		source = *httptest.NewRequest(options.Method, options.Url, body)
		if options.Form != nil {
			source.Header.Set("Content-Type", "multipart/form-data; boundary=xxx")
			source.Form = options.Form
		}

		if options.Host != "" {
			source.Host = options.Host
		}

		if options.Body != nil {
			source.Body = options.Body
		}

		if options.Header != nil {
			source.Header = options.Header
		}
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
	if r.app == nil {
		panic("app in request is nil")
	}
	return r.app
}

func (r *Request) SetApp(app inter.App) {
	r.app = app
}

func (r *Request) Make(abstract interface{}) interface{} {
	return r.App().Make(abstract)
}

func (r *Request) MakeE(abstract interface{}) (interface{}, error) {
	return r.App().MakeE(abstract)
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

func (r Request) Body() string {
	body, err := ioutil.ReadAll(r.source.Body)
	if err == io.EOF {
		return ""
	}

	return string(body)
}

func (r *Request) SetBody(body string) inter.Request {
	// Update source body
	r.source.Body = ioutil.NopCloser(strings.NewReader(body))

	// Invalidate Confetti body. Rebuild content when requested.
	r.content = support.NewValue(nil)

	return r
}

func (r *Request) Content(keyInput ...string) support.Value {
	result, err := r.ContentE(keyInput...)
	if err != nil {
		panic(err)
	}
	return result
}

func (r *Request) ContentE(keyInput ...string) (support.Value, error) {
	// Let keyInput be a optional parameter
	var key string
	var err error
	if len(keyInput) > 0 {
		key = keyInput[0]
	}

	r.content, err = r.generateContentFromBody()
	if err != nil {
		return support.Value{}, err
	}

	return r.content.GetE(key)
}

func (r Request) ContentOr(key string, defaultValue interface{}) support.Value {
	result, err := r.ContentE(key)
	if err != nil {
		return support.NewValue(defaultValue)
	}

	return result
}

func (r Request) Parameter(key string) support.Value {
	result, err := r.ParameterE(key)
	if err != nil {
		panic(err)
	}
	return result
}

func (r Request) ParameterE(key string) (support.Value, error) {
	value, err := r.parameters().GetE(key)
	if err != nil {
		err = errors.Wrap(err, "from route parameter")
	}
	return value, err
}

func (r Request) ParameterOr(key string, defaultValue interface{}) support.Value {
	value, err := r.ParameterE(key)
	if err != nil {
		return support.NewValue(defaultValue)
	}
	return value
}

func (r *Request) SetUrlValues(vars map[string]string) inter.Request {
	r.urlValues = support.NewMap(vars)
	return r
}

func (r Request) Query(key string) support.Value {
	result, err := r.QueryE(key)
	if err != nil {
		panic(err)
	}
	return result
}

func (r Request) QueryE(key string) (support.Value, error) {
	value, err := support.NewMap(r.Source().URL.Query()).GetE(key)
	if err != nil {
		err = errors.Wrap(err, "from URL query")
	}
	return value, err
}

func (r Request) QueryOr(key string, defaultValue interface{}) support.Value {
	result, err := support.NewMap(r.Source().URL.Query()).GetE(key)
	if err != nil {
		return support.NewValue(defaultValue)
	}

	return result
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

func (r *Request) File(key string) support.File {
	file, err := r.FileE(key)
	if err != nil {
		panic(err)
	}
	return file
}

func (r *Request) FileE(key string) (support.File, error) {
	var file support.File
	files, err := r.FilesE(key)

	if len(files) != 0 {
		file = files[0]
	} else {
		file = support.File{}
	}

	return file, err
}

func (r *Request) Files(key string) []support.File {
	files, err := r.FilesE(key)
	if err != nil {
		panic(err)
	}
	return files
}

func (r *Request) FilesE(key string) ([]support.File, error) {
	if r.source.MultipartForm == nil {
		err := r.source.ParseMultipartForm(defaultMaxMemory)
		if err != nil {
			return []support.File{}, err
		}
	}
	if r.source.MultipartForm != nil && r.source.MultipartForm.File != nil {
		allFiles := r.source.MultipartForm.File
		if fileHeaders := allFiles[key]; len(fileHeaders) > 0 {
			return r.getFilesByHeaders(fileHeaders)
		}
	}
	return []support.File{}, errors.New("file not found by key: " + key)
}

func (r Request) Route() inter.Route {
	return r.app.Make("route").(inter.Route)
}

func (r Request) parameters() support.Map {
	urlMap := r.urlValues
	queryMap := support.NewMap(r.Source().URL.Query())

	return support.NewMap().Merge(urlMap, queryMap)
}

func (r Request) generateContentFromBody() (support.Value, error) {
	if r.content.Filled() {
		return r.content, nil
	}

	rawDecoder, err := r.MakeE(inter.RequestBodyDecoder)
	if errors.Is(err, support.CanNotFoundValueError) {
		return support.Value{}, errors.WithStack(NoRequestBodyDecoderFoundError)
	}

	decoder := rawDecoder.(func(request inter.Request) support.Value)
	body := decoder(&r)

	return body, nil
}

func (r *Request) getFilesByHeaders(fileHeaders []*multipart.FileHeader) ([]support.File, error) {
	var result []support.File
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		result = append(result, support.NewFile(file, fileHeader))
	}
	return result, nil
}
