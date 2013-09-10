package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcesses(t *testing.T) {
	mh := new(MetricHandler)
	col := &Processes{}
	stats, _ := mh.Collect(col)
	assert.True(t, len(stats) > 0)
	assert.Equal(t, len(stats), 980)
	assert.Equal(t, stats[0].Key, "processes.Utime")
	assert.Equal(t, stats[0].Value, 25)
	assert.Equal(t, stats[0].Tags["comm"], "(init)")
	assert.Equal(t, stats[0].Tags["name"], "init")
	assert.Equal(t, stats[0].Tags["state"], "S")
	assert.Equal(t, stats[0].Tags["pid"], "1")
}

func TestNormlizeProcessName(t *testing.T) {
	assert.Equal(t, NormalizeProcessName("(int)"), "int")
	assert.Equal(t, NormalizeProcessName("(kworker/2:1H)"), "kworker")
}
