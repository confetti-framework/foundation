package mock

import (
	_ "embed"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support/str"
	"golang.org/x/text/language"
	"strings"
)

//go:embed error_template.gohtml
var errorTemplate string

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
	return errorTemplate
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
