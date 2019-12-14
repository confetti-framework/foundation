package bootstrap

import (
	"fmt"
	"github.com/lanvard/foundation"
	"lanvard/config"
	"lanvard/config/entity"
	"path/filepath"
)

type BasePathProvider struct {
}

func (l BasePathProvider) Bootstrap(app foundation.Application) foundation.Application {
	_, kernelDir, _, _ := runtime.Caller(4)
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(kernelDir)))

	config.App.BasePath = entity.BasePath(rootDir)

	return app
}
