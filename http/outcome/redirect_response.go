package outcome

import (
	"bytes"
	"github.com/confetti-framework/contract/inter"
	"html/template"
	net "net/http"
)

type redirectResponse struct {
	*Response
	targetUrl string
}

var t *template.Template

func newRedirectResponse() *redirectResponse {
	response := NewResponse(Options{Encoders: "outcome_html_encoders"})
	return &redirectResponse{Response: response}
}

func Redirect(destination string, status int) inter.Response {
	redirectResponse := newRedirectResponse()
	redirectResponse.Status(status)
	redirectResponse.setTargetUrl(destination)

	return redirectResponse
}

// Create a permanent redirect from one URI to another.
func RedirectPermanent(destination string) inter.Response {
	return Redirect(destination, net.StatusMovedPermanently)
}

// Create a temporary redirect from one URI to another.
func RedirectTemporary(destination string) inter.Response {
	return Redirect(destination, net.StatusFound)
}

// Redirect to route by route name
func RedirectToRoute(app inter.App, route string, parameters ...Parameters) inter.Response {
	destination := UrlByName(app, route, parameters...)

	redirectResponse := newRedirectResponse()
	redirectResponse.Status(net.StatusFound)
	redirectResponse.setTargetUrl(destination)

	return redirectResponse
}

func (r *redirectResponse) setTargetUrl(url string) {
	if "" == url {
		url = "/"
	}

	r.targetUrl = url

	const tmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8" />
        <meta http-equiv="refresh" content="0;url={{.}}" />

        <title>Redirecting to {{.}}</title>
    </head>
    <body>
        Redirecting to <a href="{{.}}">{{.}}</a>.
    </body>
</html>`

	buf := new(bytes.Buffer)
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(buf, url)
	if err != nil {
		panic(err)
	}

	r.Content(buf.String())

	r.headers.Set("Location", url)
}
