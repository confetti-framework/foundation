module github.com/lanvard/foundation

go 1.15

require (
	github.com/gorilla/mux v1.7.4
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/errors v0.9.1
	github.com/lanvard/routing v0.2.0
	github.com/lanvard/support v0.1.0
	github.com/lanvard/syslog v0.0.0-20201006215111-98d4d91dbaa8
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/gjson v1.6.0
	github.com/vigneshuvi/GoDateFormat v0.0.0-20190923034126-379ee8a8c45f
)

replace (
	github.com/lanvard/contract v0.0.0 => ../contract
	github.com/lanvard/routing v0.2.0 => ../routing
	github.com/lanvard/support v0.1.0 => ../support
	github.com/lanvard/errors v0.9.1 => ../errors
	github.com/lanvard/syslog v0.0.0-20201006215111-98d4d91dbaa8 => ../syslog
)
