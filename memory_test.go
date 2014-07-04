package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func aggregateStats(stats []*Metric) map[string]int64 {
	agg := map[string]int64{}
	for _, stat := range stats {
		agg[stat.Key] = stat.Value
	}
	return agg
}

func TestMemory(t *testing.T) {
	Convey("Memory", t, func() {
		m := &Memory{}
		mh := &MetricHandler{}
		stats, _ := mh.Collect(m)

		So(len(stats) > 10, ShouldBeTrue)
		So(len(stats), ShouldEqual, 42)
		So(stats[0].Key, ShouldEqual, "memory.MemTotal")
		So(stats[0].Value, ShouldEqual, 502976)

		agg := aggregateStats(stats)
		So(agg["memory.Committed_AS"], ShouldEqual, 350856)
	})
}
