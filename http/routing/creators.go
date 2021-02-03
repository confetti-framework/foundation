package routing

import "github.com/confetti-framework/contract/inter"

func createRoutes(methods []string, uri string, controller inter.Controller) *RouteCollection {
	routes := NewRouteCollection()
	for _, method := range methods {
		route := NewRoute(uri, method, controller)
		routes.Push(route)
	}

	return routes
}

func createRoute(method string, uri string, controller inter.Controller) *RouteCollection {
	methods := []string{method}
	return createRoutes(methods, uri, controller)
}
