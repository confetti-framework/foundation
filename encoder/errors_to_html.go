package encoder

import (
	"bytes"
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
	"golang.org/x/text/language"
	"html/template"
	"io/ioutil"
	"strings"
)

type ErrorToHtml struct {
	TemplateFile string
}

func (e ErrorToHtml) IsAble(object interface{}) bool {
	_, ok := object.(error)
	return ok
}

func (e ErrorToHtml) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	err, ok := object.(error)
	if !ok {
		return "", errors.New("can't convert object to html in error format")
	}

	errorMessage := str.UpperFirst(fmt.Sprintf("%v", err))

	if e.TemplateFile != "" {
		status, _ := errors.FindStatus(err)
		view := ErrorView{
			Message:      errorMessage,
			StackTrace:   stackTraceFromError(app, err),
			Status:       status,
			AppName:      e.appName(app),
			Locale:       e.locale(app),
			templatePath: e.TemplateFile,
		}

		return contentByView(view)
	}
	return errorMessage, nil
}

func stackTraceFromError(app inter.App, err error) string {
	if !app.Make("config.App.Debug").(bool) {
		return ""
	}
	stackTrace, _ := errors.FindStack(err)
	result := fmt.Sprintf("%+v", stackTrace)
	result = strings.TrimLeft(result, "\n")
	return result
}

func (e ErrorToHtml) appName(app inter.App) string {
	name, err := app.MakeE("config.App.Name")
	if err != nil {
		return ""
	}
	return name.(string)
}

func (e ErrorToHtml) locale(app inter.App) string {
	lang, err := app.MakeE("config.App.Locale")
	if err != nil {
		return ""
	}
	return lang.(language.Tag).String()
}

func contentByView(view interface{ Template() string }) (string, error) {
	content, err := ioutil.ReadFile(view.Template())
	buf := bytes.NewBufferString("")
	if err != nil {
		return "", err
	}
	t, err := template.New("error").Parse(string(content))
	if err != nil {
		return "", err
	}
	err = t.Execute(buf, view)

	return buf.String(), nil
}

type ErrorView struct {
	Message      string
	StackTrace   string
	Status       int
	AppName      string
	Locale       string
	templatePath string
}

func (v ErrorView) Template() string {
	return v.templatePath
}
