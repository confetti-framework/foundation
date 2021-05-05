package http

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	net "net/http"
	"strings"
)

func HandleHttpKernel(app inter.App, response net.ResponseWriter, request *net.Request) {
	/*
	   |--------------------------------------------------------------------------
	   | Register The Response Writer
	   |--------------------------------------------------------------------------
	   |
	   | Confetti only uses the response writer here in main.go. But we register
	   | the response writer if you need it anyway
	   |
	*/
	app.Singleton(
		(*net.ResponseWriter)(nil),
		response,
	)

	/*
	   |--------------------------------------------------------------------------
	   | Run The Application
	   |--------------------------------------------------------------------------
	   |
	   | Once we have the application, we can handle the incoming request
	   | through the kernel, and send the associated response back to
	   | the client allowing them to enjoy the creative
	   | and awesome application we have prepared for them.
	   |
	*/
	kernel := app.Make((*inter.HttpKernel)(nil)).(inter.HttpKernel)

	appRequest := NewRequest(Options{App: app, Source: *request})

	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				rec = errors.WithStack(err)
			}
			appResponse := kernel.RecoverFromMiddlewarePanic(rec)
			exposeResponse(response, appResponse)
		}
	}()

	appResponse := kernel.Handle(appRequest)

	exposeResponse(response, appResponse)
}

func exposeResponse(response net.ResponseWriter, appResponse inter.Response) {
	// Add HTTP headers
	for key, values := range appResponse.GetHeaders() {
		response.Header().Add(key, strings.Join(values, "; "))
	}

	// Add HTTP status
	response.WriteHeader(appResponse.GetStatus())

	// Add HTTP body
	_, err := response.Write([]byte(appResponse.GetBody()))
	if err != nil {
		panic(err)
	}
}
