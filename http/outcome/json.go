package outcome

import (
	"github.com/confetti-framework/contract/inter"
	"net/http"
)

type JsonResponse struct {
	*Response
}

func Json(content interface{}) inter.Response {
	return &JsonResponse{
		Response: NewResponse(Options{
			Content:  content,
			Encoders: "outcome_json_encoders",
			Headers:  http.Header{"Content-Type": {"application/json", "charset=UTF-8"}},
		}),
	}
}
