package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptParse(t *testing.T) {
	parser := &OptParser{}
	parser.Add("redis", "127.0.0.1:6379", "")
	parser.AddKey("postgres", "")

	parser.ProcessAll([]string{"--redis"})
	v := parser.Get("redis")
	assert.Equal(t, v, "127.0.0.1:6379")

	parser.ProcessAll([]string{"--redis", "1.2.3.4:1234"})
	v = parser.Get("redis")
	assert.Equal(t, v, "1.2.3.4:1234")

	parser.ProcessAll([]string{"--redis=1.2.3.4:1234"})
	v = parser.Get("redis")
	assert.Equal(t, v, "1.2.3.4:1234")

	parser.ProcessAll([]string{"-redis", "1.2.3.4:1234", "-postgres", "1.2.3.4"})
	v = parser.Get("postgres")
	assert.Equal(t, v, "1.2.3.4")
}
