package providers

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
	"strings"
)

type ConfigServiceProvider struct {
	Index []interface{}
}

func (c ConfigServiceProvider) Register(container inter.Container) inter.Container {
	container.Instance("config", c.GetMap())
	return container
}

func (c ConfigServiceProvider) GetMap() map[string]interface{} {
	result := map[string]interface{}{}

	for _, config := range c.Index {
		result[getNameSpace(config)] = config
	}

	return result
}

func getNameSpace(config interface{}) string {
	name := support.Name(config)
	parts := strings.Split(name, ".")
	if len(parts) < 2 {
		panic(errors.New("config need to be a named struct. " + name + " given."))
	}
	return parts[1]
}
