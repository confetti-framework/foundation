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
		Url:    "/users/1432",
		Route: new(mux.Route).Path("/users/{user_id}"),
	})

	urlValues := request.UrlValues()

	assert.Equal(t, 1432, urlValues.Number("user_id"))
	assert.Equal(t, "1432", urlValues.String("user_id"))
	assert.NotEqual(t, 1432, urlValues.String("user_id"))
	assert.NotEqual(t, "1432", urlValues.Number("user_id"))
}

func Test_numbers_from_uri(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users/1432,5423",
		Route: new(mux.Route).Path("/users/{user_ids}"),
	})

	urlValues := request.UrlValues()

	assert.Equal(t, []int{1432, 5423}, urlValues.Numbers("user_ids"))
	assert.Equal(t, []string{"1432", "5423"}, urlValues.Strings("user_ids"))
	assert.NotEqual(t, []int{1432, 5423}, urlValues.Strings("user_ids"))
	assert.NotEqual(t, []string{"1432", "5423"}, urlValues.Numbers("user_ids"))
}

func Test_number_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users?user_id=1432",
		Route: new(mux.Route).Path("/users"),
	})

	queryValues := request.QueryValues()

	assert.Equal(t, 1432, queryValues.Number("user_id"))
	assert.Equal(t, "1432", queryValues.String("user_id"))
	assert.NotEqual(t, 1432, queryValues.String("user_id"))
	assert.NotEqual(t, "1432", queryValues.Number("user_id"))
}

func Test_numbers_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users?user_ids=1432,5423",
		Route: new(mux.Route).Path("/users"),
	})

	queryValues := request.QueryValues()

	assert.Equal(t, []int{1432, 5423}, queryValues.Numbers("user_ids"))
	assert.Equal(t, []string{"1432", "5423"}, queryValues.Strings("user_ids"))
	assert.NotEqual(t, []int{1432, 5423}, queryValues.Strings("user_ids"))
	assert.NotEqual(t, []string{"1432", "5423"}, queryValues.Numbers("user_ids"))
}
