package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetricKeys(t *testing.T) {
	Convey("MetricKeys", t, func() {
		col := &MetricHandler{}
		es := &ElasticSearch{}
		keys := col.Keys(es)
		So(len(keys) > 0, ShouldBeTrue)
		So(keys[0], ShouldEqual, "elasticsearch.shards.Total")

	})
}
