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

	values := request.Value("user_ids")

	assert.Equal(t, []int{1432, 5423}, values.Numbers())
	assert.Equal(t, []string{"1432", "5423"}, values.Strings())
	assert.NotEqual(t, []int{1432, 5423}, values.Strings())
	assert.NotEqual(t, []string{"1432", "5423"}, values.Numbers())
}

func Test_number_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Uri:    "/users?user_id=1432",
		Route:  new(mux.Route).Path("/users"),
	})

	value := request.Value("user_id")

	assert.Equal(t, 1432, value.Number())
	assert.Equal(t, "1432", value.String())
	assert.NotEqual(t, 1432, value.String())
	assert.NotEqual(t, "1432", value.Number())
}

func Test_numbers_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Uri:    "/users?user_ids=1432,5423",
		Route:  new(mux.Route).Path("/users"),
	})

	values := request.Value("user_ids")

	assert.Equal(t, []int{1432, 5423}, values.Numbers())
	assert.Equal(t, []string{"1432", "5423"}, values.Strings())
	assert.NotEqual(t, []int{1432, 5423}, values.Strings())
	assert.NotEqual(t, []string{"1432", "5423"}, values.Numbers())
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
		support.Map{
			"age":      support.NewValue(support.NewCollection("10")),
			"first":    support.NewValue(support.NewCollection("klaas")),
			"second":   support.NewValue(support.NewCollection("bob", "tom")),
			"language": support.NewValue(support.NewCollection("Go")),
			"name":     support.NewValue(support.NewCollection("gopher")),
			"user_id":  support.NewValue(support.NewCollection("1234")),
		}, request.All())
}

func TestFormValues(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, 1234, request.Value("user_id").Number())
	assert.Equal(t, "Go", request.Value("language").String())
	assert.Equal(t, "bob", request.Value("second").String())
	assert.Equal(t, "bob", request.Values("second").First().String())
	assert.Equal(t, support.NewCollection("bob", "tom"), request.Values("second"))
	assert.Equal(t, "tom", request.Value("second.1").String())
}

//noinspection GoNilness
func TestFormValueNotFound(t *testing.T) {
	request := fakeRequestWithForm()

	value, err := request.Value("not_existing_param").NumberE()
	assert.Equal(t, 0, value)
	assert.Equal(t, "not_existing_param not found", err.Error())
}

func TestValueOr(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, "Sally", request.ValueOr("fake", "Sally").String())
	assert.Equal(t, "Go", request.ValueOr("language", "PHP").String())
	assert.Equal(t, "Go", request.ValueOr("language.0", "PHP").String())

	assert.Equal(t, 12, request.ValueOr("fake", 12).Number())
	assert.Equal(t, 10, request.ValueOr("age", 12).Number())
	assert.Equal(t, 10, request.ValueOr("age.0", 12).Number())
}

func fakeRequestWithForm() inter.Request {
	return http.NewRequest(http.Options{
		Method: method.Get,
		Host:   "https://api.lanvard.com",
		Uri:    "/user/1432?user_id=1234",
		Form: url.Values{
			"age":      {"10"},
			"language": {"Go"},
			"name":     {"gopher"},
			"first":    {"klaas"},
			"second":   {"bob", "tom"},
		},
	})
}
