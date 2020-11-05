package mock

import (
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
	"golang.org/x/text/language"
	"strings"
)

type errorViewMock struct {
	Message    string
	StackTrace string
	Status     int
	AppName    string
	Locale     string
}

func NewViewErrorMock(app inter.App, err error) inter.View {
	status, _ := errors.FindStatus(err)

	return &errorViewMock{
		Message:    str.UpperFirst(fmt.Sprintf("%v", err)),
		StackTrace: StackTrace(app, err),
		Status:     status,
		AppName:    AppName(app),
		Locale:     Locale(app),
	}
}

func (_ errorViewMock) Template() string {
	return TemplateByName("error_template.gohtml")
}

func AppName(app inter.App) string {
	name, err := app.MakeE("config.App.Name")
	if err != nil {
		return ""
	}
	return name.(string)
}

func Locale(app inter.App) string {
	lang, err := app.MakeE("config.App.Locale")
	if err != nil {
		return ""
	}
	return lang.(language.Tag).String()
}

func StackTrace(app inter.App, err error) string {
	if !app.Make("config.App.Debug").(bool) {
		return ""
	}
	stackTrace, _ := errors.FindStack(err)
	result := fmt.Sprintf("%+v", stackTrace)
	result = strings.TrimLeft(result, "\n")
	return result
}
