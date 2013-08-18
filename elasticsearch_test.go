package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseElasticSearch(t *testing.T) {
	mh := &MetricHandler{}
	es := &ElasticSearch{ RawStatus: readFile("fixtures/elasticsearch_status.json") }

	all, _ := mh.Collect(es)

	mapped := map[string]*Metric {}
	for _, m := range all {
		k := m.Key
		if name, ok := m.Tags["index_name"]; ok {
			k = name + "." + k
		}
		mapped[k] = m
	}
	assert.Equal(t, mapped["elasticsearch.shards.Total"].Value, int64(20))
	assert.Equal(t, mapped["elasticsearch.shards.Successful"].Value, int64(10))
	assert.Equal(t, mapped["index.elasticsearch.indices.index.SizeInBytes"].Value, int64(2161))
	assert.Equal(t, mapped["index.elasticsearch.indices.merges.TotalDocs"].Value, int64(1))
	assert.True(t, len(mapped) > 5)
}
