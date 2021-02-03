package outcome

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/encoder"
	"net/http"
	"strings"
)

type Response struct {
	app          inter.App
	content      interface{}
	body         string
	headers      http.Header
	cookies      []http.Cookie
	status       int
	encoderAlias string
}

type Options struct {
	App      inter.App
	Content  interface{}
	Headers  http.Header
	Status   int
	Encoders string
}

func NewResponse(options Options) *Response {
	// handle default
	options = applyDefaultOptions(options)

	response := &Response{}
	response.app = options.App
	response.status = options.Status
	response.headers = options.Headers
	response.encoderAlias = options.Encoders
	response.content = options.Content

	return response
}

func (r Response) App() inter.App {
	return r.app
}

func (r *Response) SetApp(app inter.App) {
	r.app = app
}

// Receive the raw content. Only available before it has been converted
// to a body (by "response_decorators")
func (r *Response) GetContent() interface{} {
	if r.body != "" {
		panic("can't get content. GetContent is already converted to body")
	}
	return r.content
}

// Put the raw content. "response_decorators" are responsible
// for converting the content into a body
func (r *Response) Content(content interface{}) {
	if r.body != "" {
		panic("can't set content. GetContent is already converted to body")
	}
	r.content = content
}

func (r Response) GetBody() string {
	result, err := r.GetBodyE()
	if err != nil {
		panic(err)
	}

	return result
}

func (r Response) GetBodyE() (string, error) {
	if r.body != "" {
		return r.body, nil
	}
	if r.encoderAlias == "" {
		return "", errors.New("can't transform response object to string. No response encoder alias defined in outcome.Response")
	}
	instance, err := r.app.MakeE(r.encoderAlias)
	if err != nil {
		return "", err
	}
	encoders, ok := instance.([]inter.Encoder)
	if !ok {
		return "", errors.New("no valid response encoder found")
	}

	body, err := encoder.EncodeThrough(r.app, r.content, encoders)
	r.body = body

	return r.body, err
}

func (r *Response) Body(body string) inter.Response {
	r.body = body

	return r
}

func (r Response) GetStatus() int {
	return r.status
}

func (r *Response) Status(status int) inter.Response {
	r.status = status

	return r
}

func (r Response) GetHeader(key string) string {
	return strings.Join(r.headers.Values(key), "; ")
}

func (r Response) GetHeaders() http.Header {
	return r.headers
}

func (r *Response) Header(key string, values ...string) inter.Response {
	r.headers[key] = values
	return r
}

func (r *Response) Headers(headers http.Header) inter.Response {
	r.headers = headers
	return r
}

func (r *Response) Filename(filename string) inter.Response {
	r.Header("Content-Disposition", "attachment", "filename=\""+filename+"\"")
	return r
}

func (r *Response) ShowInBrowser() inter.Response {
	r.Header("Content-Disposition", "inline")
	return r
}

func (r *Response) Cookie(cookies ...http.Cookie) inter.Response {
	r.cookies = append(r.cookies, cookies...)
	return r
}

func (r Response) GetCookies() []http.Cookie {
	return r.cookies
}

func applyDefaultOptions(options Options) Options {
	if options.Status == 0 {
		options.Status = http.StatusOK
	}
	if options.Headers == nil {
		options.Headers = make(http.Header)
	}
	if options.Encoders == "" {
		options.Encoders = "default_response_outcome"
	}

	return options
}
