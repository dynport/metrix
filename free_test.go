package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var freeResult = map[string]int{
	"free.mem.total":   8177440,
	"free.mem.used":    7097488,
	"free.mem.free":    1079952,
	"free.mem.buffers": 1230276,
	"free.mem.cached":  1782336,
	"free.swap.total":  522236,
	"free.swap.used":   22176,
	"free.swap.free":   500060,
}

func TestFree(t *testing.T) {
	mh := new(MetricHandler)
	free := &Free{RawStatus: readFile("fixtures/free.txt")}
	stats, _ := mh.Collect(free)
	agg := aggregateStats(stats)
	for k, v := range freeResult {
		assert.Equal(t, agg[k], v, fmt.Sprintf("expected %s to have value %v but was %v", k, v, agg[k]))
	}
}
