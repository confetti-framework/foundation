package contract

import "github.com/lanvard/foundation"

type BootServiceProvider interface {
	Boot(app *foundation.Application) *foundation.Application
}
