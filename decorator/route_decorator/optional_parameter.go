package route_decorator

import (
	"github.com/confetti-framework/contract/inter"
	"regexp"
	"strings"
)

type OptionalParameter struct{}

func (o OptionalParameter) Decorate(route inter.Route) inter.Route {
	uri := route.Uri()
	domain := route.Domain()

	if !strings.Contains(uri, "?}") && !strings.Contains(domain, "?}") {
		return route
	}

	route.SetUri(o.implementOptionalParameter(uri))
	route.SetDomain(o.implementOptionalParameter(domain))

	return route
}

func (o OptionalParameter) implementOptionalParameter(uri string) string {
	regex := regexp.MustCompile(`{(?P<name>\w*)\?}`)
	submatch := regex.FindAllStringSubmatch(uri, 10)

	for _, parts := range submatch {
		newMatch := "{" + parts[1] + ":.*}"
		uri = strings.Replace(uri, parts[0], newMatch, 1)
	}

	return uri
}
