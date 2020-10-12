package request

import (
	"github.com/gorilla/mux"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/method"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/foundation/http/request_helper"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_number_from_uri(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/user/1432",
		Route:  new(mux.Route).Path("/user/{user_id}"),
	})

	urlValue := request.Parameter("user_id")

	assert.Equal(t, 1432, urlValue.Number())
	assert.Equal(t, "1432", urlValue.String())
	assert.NotEqual(t, 1432, urlValue.String())
	assert.NotEqual(t, "1432", urlValue.Number())
}

func Test_numbers_from_uri(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/user/1432,5423",
		Route:  new(mux.Route).Path("/user/{user_ids}"),
	})

	values := request.Parameter("user_ids")

	assert.Equal(t, []interface{}{"1432", "5423"}, values.Split(",").Raw())
	assert.NotEqual(t, []int{1432, 5423}, values.Split(",").Raw())
}

func Test_number_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users?user_id=1432",
		Route:  new(mux.Route).Path("/users"),
	})

	value := request.Parameter("user_id")

	assert.Equal(t, 1432, value.Number())
	assert.Equal(t, "1432", value.String())
	assert.NotEqual(t, 1432, value.String())
	assert.NotEqual(t, "1432", value.Number())
}

func Test_numbers_from_query(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users?user_ids=1432,5423",
		Route:  new(mux.Route).Path("/users"),
	})

	values := request.Parameter("user_ids")

	assert.Equal(t, []interface{}{"1432", "5423"}, values.Split(",").Raw())
	assert.NotEqual(t, []int{1432, 5423}, values.Split(",").Raw())
}

func Test_get_url(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Host:   "https://api.lanvard.com",
		Url:    "/user/1432?test=123",
	})

	assert.Equal(t, "GET", request.Method())
	assert.True(t, request_helper.IsMethod(request, "GET"))
	assert.Equal(t, "/user/1432", request.Path())
	assert.Equal(t, "https://api.lanvard.com/user/1432", request.Url())
	assert.Equal(t, "https://api.lanvard.com/user/1432?test=123", request.FullUrl())
}

func Test_all_values(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t,
		support.Map{
			"age":      support.NewValue(support.NewCollection("10")),
			"first":    support.NewValue(support.NewCollection("klaas")),
			"second":   support.NewValue(support.NewCollection("bob", "tom")),
			"language": support.NewValue(support.NewCollection("Go")),
			"name":     support.NewValue(support.NewCollection("gopher")),
		}, request.Body().Source())
}

func Test_form_values(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, 1234, request.Parameter("user_id").Number())
	assert.Equal(t, "Go", request.Body("language").String())
	assert.Equal(t, "bob", request.Body("second").String())
	assert.Equal(t, "bob", request.Body("second").Collection().First().String())
	assert.Equal(t, support.NewCollection("bob", "tom"), request.Body("second").Collection())
	assert.Equal(t, "tom", request.Body("second.1").String())
	assert.Equal(t, "tom", request.Body("").Map()["second"].Collection()[1].String())
}

func Test_form_value_not_found(t *testing.T) {
	request := fakeRequestWithForm()

	value, err := request.Parameter("not_existing_param").NumberE()
	assert.Equal(t, 0, value)
	//goland:noinspection GoNilness
	assert.Equal(t, "no value found with key 'not_existing_param'", err.Error())
}

func Test_value_or(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, "Sally", request.BodyOr("fake", "Sally").String())
	assert.Equal(t, "Go", request.BodyOr("language", "PHP").String())
	assert.Equal(t, "Go", request.BodyOr("language.0", "PHP").String())

	assert.Equal(t, 12, request.BodyOr("fake", 12).Number())
	assert.Equal(t, 10, request.BodyOr("age", 12).Number())
	assert.Equal(t, 10, request.BodyOr("age.0", 12).Number())
}

func Test_request_content_type_json(t *testing.T) {
	// Given
	request := fakeRequestWithJsonBody()

	// When
	response := middleware.RequestBodyDecoder{}.Handle(request, func(request inter.Request) inter.Response {
		value := request.Body("data.foo.0.bar.1.bar")
		return outcome.Html(value)
	})
	response.SetApp(request.App())

	// Then
	assert.Equal(t, "A02", response.Content())
}

func fakeRequestWithForm() inter.Request {
	return http.NewRequest(http.Options{
		Method: method.Get,
		Host:   "https://api.lanvard.com",
		Url:    "/user/1432?user_id=1234",
		Form: url.Values{
			"age":      {"10"},
			"language": {"Go"},
			"name":     {"gopher"},
			"first":    {"klaas"},
			"second":   {"bob", "tom"},
		},
	})
}

func fakeRequestWithJsonBody() inter.Request {
	return http.NewRequest(http.Options{
		App:    foundation.NewApp(),
		Method: method.Get,
		Host:   "https://api.lanvard.com",
		Url:    "/user/2432?comment_id=1234",
		Header: map[string][]string{
			"Content-Type": {"text/json; charset=UTF-8"},
		},
		Content: `{"data":{"foo":[{"foo":{"foo":"NL"},"bar":[{"bar":"A01"},{"bar":"A02"}]}]}}`,
	})
}
