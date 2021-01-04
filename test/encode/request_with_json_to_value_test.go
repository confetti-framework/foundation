package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/http"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_json_to_value(t *testing.T) {
	request := http.NewRequest(http.Options{Content: `{"name":{"first":"Janet","last":"Prichard"},"age":47}`})
	value := encoder.RequestWithJsonToValue(request)

	require.Equal(t, "Janet", value.Get("name.first").String())
}

func Test_deep_json_to_value(t *testing.T) {
	content := `{
  "data": {
    "tracktraces": [
      {
        "history": [
          {
            "code": "A01",
            "status": 2,
            "main": "registered"
          },
          {
            "code": "B01",
            "status": 3,
            "main": "handed_to_carrier"
          },
          {
            "code": "J01",
            "status": 4,
            "main": "sorting"
          },
          {
            "code": "J05",
            "status": 5,
            "main": "distribution"
          }
        ]
      }
    ]
  }
}`
	request := http.NewRequest(http.Options{Content: content})
	value := encoder.RequestWithJsonToValue(request)

	require.Equal(t, "J01", value.Get("data.tracktraces.0.history.2.code").String())
}
