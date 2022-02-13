package console

import (
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/confetti-framework/contract/inter"
	"github.com/jedib0t/go-pretty/v6/table"
)

// RouteListCommand to give you an example of what a command might look like.
type RouteListCommand struct {
}

func (e RouteListCommand) Name() string {
	return "route:list"
}

func (e RouteListCommand) Description() string {
	return "Displays a list of all registered routes"
}

// Handle contains the logic of the command
func (e RouteListCommand) Handle(c inter.Cli) inter.ExitCode {
	t := c.Table()

	var routes inter.RouteCollection
	list, err := c.App().MakeE("routes")

	if err != nil {
		c.Error("Could not list routes: %s", err)
		return inter.Failure
	}

	routes = list.(inter.RouteCollection)

	t.AppendHeader(table.Row{"\x1b[33mMethod\x1b[0m", "\x1b[33mUri\x1b[0m", "\x1b[33mController\x1b[0m", "\x1b[33mName\x1b[0m"})

	for _, route := range routes.All() {
		controller := route.Controller()
		controllerName := runtime.FuncForPC(reflect.ValueOf(controller).Pointer()).Name()

		t.AppendRow(table.Row{
			route.Method(),
			getCleanRouteUri(route),
			controllerName,
			route.Name(),
		})
	}

	t.Render()

	return inter.Success
}

func getCleanRouteUri(route inter.Route) string {
	uri := route.Uri()

	for _, prefix := range route.RouteOptions().Prefixes() {
		uri = path.Join(prefix, uri)
	}

	return strings.Replace(uri, "{allow_slash:\\/?}", "", 1)
}
