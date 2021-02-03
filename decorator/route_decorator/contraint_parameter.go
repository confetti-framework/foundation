package route_decorator

import (
	"github.com/confetti-framework/contract/inter"
	"strings"
)

type ConstrainParameters struct{}

func (c ConstrainParameters) Decorate(route inter.Route) inter.Route {
	uri := route.Uri()
	if !strings.Contains(uri, "}") {
		return route
	}

	for parameter, constrainRegex := range route.Constraint() {
		oldMatch := "{" + parameter + "}"
		newMatch := "{" + parameter + ":" + constrainRegex + "}"
		uri = strings.Replace(uri, oldMatch, newMatch, 10)
	}

	return route.SetUri(uri)
}
