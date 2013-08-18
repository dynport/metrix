package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseRedis(t *testing.T) {
	logger.LogLevel = WARN
	mh := &MetricHandler{}
	es := &Redis{ Raw: readFile("fixtures/redis_info.txt") }

	all, _ := mh.Collect(es)
	assert.True(t, len(all) > 0)
}
