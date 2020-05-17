package http

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values map[string][]inter.Value

func NewUrlByValues(rawValues map[string]string) Values {
	result := Values{}

	for key, value := range rawValues {
		result.Add(key, support.NewValue(value))
	}

	return result
}

func NewUrlByMultiValues(rawMapWithValues map[string][]string) Values {
	result := Values{}

	for key, rawValues := range rawMapWithValues {
		for _, value := range rawValues {
			result.Add(key, support.NewValue(value))
		}
	}

	return result
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Values) Get(key string) inter.Value {
	vs := v[key]
	if len(vs) == 0 {
		return nil
	}
	return vs[0]
}

func (v Values) GetMulti(key string) []inter.Value {
	values, ok := v[key]
	if ! ok {
		return nil
	}

	return values
}

// Set sets the key to value. It replaces any existing
// values.
func (v Values) Set(key string, value inter.Value) {
	v[key] = []inter.Value{value}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v Values) Add(key string, value inter.Value) {
	v[key] = append(v[key], value)
}

// Del deletes the values associated with key.
func (v Values) Del(key string) {
	delete(v, key)
}
