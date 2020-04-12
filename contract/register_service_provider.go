package contract

import "github.com/lanvard/foundation"

type RegisterServiceProvider interface {
	Register(app *foundation.Application) *foundation.Application
}
