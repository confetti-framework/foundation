package http

import (
	"github.com/lanvard/contract/entity"
	"net/url"
	"strconv"
	"strings"
)

type UrlValues struct {
	source url.Values
}

func NewUrlByValues(values map[string]string) UrlValues {
	source := url.Values{}

	for key, value := range values {
		source.Add(key, value)
	}

	return UrlValues{source: source}
}

func NewUrlByMultiValues(values map[string][]string) UrlValues {
	return UrlValues{source: values}
}

func (u UrlValues) Source() url.Values {
	return u.source
}

func (u UrlValues) String(key string) string {
	value, err := u.StringE(key)
	if nil != err {
		panic(err)
	}

	return value
}

func (u UrlValues) StringE(key string) (string, error) {
	value, found := u.get(key)

	if ! found {
		return "", u.getException(key)
	}

	return value, nil
}

func (u UrlValues) Strings(key string) []string {
	values, err := u.StringsE(key)
	if nil != err {
		panic(err)
	}

	return values
}

func (u UrlValues) StringsE(key string) ([]string, error) {
	values, found := u.getSlice(key)
	if ! found {
		return nil, u.getException(key)
	}

	return values, nil
}

func (u UrlValues) Number(key string) int {
	values, err := u.NumberE(key)
	if nil != err {
		panic(err)
	}

	return values
}

func (u UrlValues) NumberE(key string) (int, error) {
	result, found := u.get(key)
	if ! found {
		return 0, u.getException(key)
	}

	if "" == result {
		return 0, nil
	}

	return strconv.Atoi(result)
}

func (u UrlValues) Numbers(key string) []int {
	values, err := u.NumbersE(key)
	if nil != err {
		panic(err)
	}

	return values
}

func (u UrlValues) NumbersE(key string) ([]int, error) {
	rawValues, found := u.getSlice(key)
	if ! found {
		return nil, u.getException(key)
	}

	var result []int

	for _, rawValue := range rawValues {
		value, err := strconv.Atoi(rawValue)
		if nil != err {
			return nil, err
		}

		result = append(result, value)
	}

	return result, nil
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (u UrlValues) get(key string) (value string, found bool) {
	if u.source == nil {
		return "", false
	}

	vs := u.Source()[key]
	if len(vs) == 0 {
		return "", false
	}

	return vs[0], true
}

func (u UrlValues) getSlice(key string) (values []string, found bool) {
	rawValues, found := u.get(key)
	if ! found {
		return nil, false
	}

	return strings.Split(rawValues, ","), true
}

func (u UrlValues) getException(key string) entity.Exception {
	return entity.Exception{Code: "value_not_found_in_url", Human: "key %key not found in uri",
		Values: map[string]string{"key": key}}
}
