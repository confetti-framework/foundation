package outcome

import (
	"bytes"
	"html/template"
)

type RedirectResponse struct {
	*Response
	targetUrl string
}

var t *template.Template

func (r *RedirectResponse) SetTargetUrl(url string) {
	if "" == url {
		panic("Cannot redirect to an empty URL.")
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

	r.SetContent(buf.String())

	r.headers.Set("Location", url)
}

func newRedirectResponse() *RedirectResponse {
	response := NewResponse()
	return &RedirectResponse{Response: response}
}

func Redirect(destination string, status int) *RedirectResponse {
	redirectResponse := newRedirectResponse()
	redirectResponse.status = status
	redirectResponse.SetTargetUrl(destination)

	return redirectResponse
}
