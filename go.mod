module github.com/confetti-framework/foundation

go 1.15

require (
	github.com/confetti-framework/contract v0.1.2
	github.com/confetti-framework/errors v0.11.0
	github.com/confetti-framework/support v0.2.0
	github.com/confetti-framework/syslog v0.1.1
	github.com/gorilla/mux v1.8.0
	github.com/jedib0t/go-pretty/v6 v6.1.0
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.7.0
	github.com/thatisuday/commando v1.0.4 // indirect
	github.com/tidwall/gjson v1.6.8
	github.com/vigneshuvi/GoDateFormat v0.0.0-20210204121036-67364dc23c79
	golang.org/x/text v0.3.5
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/confetti-framework/contract v0.1.2 => ../contract
	github.com/confetti-framework/support v0.2.0 => ../support
)
