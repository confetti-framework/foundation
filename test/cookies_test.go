package test

import (
	"github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCookieGetFromRequest(t *testing.T) {
	options := http.Options{
		Headers: http.Header{
			"Cookie": []string{"latest_page=home"},
		},
	}
	request := http.NewRequest(options)

	assert.Equal(t, "home", request.Cookie("latest_page").Value)
}
