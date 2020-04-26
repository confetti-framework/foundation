package method

import "net/http"

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	Get     = http.MethodGet
	Head    = http.MethodHead
	Post    = http.MethodPost
	Put     = http.MethodPut
	Patch   = http.MethodPatch // RFC 578
	Delete  = http.MethodDelete
	Connect = http.MethodConnect
	Options = http.MethodOptions
	Trace   = http.MethodTrace
)
