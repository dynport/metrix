package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseElasticSearch(t *testing.T) {
	Convey("ElasticSearch", t, func() {
		mh := &MetricHandler{}
		es := &ElasticSearch{RawStatus: readFile("fixtures/elasticsearch_status.json")}

		all, _ := mh.Collect(es)

		mapped := map[string]*Metric{}
		for _, m := range all {
			k := m.Key
			if name, ok := m.Tags["index_name"]; ok {
				k = name + "." + k
			}
			mapped[k] = m
		}
		So(mapped["elasticsearch.shards.Total"].Value, ShouldEqual, int64(20))
		So(mapped["elasticsearch.shards.Successful"].Value, ShouldEqual, int64(10))
		So(mapped["index.elasticsearch.indices.index.SizeInBytes"].Value, ShouldEqual, int64(2161))
		So(mapped["index.elasticsearch.indices.merges.TotalDocs"].Value, ShouldEqual, int64(1))
		So(len(mapped) > 5, ShouldBeTrue)
	})
}
