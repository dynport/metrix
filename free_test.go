package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFree(t *testing.T) {
	mh := new(MetricHandler)
	free := &Free{RawStatus: readFile("fixtures/free.txt")}
	stats, _ := mh.Collect(free)
	agg := aggregateStats(stats)
	assert.Equal(t, agg["free.mem.total"], 8177440)
	assert.Equal(t, agg["free.mem.used"], 7097488)
	assert.Equal(t, agg["free.mem.free"], 1079952)
	assert.Equal(t, agg["free.mem.buffers"], 1230276)
	assert.Equal(t, agg["free.mem.cached"], 1782336)

	assert.Equal(t, agg["free.swap.total"], 522236)
	assert.Equal(t, agg["free.swap.used"], 22176)
	assert.Equal(t, agg["free.swap.free"], 500060)
}
