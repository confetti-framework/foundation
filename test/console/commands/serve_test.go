package commands

import (
	"github.com/confetti-framework/foundation/console"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serve_get_name(t *testing.T) {
	assert.Equal(t, "app:serve", console.AppServe{}.Name())
}

func Test_serve_get_description(t *testing.T) {
	assert.Equal(t, "Start the http server to handle requests", console.AppServe{}.Description())
}
