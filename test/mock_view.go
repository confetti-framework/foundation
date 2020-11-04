package test

import (
	"path/filepath"
	"runtime"
)

type ViewMock struct{}

func (v ViewMock) Template() string {
	_, file, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(file)
	return currentDir + "/mock_template.gohtml"
}
