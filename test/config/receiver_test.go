package config

import (
	"github.com/lanvard/foundation/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEmptyStructByEmptyKey(t *testing.T) {
	result, err := config.GetE(struct{}{}, "")
	assert.Equal(t, result, struct{}{})
	assert.Nil(t, err)
}

func TestGetFilledStructByEmptyKey(t *testing.T) {
	mockConfig := struct{ title string }{title: "The horse"}
	result, err := config.GetE(mockConfig, "")
	assert.Equal(t, result, struct{ title string }{title: "The horse"})
	assert.Nil(t, err)
}

func TestGetByNotExistingKey(t *testing.T) {
	mockConfig := struct{ title string }{title: "The horse"}
	result, err := config.GetE(mockConfig, "age")
	assert.Nil(t, result)
	assert.EqualError(t, err, "age: can't find value")
}

func TestGetByExistingKey(t *testing.T) {
	mockConfig := struct{ Title string }{Title: "The horse"}
	result, err := config.GetE(mockConfig, "Title")
	assert.Equal(t, "The horse", result)
	assert.Nil(t, err)
}

// func TestGetNestedFromStruct(t *testing.T) {
// 	mockConfig := struct{ House map[string]string }{House: map[string]string{"door": "blue"}}
// 	result, err := config.GetE(mockConfig, "House.door")
// 	assert.Equal(t, "blue", result)
// 	assert.Nil(t, err)
// }

// func TestGetByAsterisk(t *testing.T) {
// 	mockConfig := struct { Title string}{Title: "The horse"}
// 	result, err := config.GetE(mockConfig, "Title")
// 	assert.Equal(t, "The horse", result)
// 	assert.Nil(t, err)
// }

func TestGetEmptyStructByEmptyKeyWithoutError(t *testing.T) {
	result := config.Get(struct{}{}, "")
	assert.Equal(t, result, struct{}{})
}

func TestGetFilledStructByEmptyKeyWithoutError(t *testing.T) {
	mockConfig := struct{ title string }{title: "The horse"}
	result := config.Get(mockConfig, "")
	assert.Equal(t, result, struct{ title string }{title: "The horse"})
}

func TestGetByNotExistingKeyWithPanic(t *testing.T) {
	mockConfig := struct{ title string }{title: "The horse"}

	assert.PanicsWithError(t, "age: can't find value", func() {
		config.Get(mockConfig, "age")
	})
}

func TestGetByExistingKeyWithoutError(t *testing.T) {
	mockConfig := struct{ Title string }{Title: "The horse"}
	result := config.Get(mockConfig, "Title")
	assert.Equal(t, "The horse", result)
}
