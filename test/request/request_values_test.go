package request

import (
	"github.com/gorilla/mux"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/foundation/http/request_helper"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/confetti-framework/routing/outcome"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/assert"
	net "net/http"
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

	assert.Equal(t, 1432, urlValue.Int())
	assert.Equal(t, "1432", urlValue.String())
	assert.NotEqual(t, 1432, urlValue.String())
	assert.NotEqual(t, "1432", urlValue.Int())
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

	assert.Equal(t, 1432, value.Int())
	assert.Equal(t, "1432", value.String())
	assert.NotEqual(t, 1432, value.String())
	assert.NotEqual(t, "1432", value.Int())
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
		Host:   "https://api.confetti-framework.com",
		Url:    "/user/1432?test=123",
	})

	assert.Equal(t, "GET", request.Method())
	assert.True(t, request_helper.IsMethod(request, "GET"))
	assert.Equal(t, "/user/1432", request.Path())
	assert.Equal(t, "https://api.confetti-framework.com/user/1432", request.Url())
	assert.Equal(t, "https://api.confetti-framework.com/user/1432?test=123", request.FullUrl())
}

func Test_all_values(t *testing.T) {
	request := fakeRequestWithForm()
	request.App().Singleton(inter.RequestBodyDecoder, encoder.RequestWithFormToValue)

	assert.Equal(t,
		support.Map{
			"age":      support.NewValue(support.NewCollection("10")),
			"first":    support.NewValue(support.NewCollection("klaas")),
			"second":   support.NewValue(support.NewCollection("bob", "tom")),
			"language": support.NewValue(support.NewCollection("Go")),
			"name":     support.NewValue(support.NewCollection("gopher")),
		}, request.Content().Source())
}

func Test_form_values(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, 1234, request.Parameter("user_id").Int())
	assert.Equal(t, "Go", request.Content("language").String())
	assert.Equal(t, "bob", request.Content("second").String())
	assert.Equal(t, "bob", request.Content("second").Collection().First().String())
	assert.Equal(t, support.NewCollection("bob", "tom"), request.Content("second").Collection())
	assert.Equal(t, "tom", request.Content("second.1").String())
	assert.Equal(t, "tom", request.Content("").Map()["second"].Collection()[1].String())
}

func Test_form_value_not_found(t *testing.T) {
	request := fakeRequestWithForm()

	value, err := request.ParameterE("not_existing_param")
	assert.Equal(t, 0, value.Int())
	//goland:noinspection GoNilness
	assert.Equal(t, "key 'not_existing_param': can not found value", err.Error())
}

func Test_value_or(t *testing.T) {
	request := fakeRequestWithForm()

	assert.Equal(t, "Sally", request.ContentOr("fake", "Sally").String())
	assert.Equal(t, "Go", request.ContentOr("language", "PHP").String())
	assert.Equal(t, "Go", request.ContentOr("language.0", "PHP").String())

	assert.Equal(t, 12, request.ContentOr("fake", 12).Int())
	assert.Equal(t, 10, request.ContentOr("age", 12).Int())
	assert.Equal(t, 10, request.ContentOr("age.0", 12).Int())
}

func Test_request_content_type_json(t *testing.T) {
	// Given
	request := fakeRequestWithJsonBody()

	// When
	response := middleware.RequestBodyDecoder{}.Handle(request, func(request inter.Request) inter.Response {
		value := request.Content("data.foo.0.bar.1.bar")
		return outcome.Html(value)
	})
	response.SetApp(request.App())

	// Then
	assert.Equal(t, "A02", response.GetBody())
}

func fakeRequestWithForm() inter.Request {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Host:   "https://api.confetti-framework.com",
		Url:    "/user/1432?user_id=1234",
		Header: map[string][]string{"Content-Type": {"multipart/form-data; boundary=xxx"}},
		Form: url.Values{
			"age":      {"10"},
			"language": {"Go"},
			"name":     {"gopher"},
			"first":    {"klaas"},
			"second":   {"bob", "tom"},
		},
		App: foundation.NewApp(),
	})
	middleware.RequestBodyDecoder{}.Handle(request, emptyController)
	return request
}

func fakeRequestWithJsonBody() inter.Request {
	app := foundation.NewApp()
	app.Bind("outcome_html_encoders", mock.HtmlEncoders)

	return http.NewRequest(http.Options{
		App:    app,
		Method: method.Post,
		Host:   "https://api.confetti-framework.com",
		Url:    "/user/2432?comment_id=1234",
		Header: map[string][]string{
			"Content-Type": {"text/json; charset=UTF-8"},
		},
		Content: `{"data":{"foo":[{"foo":{"foo":"NL"},"bar":[{"bar":"A01"},{"bar":"A02"}]}]}}`,
	})
}

func Test_get_content_from_request_with_method_get(t *testing.T) {
	request := http.NewRequest(http.Options{
		App:    foundation.NewApp(),
		Method: method.Get,
		Host:   "https://api.confetti-framework.com",
	})

	content, err := request.ContentE("")
	assert.Equal(t, support.NewValue(nil), content)
	assert.EqualError(t, err, "no request body decoder found. Check the headers and http method")
	status, ok := errors.FindStatus(err)
	assert.Equal(t, net.StatusUnsupportedMediaType, status)
	assert.True(t, ok)
}

var emptyController = func(request inter.Request) inter.Response { return nil }
