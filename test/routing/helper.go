package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/test/mock"
)

func emptyController() func(request inter.Request) inter.Response {
	return func(request inter.Request) inter.Response { return nil }
}

func newRequest(options http.Options) inter.Request {
	app := foundation.NewApp()
	app.Bind("outcome_html_encoders", mock.HtmlEncoders)
	app.Bind("response_decorators", []inter.ResponseDecorator{})
	options.App = app
	return http.NewRequest(options)
}
