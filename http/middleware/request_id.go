package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/google/uuid"
)

type RequestID struct{}

const requestHeaderIDName = "x-request-id"

func (_ RequestID) Handle(request inter.Request, next inter.Next) inter.Response {
	var reqID string

	if reqID = request.Headers().Get(requestHeaderIDName); reqID == "" {
		reqID = uuid.New().String()
		request.Headers().Set(requestHeaderIDName, reqID)
	}

	resp := next(request)
	resp.GetHeaders().Set(requestHeaderIDName, reqID)

	return resp
}
