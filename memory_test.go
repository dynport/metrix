package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	m := &Memory{}
	mh := &MetricHandler{}
	stats, _ := mh.Collect(m)

	assert.True(t, len(stats) > 10)
	assert.Equal(t, len(stats), 42)
	assert.Equal(t, stats[0].Key, "memory.MemTotal")
	assert.Equal(t, stats[0].Value, 502976)
}
