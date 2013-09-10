package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCpu(t *testing.T) {
	mh := &MetricHandler{}
	stats, _ := mh.Collect(&Cpu{})
	assert.Equal(t, len(stats), 19)
	assert.Equal(t, stats[0].Key, "cpu.User")
	assert.Equal(t, stats[0].Tags["total"], "true")
	assert.Equal(t, stats[0].Value, 3700)

	assert.Equal(t, stats[14].Key, "cpu.Ctxt")
	assert.Equal(t, stats[14].Value, 957584)
}
