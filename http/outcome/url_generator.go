package outcome

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/decorator/route_decorator"
	"github.com/confetti-framework/foundation/http/http_helper"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/support"
)

type Parameters map[string]interface{}

// Receive the URL to a named route by app, name, uri parameters, query parameters
func UrlByName(app inter.App, name string, parameters ...Parameters) string {
	var pairs []string
	var result fmt.Stringer

	routes := app.Make("routes").(inter.RouteCollection)
	route, err := RouteByName(routes, name)
	if err != nil {
		panic("URL cannot be generated because " + err.Error())
	}

	UriParameters := Parameters{}
	if len(parameters) > 0 {
		UriParameters = parameters[0]
	}

	QueryParameters := Parameters{}
	if len(parameters) > 1 {
		QueryParameters = parameters[1]
	}

	// Remove Confetti custom placeholder
	route = route_decorator.UriSuffixSlash{}.Revert(route)

	muxRoute := http_helper.MuxFromRoute(route)

	for name, value := range UriParameters {
		pairs = append(pairs, name, support.NewValue(value).String())
	}

	for name, value := range QueryParameters {
		pairs = append(pairs, name, support.NewValue(value).String())
	}

	for name := range QueryParameters {
		muxRoute.Queries(name, "{"+name+"}")
	}

	muxRoute.Schemes("https")

	result, err = muxRoute.URL(pairs...)

	if err != nil {
		panic("URL cannot be generated because " + err.Error())
	}

	return result.String()
}

// Receive inter.Route by name
func RouteByName(routes inter.RouteCollection, name string) (inter.Route, error) {
	var matchedRoutes []inter.Route
	for _, route := range routes.All() {
		if route.Method() != method.Head && name == route.Name() {
			matchedRoutes = append(matchedRoutes, route)
		}
	}

	if len(matchedRoutes) == 0 {
		return nil, errors.New("no route found matching name " + name)
	}

	if len(matchedRoutes) > 1 {
		return nil, errors.New("multiple routes were found matching name " + name)
	}

	return matchedRoutes[0], nil
}
