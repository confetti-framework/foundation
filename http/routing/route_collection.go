package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/decorator/route_decorator"
	"github.com/confetti-framework/foundation/http/http_helper"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/gorilla/mux"
)

type RouteCollection struct {
	// An array of the routes keyed by method.
	routesMapRoutes inter.MapMethodRoutes
	routes          []inter.Route
	decorators      []inter.RouteDecorator
	middlewares     []inter.HttpMiddleware
}

func NewRouteCollection(routeCollections ...inter.RouteCollection) *RouteCollection {
	collection := RouteCollection{}
	collection.Merge(flatten(routeCollections))
	collection.decorators = []inter.RouteDecorator{
		route_decorator.ConstrainParameters{},
		route_decorator.OptionalParameter{},
		route_decorator.UriPrefixSlash{},
		route_decorator.UriSuffixSlash{},
	}

	return &collection
}

func Group(routeCollections ...inter.RouteCollection) *RouteCollection {
	return NewRouteCollection(routeCollections...)
}

func DecorateRoutes(routes *RouteCollection) {
	for _, route := range routes.All() {
		route_decorator.Decorate(route, routes.decorators)
	}
	routes.decorators = []inter.RouteDecorator{}
}

func (c *RouteCollection) Push(route inter.Route) inter.RouteCollection {
	if c.routesMapRoutes == nil {
		c.routesMapRoutes = make(inter.MapMethodRoutes)
	}

	routesByMethod := c.routesMapRoutes[route.Method()]
	c.routesMapRoutes[route.Method()] = append(routesByMethod, route)

	c.routes = append(c.routes, route)

	return c
}

func (c *RouteCollection) Merge(routeCollection inter.RouteCollection) inter.RouteCollection {
	for _, route := range routeCollection.All() {
		c.Push(route)
	}

	return c
}

func (c RouteCollection) All() []inter.Route {
	return c.routes
}

func (c RouteCollection) Match(request inter.Request) inter.Route {
	// Normally, the routes are already decorated. In that case, no decorators
	// are present.
	DecorateRoutes(&c)
	routes, _ := request.App().MakeE("routes")
	if routes == nil {
		request.App().Singleton("routes", &c)
	}

	routesWithUrl := c.getByMethod(request.Method())

	// First, we will see if we can find a matching route for this current request
	// method. If we can, great, we can just return it so that it can be called
	// by the consumer. Otherwise we will check for routes with another verb.
	route, found := c.matchAgainstRoutes(routesWithUrl, request)

	if found {
		return route
	}

	// If no route was found we will now check if a matching route is specified by
	// another HTTP verb. If it is we will need to throw a MethodNotAllowed and
	// inform the user agent of which HTTP verb it should use for this route.
	ok := c.hasAlternateMethod(request)

	if ok {
		return c.getErrorRoute(errors.Wrap(MethodNotAllowedError, "method %s is not supported for this url", request.Method()))
	}

	return c.getErrorRoute(RouteNotFoundError)
}

// Set a group of global where patterns on the routes.
func (c *RouteCollection) Where(parameter, regex string) inter.RouteCollection {
	for _, route := range c.routes {
		route.SetConstraint(parameter, regex)
	}

	return c
}

func (c *RouteCollection) WhereMulti(constraints map[string]string) inter.RouteCollection {
	for parameter, regex := range constraints {
		c.Where(parameter, regex)
	}

	return c
}

func (c *RouteCollection) Domain(domain string) inter.RouteCollection {
	for _, route := range c.routes {
		route.SetDomain(domain)
	}

	return c
}

func (c *RouteCollection) Prefix(prefix string) inter.RouteCollection {
	for _, route := range c.routes {
		route.SetPrefix(prefix)
	}

	return c
}

func (c *RouteCollection) Name(name string) inter.RouteCollection {
	for _, route := range c.routes {
		route.SetName(name)
	}

	return c
}

func (c *RouteCollection) Middleware(middleware ...inter.HttpMiddleware) inter.RouteCollection {
	for _, route := range c.routes {
		route.SetMiddleware(middleware)
	}

	c.middlewares = append(c.middlewares, middleware...)

	return c
}

func (c *RouteCollection) WithoutMiddleware(middleware ...inter.HttpMiddleware) inter.RouteCollection {
	for _, route := range c.routes {
		route.SetExcludeMiddleware(middleware)
	}

	return c
}

// Set the (redirect) destination url
func (c *RouteCollection) setDestination(destination string) *RouteCollection {
	for _, route := range c.routes {
		route.SetDestination(destination)
	}

	return c
}

// Set the http status
func (c *RouteCollection) setStatus(status int) *RouteCollection {
	for _, route := range c.routes {
		route.SetStatus(status)
	}

	return c
}

func (c RouteCollection) getByMethod(method string) []inter.Route {
	return c.routesMapRoutes[method]
}

func (c *RouteCollection) matchAgainstRoutes(routes []inter.Route, request inter.Request) (inter.Route, bool) {
	source := request.Source()
	var match mux.RouteMatch

	for _, route := range routes {
		muxRoute := http_helper.MuxFromRoute(route)
		if err := muxRoute.GetError(); err != nil {
			return c.getErrorRoute(err), true
		}
		ok := muxRoute.Match(&source, &match)

		if ok {
			request.SetUrlValues(match.Vars)
			if request.App() == nil {
				return c.getErrorRoute(AppNotFoundError), true
			}
			request.App().Singleton("route", route)

			return route, true
		}
	}

	var nil inter.Route
	return nil, false
}

func (c RouteCollection) hasAlternateMethod(request inter.Request) bool {
	for _, route := range c.routes {
		var match mux.RouteMatch
		source := request.Source()
		ok := http_helper.MuxFromRoute(route).Match(&source, &match)
		if ok {
			return true
		}
	}

	return false
}

func flatten(collections []inter.RouteCollection) inter.RouteCollection {
	result := &RouteCollection{}

	for _, collection := range collections {
		result.Merge(collection)
	}

	return result
}

func (c RouteCollection) getErrorRoute(err error) *Route {
	status, _ := errors.FindStatus(err)
	r := &Route{
		controller: func(request inter.Request) inter.Response {
			return outcome.Html(err)
		},
		routeOptions: RouteOptions{status: status},
	}
	r.SetMiddleware(c.middlewares)
	return r
}
