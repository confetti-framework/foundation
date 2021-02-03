package outcome

import (
	"github.com/confetti-framework/contract/inter"
	"net/http"
)

type HtmlResponse struct {
	*Response
}

func Html(content interface{}) inter.Response {
	return &HtmlResponse{
		Response: NewResponse(Options{
			Content:  content,
			Encoders: "outcome_html_encoders",
			Headers:  http.Header{"Content-Type": {"text/html", "charset=UTF-8"}},
		}),
	}
}
