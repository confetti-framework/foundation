package outcome

import (
	"encoding/json"
	"github.com/lanvard/contract/inter"
)

type Response struct {
	content interface{}
}

func NewResponse() Response {
	return Response{}
}

func Json(content interface{}) inter.Response {
	return Response{content: content}
}

func Http(content interface{}) inter.Response {
	return Response{content: content}
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

	// @todo handle unknown type
	panic("Don't know type")
}

func (r Response) SetContent(content string) inter.Response {
	r.content = content

	return r
}
