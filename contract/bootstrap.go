package contract

import "github.com/lanvard/foundation"

type Bootstrap interface {
	Bootstrap(app *foundation.Application) *foundation.Application
}
