package request

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"github.com/stretchr/testify/require"
	net "net/http"
	"testing"
)

func Test_request_cookie_get_first(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home"})

	cookie := request.Cookie("latest_page")
	require.Equal(t, "home", cookie)
}

func Test_request_cookie_get_second(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home;navigated=landing_page"})

	require.Equal(t, "landing_page", request.Cookie("navigated"))
}

func Test_request_cookie_not_found_must_panic(t *testing.T) {
	request := requestByCookies([]string{})

	require.PanicsWithError(t, "http: named cookie not present", func() {
		request.Cookie("navigated")
	})
}

func Test_request_cookie_eget_first(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home"})

	value, err := request.CookieE("latest_page")
	require.NoError(t, err)
	require.Equal(t, "home", value)
}

func Test_request_cookie_eget_second(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home;navigated=landing_page"})

	value, err := request.CookieE("navigated")
	require.NoError(t, err)
	require.Equal(t, "landing_page", value)
}

func Test_request_cookie_enot_found_must_give_an_error(t *testing.T) {
	request := requestByCookies([]string{})

	value, err := request.CookieE("navigated")
	require.EqualError(t, err, "http: named cookie not present")
	require.Equal(t, "", value)
}

func requestByCookies(cookies []string) inter.Request {
	options := http.Options{
		Header: net.Header{
			"Cookie": cookies,
		},
	}
	return http.NewRequest(options)
}
