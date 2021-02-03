package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
	"strings"
)

type Route struct {
	uri          string
	domain       string
	method       string
	controller   inter.Controller
	routeOptions RouteOptions
	middlewares  []inter.HttpMiddleware
}

func NewRoute(url string, method string, controller inter.Controller) inter.Route {
	return &Route{uri: url, method: method, controller: controller}
}

func (r Route) Uri() string {
	return r.uri
}

func (r *Route) SetUri(uri string) inter.Route {
	r.uri = uri

	return r
}

func (r Route) Domain() string {
	return r.domain
}

func (r *Route) SetDomain(domain string) inter.Route {
	r.domain = domain

	return r
}

func (r Route) Method() string {
	return r.method
}

func (r Route) RouteOptions() inter.RouteOptions {
	return r.routeOptions
}

func (r Route) Controller() inter.Controller {
	return r.controller
}

func (r *Route) SetPrefix(prefix string) inter.Route {
	// If a prefix is added at a higher level, it must apply at the beginning of the url
	r.routeOptions.prefixes = append([]string{prefix}, r.routeOptions.prefixes...)

	return r
}

func (r *Route) SetDestination(destination string) inter.Route {
	r.routeOptions.destination = destination

	return r
}

func (r *Route) SetStatus(status int) inter.Route {
	r.routeOptions.status = status

	return r
}

func (r Route) Constraint() map[string]string {
	return r.routeOptions.constraints
}

// Get the name of the route instance
func (r Route) Name() string {
	return r.routeOptions.name
}

// Push or change the route name.
func (r *Route) SetName(name string) inter.Route {
	r.routeOptions.name = name + r.routeOptions.name

	return r
}

// Determine whether the route's name matches the given patterns
func (r Route) Named(patterns ...string) bool {
	name := r.routeOptions.name
	if name == "" {
		return false
	}

	for _, pattern := range patterns {
		if strings.Contains(name, pattern) {
			return true
		}
	}

	return false
}

func (r *Route) SetConstraint(parameter string, regex string) inter.Route {
	if r.routeOptions.constraints == nil {
		r.routeOptions.constraints = make(map[string]string)
	}

	r.routeOptions.constraints[parameter] = regex

	return r
}

func (r Route) Middleware() []inter.HttpMiddleware {
	return r.middlewares
}

func (r *Route) SetMiddleware(middlewaresToStore []inter.HttpMiddleware) inter.Route {
	excluded := getRouteNames(r.routeOptions.excludeMiddlewares)

	// validate and sort the middlewares
	var validMiddlewares []inter.HttpMiddleware
	for _, middleware := range middlewaresToStore {
		if !excluded.Contains(support.Name(middleware)) {
			// put the middleware first in the slice
			validMiddlewares = append(validMiddlewares, middleware)
		}
	}

	r.middlewares = append(validMiddlewares, r.middlewares...)

	return r
}

func (r *Route) SetExcludeMiddleware(middlewares []inter.HttpMiddleware) inter.Route {
	r.routeOptions.excludeMiddlewares = append(r.routeOptions.excludeMiddlewares, middlewares...)

	return r
}

type RouteOptions struct {
	prefixes           []string
	destination        string
	status             int
	constraints        map[string]string      // The regular expression requirements
	name               string                 // Named routes allow the convenient generation of URLs or redirects
	excludeMiddlewares []inter.HttpMiddleware // prevent the middleware from being applied to this route
}

func (r RouteOptions) Prefixes() []string {
	return r.prefixes
}

func (r RouteOptions) Status() int {
	return r.status
}

func getRouteNames(middlewares []inter.HttpMiddleware) support.Collection {
	names := support.NewCollection()

	for _, middleware := range middlewares {
		names = names.Push(support.Name(middleware))
	}

	return names
}
