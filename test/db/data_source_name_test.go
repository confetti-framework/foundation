package db

import (
	"github.com/confetti-framework/foundation/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_network_address_empty(t *testing.T) {
	sql := db.MySQL{}
	assert.Equal(t, "/", sql.NetworkAddress())
}

func Test_network_address_user(t *testing.T) {
	sql := db.MySQL{
		Username: "user",
	}
	assert.Equal(t, "user@/", sql.NetworkAddress())
}

func Test_network_address_user_with_password(t *testing.T) {
	sql := db.MySQL{
		Username: "user",
		Password: "password",
	}
	assert.Equal(t, "user:password@/", sql.NetworkAddress())
}