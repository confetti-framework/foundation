package http_helper

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/gorilla/mux"
)

func MuxFromRoute(route inter.Route) *mux.Route {
	muxRoute := new(mux.Route)
	for _, prefix := range route.RouteOptions().Prefixes() {
		muxRoute.PathPrefix(prefix)
	}

	muxRoute.Path(route.Uri())
	if route.Domain() != "" {
		muxRoute.Host(route.Domain())
	}

	return muxRoute
}
