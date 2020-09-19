package test

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestCookieGetFirst(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home"})

	cookie := request.Cookie("latest_page")
	assert.Equal(t, "home", cookie)
}

func TestRequestCookieGetSecond(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home;navigated=landing_page"})

	assert.Equal(t, "landing_page", request.Cookie("navigated"))
}

func TestRequestCookieNotFoundMustPanic(t *testing.T) {
	request := requestByCookies([]string{})

	assert.PanicsWithError(t, "http: named cookie not present", func() {
		request.Cookie("navigated")
	})
}

func TestRequestCookieEGetFirst(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home"})

	value, err := request.CookieE("latest_page")
	assert.Nil(t, err)
	assert.Equal(t, "home", value)
}

func TestRequestCookieEGetSecond(t *testing.T) {
	request := requestByCookies([]string{"latest_page=home;navigated=landing_page"})

	value, err := request.CookieE("navigated")
	assert.Nil(t, err)
	assert.Equal(t, "landing_page", value)
}

func TestRequestCookieENotFoundMustGiveAnError(t *testing.T) {
	request := requestByCookies([]string{})

	value, err := request.CookieE("navigated")
	assert.EqualError(t, err, "http: named cookie not present")
	assert.Equal(t, "", value)
}

func requestByCookies(cookies []string) inter.Request {
	options := http.Options{
		Headers: http.Header{
			"Cookie": cookies,
		},
	}
	return http.NewRequest(options)
}
