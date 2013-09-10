package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRedis(t *testing.T) {
	logger.LogLevel = WARN
	mh := &MetricHandler{}
	es := &Redis{Raw: readFile("fixtures/redis_info.txt")}

	all, _ := mh.Collect(es)
	assert.True(t, len(all) > 0)
}
