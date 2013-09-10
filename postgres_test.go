package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePostgresUrl(t *testing.T) {
	raw := "psql://ff:ffpwd@127.0.0.1:1234/ff_test"
	u := ParsePostgresUrl(raw)
	assert.Equal(t, u.Host, "127.0.0.1")
	assert.Equal(t, u.User, "ff")
	assert.Equal(t, u.Password, "ffpwd")
	assert.Equal(t, u.Database, "ff_test")
	assert.Equal(t, u.Port, 1234)

	s, _ := u.ConnectString()
	assert.Equal(t, s, "host=127.0.0.1 dbname=ff_test user=ff password=ffpwd port=1234 sslmode=disable")
}
