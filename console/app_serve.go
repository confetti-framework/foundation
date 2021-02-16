package console

import (
	"github.com/confetti-framework/contract/inter"
	"io"
)

type AppServe struct {
	Port int `short:"p" flag:"--port"`
}

func (s AppServe) Name() string {
	return "app:serve"
}

func (s AppServe) Description() string {
	return "Start the http server to handle requests."
}

func (s AppServe) Handle(app inter.App, output io.Writer) inter.ExitCode {
	println("the port:", s.Port)

	return inter.Success
}
