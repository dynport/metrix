package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadAvg(t *testing.T) {
	mh := &MetricHandler{}
	stats, _ := mh.Collect(&LoadAvg{})
	assert.Equal(t, len(stats), 3)
	assert.Equal(t, stats[0].Value, 3)
	assert.Equal(t, stats[1].Value, 7)
	assert.Equal(t, stats[2].Value, 8)
}
