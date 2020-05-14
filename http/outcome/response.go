package outcome

import (
	"encoding/json"
	"github.com/lanvard/contract/inter"
	"net/http"
)

type Response struct {
	content interface{}
	headers http.Header
	status  int
}

func NewResponse() *Response {
	headers := make(http.Header)
	return &Response{status: http.StatusOK, headers: headers}
}

func Json(content interface{}) inter.Response {
	return &Response{status: http.StatusOK, content: content}
}

func Http(content interface{}) inter.Response {
	return &Response{status: http.StatusOK, content: content}
}

func (r Response) Content() string {

	stringer, ok := r.content.(string)
	if ok {
		return stringer
	}

	mapper, ok := r.content.(inter.ToMap)
	if ok {
		jsonString, err := json.Marshal(mapper.ToMap())
		if nil != err {
			panic(err.Error())
		}

		return string(jsonString)
	}

	return ""
}

func (r *Response) SetContent(content string) inter.Response {
	r.content = content

	return r
}

func (r Response) Status() int {
	return r.status
}

func (r *Response) SetStatus(status int) inter.Response {
	r.status = status

	return r
}

func (r Response) Header(key string) string {
	return r.headers.Get(key)
}

func (r Response) Headers() http.Header {
	return r.headers
}