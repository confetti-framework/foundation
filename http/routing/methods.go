package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/foundation/http/outcome"
	net "net/http"
)

// Source of the methods supported by the router.
var allMethods = []string{
	method.Get,
	method.Head,
	method.Post,
	method.Put,
	method.Patch,
	method.Delete,
	method.Options,
}

// Register new GET routes
func Get(uri string, controller inter.Controller) *RouteCollection {
	methods := []string{method.Get, method.Head}
	return createRoutes(methods, uri, controller)
}

// Register new POST routes
func Post(uri string, controller inter.Controller) *RouteCollection {
	return createRoute(method.Post, uri, controller)
}

// Register new PUT routes
func Put(uri string, controller inter.Controller) *RouteCollection {
	return createRoute(method.Put, uri, controller)
}

// Register new PATCH routes
func Patch(uri string, controller inter.Controller) *RouteCollection {
	return createRoute(method.Patch, uri, controller)
}

// Register new DELETE routes
func Delete(uri string, controller inter.Controller) *RouteCollection {
	return createRoute(method.Delete, uri, controller)
}

// Register new OPTIONS routes
func Options(uri string, controller inter.Controller) *RouteCollection {
	return createRoute(method.Options, uri, controller)
}

// Register a new route responding to all methods
func Any(uri string, controller inter.Controller) *RouteCollection {
	return createRoutes(allMethods, uri, controller)
}

// Register a new route with multiple methods
func Match(methods []string, uri string, controller inter.Controller) *RouteCollection {
	return createRoutes(methods, uri, controller)
}

func View(uri string, view inter.View) inter.RouteCollection {
	return Get(uri, func(request inter.Request) inter.Response {
		return outcome.Html(view)
	})
}

// Create a redirect from one URI to another.
func Redirect(uri string, destination string, status int) *RouteCollection {
	return Any(uri, redirectController).setDestination(destination).setStatus(status)
}

// Create a permanent redirect from one URI to another.
func RedirectPermanent(uri string, destination string) *RouteCollection {
	return Redirect(uri, destination, net.StatusMovedPermanently)
}

// Create a temporary redirect from one URI to another.
func RedirectTemporary(uri string, destination string) *RouteCollection {
	return Redirect(uri, destination, net.StatusFound)
}

func Fallback(controller inter.Controller) inter.RouteCollection {
	return Any("/{any}", controller).Where("any", ".*")
}
