package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_add_simple_cookie_to_response(t *testing.T) {
	cookie := http.Cookie{Name: "flow_id", Value: "123"}

	response := getResponse().Cookie(cookie)
	require.Equal(t, "123", response.GetCookies()[0].Value)
}

func Test_add_multiple_cookies(t *testing.T) {
	cookieFlow := http.Cookie{Name: "flow_id", Value: "123"}
	cookieRoom := http.Cookie{Name: "room", Value: "8"}

	response := getResponse().Cookie(cookieFlow, cookieRoom)
	require.Equal(t, "123", response.GetCookies()[0].Value)
	require.Equal(t, "8", response.GetCookies()[1].Value)
}

func getResponse() inter.Response {
	return outcome.NewResponse(outcome.Options{})
}
