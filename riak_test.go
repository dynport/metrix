package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRiakMetrics(t *testing.T) {
	Convey("RiakMetrics", t, func() {
		mh := &MetricHandler{}
		col := &Riak{Raw: readFile("fixtures/riak.json")}
		metrics, _ := mh.Collect(col)

		So(len(metrics), ShouldEqual, 126)
		m1 := metrics[0]
		So(m1.Key, ShouldEqual, "riak.VNodeGets")
		So(m1.Value, ShouldEqual, 241)

	})
}
