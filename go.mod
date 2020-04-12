module github.com/lanvard/foundation

go 1.14

require (
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/http v0.0.0
	github.com/lanvard/routing v0.0.0
	github.com/lanvard/support v0.0.0
	github.com/stretchr/testify v1.5.1
)

replace (
	github.com/lanvard/contract v0.0.0 => ../contract
	github.com/lanvard/http v0.0.0 => ../http
	github.com/lanvard/routing v0.0.0 => ../routing
	github.com/lanvard/support v0.0.0 => ../support
)
