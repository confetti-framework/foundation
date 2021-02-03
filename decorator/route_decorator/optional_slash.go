package route_decorator

import (
	"github.com/confetti-framework/contract/inter"
	"strings"
)

const optionalSlash = "{allow_slash:\\/?}"

type UriSuffixSlash struct{}

func (o UriSuffixSlash) Decorate(route inter.Route) inter.Route {
	uri := route.Uri()

	if !strings.Contains(uri, "?") && !strings.HasSuffix(uri, "/") {
		route.SetUri(uri + optionalSlash)
	}

	return route
}

// Revert decorator
func (o UriSuffixSlash) Revert(route inter.Route) inter.Route {
	uri := route.Uri()

	uri = strings.ReplaceAll(uri, optionalSlash, "")

	return route.SetUri(uri)
}
