package console

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	net "net/http"
	"strconv"
	"time"
)

// AppServe starts the http server to handle requests.
type AppServe struct {
	Port int `short:"p" flag:"port"`
}

// Name of the command
func (s AppServe) Name() string {
	return "app:serve"
}

// Description of the command
func (s AppServe) Description() string {
	return "Start the http server to handle requests."
}

// Handle contains the logic of the command
func (s AppServe) Handle(c inter.Cli) inter.ExitCode {
	name := c.App().Make("config.App.Name").(string)
	appProvider := c.App().Make(inter.AppProvider).(func() inter.App)

	// This bootstraps the framework and gets it ready for use, then it will load up
	// this application so that we can run it and send the responses back to the
	// user.
	handler := func(response net.ResponseWriter, request *net.Request) {
		app := appProvider()
		http.HandleHttpKernel(app, response, request)
	}

	c.Info("Start %s to handle requests", name)
	server := &net.Server{
		Addr:         s.getPortAddr(c.App()),
		Handler:      net.HandlerFunc(handler),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != net.ErrServerClosed {
		c.Error("Could not %s", err)
		return inter.Failure
	}

	c.Info("Server stopped")

	return inter.Success
}

func (s AppServe) getPortAddr(app inter.App) string {
	var port int
	if s.Port != 0 {
		port = s.Port
	} else {
		port = app.Make("config.App.Port").(int)
	}
	return ":" + strconv.Itoa(port)
}
