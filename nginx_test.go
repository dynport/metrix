package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNginx(t *testing.T) {
	logger.LogLevel = INFO
	mh := &MetricHandler{}
	nginx := &Nginx{Raw: readFile("fixtures/nginx.status")}

	all, _ := mh.Collect(nginx)
	assert.Equal(t, len(all), 7)

	assert.Equal(t, all[0].Key, "nginx.ActiveConnections")
	assert.Equal(t, all[0].Value, 10)

	assert.Equal(t, all[6].Key, "nginx.Waiting")
	assert.Equal(t, all[6].Value, 70)
}
