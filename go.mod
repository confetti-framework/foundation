module github.com/lanvard/foundation

require (
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/http v0.0.0
	github.com/lanvard/routing v0.0.0
	github.com/lanvard/support v0.0.0
	github.com/stretchr/testify v1.5.1
	upspin.io v0.0.0-20200124224713-5ddde7b8e6ff // indirect
)

replace (
	github.com/lanvard/contract v0.0.0 => ../contract
	github.com/lanvard/http v0.0.0 => ../http
	github.com/lanvard/routing v0.0.0 => ../routing
	github.com/lanvard/support v0.0.0 => ../support
)

go 1.14
