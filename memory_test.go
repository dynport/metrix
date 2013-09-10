package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func aggregateStats(stats []*Metric) map[string]int64 {
	agg := map[string]int64{}
	for _, stat := range stats {
		agg[stat.Key] = stat.Value
	}
	return agg
}

func TestMemory(t *testing.T) {
	m := &Memory{}
	mh := &MetricHandler{}
	stats, _ := mh.Collect(m)

	assert.True(t, len(stats) > 10)
	assert.Equal(t, len(stats), 42)
	assert.Equal(t, stats[0].Key, "memory.MemTotal")
	assert.Equal(t, stats[0].Value, 502976)

	agg := aggregateStats(stats)
	assert.Equal(t, agg["memory.Committed_AS"], 350856)
}
