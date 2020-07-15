package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support/transformer"
)

type RequestJsonBody struct{}

// This ensures the request can be decoded from JSON.
func (j RequestJsonBody) Handle(request inter.Request, next inter.Next) inter.Response {
	//     public function isJson()
	//    {
	//        return Str::contains($this->header('CONTENT_TYPE'), ['/json', '+json']);
	//    }
	request.App().Singleton(inter.RequestBodyDecoder, transformer.JsonToValue)
	return next(request)
}
