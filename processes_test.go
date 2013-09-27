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
	for _, stat := range stats {
		if stat.Tags["name"] == "init" && stat.Key == "processes.Utime" {
			assert.Equal(t, stat.Key, "processes.Utime")
			assert.Equal(t, stat.Value, 25)
			assert.Equal(t, stat.Tags["comm"], "(init)")
			assert.Equal(t, stat.Tags["name"], "init")
			assert.Equal(t, stat.Tags["state"], "S")
			assert.Equal(t, stat.Tags["pid"], "1")
		}
	}
}

func TestNormlizeProcessName(t *testing.T) {
	assert.Equal(t, NormalizeProcessName("(int)"), "int")
	assert.Equal(t, NormalizeProcessName("(kworker/2:1H)"), "kworker")
}

func BenchmarkNormlizeProcessName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalizeProcessName("(kworker/2:1H)")
	}
}

func BenchmarkCollect(b *testing.B) {
	mh := new(MetricHandler)
	col := &Processes{}
	for i := 0; i < b.N; i++ {
		mh.Collect(col)
	}
}
