package outcome

import "github.com/lanvard/contract/inter"

type Response struct {
	content string
}

func NewResponse() Response {
	return Response{}
}

func Json(content string) inter.Response {
	return &Response{content: content}
}

func (r *Response) Content() string {
	return r.content
}

func (r *Response) SetContent(content string) inter.Response {
	r.content = content

	return r
}
