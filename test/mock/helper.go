package mock

import (
	"path/filepath"
	"runtime"
)

func TemplateByName(fileName string) string {
	_, file, _, _ := runtime.Caller(0)
	if fileName == "" {
		return filepath.Dir(file)
	}
	return filepath.Join(filepath.Dir(file), fileName)
}
