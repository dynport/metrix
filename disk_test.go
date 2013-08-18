package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDiskStat(t *testing.T) {
	mh := &MetricHandler{}
	disk := &Disk{}
	stats, _ := mh.Collect(disk)

	assert.Equal(t, len(stats), 11)
	assert.Equal(t, stats[0].Key, "disk.ReadsCompleted")
	assert.Equal(t, stats[0].Value, int64(9748))
	assert.Equal(t, stats[0].Tags["name"], "sda")
}
