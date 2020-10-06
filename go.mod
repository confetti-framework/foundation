module github.com/lanvard/foundation

go 1.15

require (
	github.com/gorilla/mux v1.7.4
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/routing v0.2.0
	github.com/lanvard/support v0.1.0
	github.com/lanvard/syslog v0.0.0-20201006212613-1a854c73ef48
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.5.1
	github.com/tidwall/gjson v1.6.0
)

replace (
	github.com/lanvard/contract v0.0.0 => ../contract
	github.com/lanvard/routing v0.2.0 => ../routing
	github.com/lanvard/support v0.1.0 => ../support
	github.com/lanvard/syslog v0.0.0-20201006212613-1a854c73ef48 => ../syslog
)
