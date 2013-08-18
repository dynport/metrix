package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMetricKeys(t *testing.T) {
	col := &MetricHandler{}
	es := &ElasticSearch{}
	keys := col.Keys(es)
	assert.True(t, len(keys) > 0)
	assert.Equal(t, keys[0], "elasticsearch.shards.Total")
}
