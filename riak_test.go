package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRiakMetrics(t *testing.T) {
	mh := &MetricHandler{}
	col := &Riak{ Raw: readFile("fixtures/riak.json") }
	metrics, _ := mh.Collect(col)

	assert.Equal(t, len(metrics), 126)
	m1 := metrics[0]
	assert.Equal(t, m1.Key, "riak.VNodeGets")
	assert.Equal(t, m1.Value, 241)
}
