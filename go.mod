module github.com/lanvard/foundation

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/support v0.1.0
	github.com/stretchr/testify v1.5.1
)

replace (
	github.com/lanvard/contract v0.0.0 => ../contract
	github.com/lanvard/support v0.1.0 => ../support
)
