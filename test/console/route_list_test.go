package console

import (
	"bytes"
	"testing"

	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
)

func Test_route_list_without_routes_errors(t *testing.T) {
	writer, app := setUp()
	var writerErr bytes.Buffer

	app.Bind("config.App.OsArgs", []interface{}{"/main", "route:list"})

	code := console.Kernel{
		App:       app,
		Writer:    &writer,
		WriterErr: &writerErr,
		Commands:  []inter.Command{console.RouteListCommand{}},
	}.Handle()

	require.Equal(t, inter.Failure, code)
	require.Contains(t, writerErr.String(), "Could not list routes: get instance 'routes' from container: key 'routes': can not found value")
}

func storeUsers(request inter.Request) inter.Response { return nil }

func Test_route_list_shows_routes(t *testing.T) {
	writer, app := setUp()
	var writerErr bytes.Buffer

	app.Bind("config.App.OsArgs", []interface{}{"/main", "route:list"})
	app.Bind("routes", routing.NewRouteCollection(routing.Group(
		routing.Get("/", func(request inter.Request) inter.Response { return nil }),
		routing.Post("/users", storeUsers).Name("users.store"),
	).Prefix("/api")))

	code := console.Kernel{
		App:       app,
		Writer:    &writer,
		WriterErr: &writerErr,
		Commands:  []inter.Command{console.RouteListCommand{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	result := TrimDoubleSpaces(writer.String())
	require.Contains(t, result, "\x1b[33mMETHOD\x1b[0m \x1b[33mURI\x1b[0m \x1b[33mCONTROLLER\x1b[0m \x1b[33mNAME\x1b[0m")
	require.Contains(t, result, "GET /api github.com/confetti-framework/foundation/test/console.Test_route_list_shows_routes.func1")
	require.Contains(t, result, "HEAD /api github.com/confetti-framework/foundation/test/console.Test_route_list_shows_routes.func1")
	require.Contains(t, result, "POST /api/users github.com/confetti-framework/foundation/test/console.storeUsers users.store")
}