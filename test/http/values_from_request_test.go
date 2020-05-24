package test

import (
	"github.com/gorilla/mux"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/method"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_number_from_uri(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/user/1432",
		Route: new(mux.Route).Path("/user/{user_id}"),
	})

	urlValue := request.UrlValue("user_id")

	assert.Equal(t, 1432, urlValue.Number())
	assert.Equal(t, "1432", urlValue.String())
	assert.NotEqual(t, 1432, urlValue.String())
	assert.NotEqual(t, "1432", urlValue.Number())
}

func Test_numbers_from_uri(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/user/1432,5423",
		Route: new(mux.Route).Path("/user/{user_ids}"),
	})

	urlValues := request.UrlValue("user_ids")

	assert.Equal(t, []int{1432, 5423}, urlValues.Numbers())
	assert.Equal(t, []string{"1432", "5423"}, urlValues.Strings())
	assert.NotEqual(t, []int{1432, 5423}, urlValues.Strings())
	assert.NotEqual(t, []string{"1432", "5423"}, urlValues.Numbers())
}

func Test_number_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users?user_id=1432",
		Route: new(mux.Route).Path("/users"),
	})

	queryValue := request.QueryValue("user_id")

	assert.Equal(t, 1432, queryValue.Number())
	assert.Equal(t, "1432", queryValue.String())
	assert.NotEqual(t, 1432, queryValue.String())
	assert.NotEqual(t, "1432", queryValue.Number())
}

func Test_numbers_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users?user_ids=1432,5423",
		Route: new(mux.Route).Path("/users"),
	})

	queryValues := request.QueryValue("user_ids")

	assert.Equal(t, []int{1432, 5423}, queryValues.Numbers())
	assert.Equal(t, []string{"1432", "5423"}, queryValues.Strings())
	assert.NotEqual(t, []int{1432, 5423}, queryValues.Strings())
	assert.NotEqual(t, []string{"1432", "5423"}, queryValues.Numbers())
}
