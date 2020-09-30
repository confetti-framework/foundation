package providers

import (
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
		name := strings.Split(support.Name(config), ".")[1]
		result[name] = config
	}

	return result
}
