package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiles(t *testing.T) {
	files := &Files{RawStatus: readFile("fixtures/lsof.txt")}
	assert.Equal(t, files.Prefix(), "files")
	assert.NotEmpty(t, files.Keys())

	mh := new(MetricHandler)
	stats, _ := mh.Collect(files)
	assert.True(t, len(stats) > 4)
	assert.Equal(t, len(stats), 74)

	names := map[string]int{}
	for _, s := range stats {
		names[s.Tags["name"]]++
	}
	assert.Equal(t, names["kworker/0"], 0)
	assert.Equal(t, names["kworker"], 8)
}

func TestParseLsofOutput(t *testing.T) {
	return
}
