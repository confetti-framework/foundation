package test

import (
	"github.com/gorilla/mux"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/method"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_number_from_uri(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Uri:    "/user/1432",
		Route:  new(mux.Route).Path("/user/{user_id}"),
	})

	urlValue := request.Value("user_id")

	assert.Equal(t, 1432, urlValue.Number())
	assert.Equal(t, "1432", urlValue.String())
	assert.NotEqual(t, 1432, urlValue.String())
	assert.NotEqual(t, "1432", urlValue.Number())
}

func Test_numbers_from_uri(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Uri:    "/user/1432,5423",
		Route:  new(mux.Route).Path("/user/{user_ids}"),
	})

	urlValues := request.Value("user_ids")

	assert.Equal(t, []int{1432, 5423}, urlValues.Numbers())
	assert.Equal(t, []string{"1432", "5423"}, urlValues.Strings())
	assert.NotEqual(t, []int{1432, 5423}, urlValues.Strings())
	assert.NotEqual(t, []string{"1432", "5423"}, urlValues.Numbers())
}

func Test_number_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Uri:    "/users?user_id=1432",
		Route:  new(mux.Route).Path("/users"),
	})

	queryValue := request.Value("user_id")

	assert.Equal(t, 1432, queryValue.Number())
	assert.Equal(t, "1432", queryValue.String())
	assert.NotEqual(t, 1432, queryValue.String())
	assert.NotEqual(t, "1432", queryValue.Number())
}

func Test_numbers_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Uri:    "/users?user_ids=1432,5423",
		Route:  new(mux.Route).Path("/users"),
	})

	queryValues := request.Value("user_ids")

	assert.Equal(t, []int{1432, 5423}, queryValues.Numbers())
	assert.Equal(t, []string{"1432", "5423"}, queryValues.Strings())
	assert.NotEqual(t, []int{1432, 5423}, queryValues.Strings())
	assert.NotEqual(t, []string{"1432", "5423"}, queryValues.Numbers())
}

func TestGetUrl(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Host:   "https://api.lanvard.com",
		Uri:    "/user/1432?test=123",
	})

	assert.Equal(t, "GET", request.Method())
	assert.True(t, request.IsMethod("GET"))
	assert.Equal(t, "/user/1432", request.Path())
	assert.Equal(t, "https://api.lanvard.com/user/1432", request.Url())
	assert.Equal(t, "https://api.lanvard.com/user/1432?test=123", request.FullUrl())
}

func TestAllValues(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t,
		support.Bag{
			"field1": support.Collection{support.NewValue("value1")},
			"field2": support.Collection{
				support.NewValue("initial-value2"),
				support.NewValue("value2")},
			"language": support.Collection{support.NewValue("Go")},
			"name":     support.Collection{support.NewValue("gopher")},
			"user_id":  support.Collection{support.NewValue("1234")},
		}, request.All())
}

func TestFormValues(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, 1234, request.Value("user_id").Number())
	assert.Equal(t, "Go", request.Value("language").String())
	assert.Equal(t, "initial-value2", request.Value("field2").String())
	assert.Equal(t, "initial-value2", request.Values("field2").First().String())
	assert.Equal(t, support.NewCollection("initial-value2", "value2"), request.Values("field2"))
}

//noinspection GoNilness
func TestFormValueNotFound(t *testing.T) {
	request := fakeRequestWithForm()

	_, err := request.Value("not_existing_param").NumberE()
	assert.Equal(t, "not_existing_param not found", err.Error())
}

func TestValueOr(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, "Sally", request.ValueOr("fake", "Sally").String())
	assert.Equal(t, "Go", request.ValueOr("language", "PHP").String())
}

func fakeRequestWithForm() inter.Request {
	return http.NewRequest(http.Options{
		Method: method.Get,
		Host:   "https://api.lanvard.com",
		Uri:    "/user/1432?user_id=1234",
		Form: url.Values{
			"language": []string{"Go"},
			"name":     []string{"gopher"},
			"field1":   []string{"value1"},
			"field2":   []string{"initial-value2", "value2"},
		},
	})
}