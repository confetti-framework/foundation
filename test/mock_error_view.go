package test

import (
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support/str"
	"golang.org/x/text/language"
	"path/filepath"
	"runtime"
	"strings"
)

type viewErrorMock struct {
	Message    string
	StackTrace string
	Status     int
	AppName    string
	Locale     string
}

func NewViewErrorMock(app inter.App, err error) inter.View {
	status, _ := errors.FindStatus(err)

	return &viewErrorMock{
		Message:    str.UpperFirst(fmt.Sprintf("%v", err)),
		StackTrace: StackTrace(app, err),
		Status:     status,
		AppName:    AppName(app),
		Locale:     Locale(app),
	}
}

func (v viewErrorMock) Template() string {
	_, file, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(file)
	return currentDir + "/mock_error_template.gohtml"
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
