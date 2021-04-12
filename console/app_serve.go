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
	Host string `short:"h" flag:"host" description:"The host address to serve the application on [default: \"127.0.0.1\"]"`
	Port int `short:"p" flag:"port" description:"The port to serve the application on"`
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

	c.Line("\u001B[32mStarting %s server:\u001B[0m http://%s", name, s.getAddr(c.App()))
	server := &net.Server{
		Addr:         s.getAddr(c.App()),
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
	return strconv.Itoa(port)
}

func (s AppServe) getHostAddr(app inter.App) string {
	var host string
	if len(s.Host) != 0 {
		host = s.Host
	} else {
		host = "127.0.0.1"
	}
	return host
}

func (s AppServe) getAddr(app inter.App) string {
	return s.getHostAddr(app) + ":" + s.getPortAddr(app)
}
