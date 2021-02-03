package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
)

func redirectController(request inter.Request) inter.Response {
	rawRoute := request.App().Make("route")
	if rawRoute == nil {
		panic("no route found in request")
	}
	route := rawRoute.(*Route)
	options := route.routeOptions
	return outcome.Redirect(options.destination, options.status)
}
