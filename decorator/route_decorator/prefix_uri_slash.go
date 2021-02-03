package route_decorator

import (
	"github.com/confetti-framework/contract/inter"
	"strings"
)

type UriPrefixSlash struct{}

func (o UriPrefixSlash) Decorate(route inter.Route) inter.Route {
	uri := route.Uri()

	if !strings.HasPrefix(uri, "/") {
		route.SetUri("/" + uri)
	}

	return route
}
