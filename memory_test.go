package main

import (
	"io/ioutil"
	"sort"
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

		sort.Sort(stats)
		So(stats[0].Key, ShouldEqual, "memory.active")
		So(stats[0].Value, ShouldEqual, 61740)

		agg := aggregateStats(stats)

		So(agg["memory.committed_as"], ShouldEqual, 350856)

		Convey("Load", func() {
			b := mustRead(t, "proc/meminfo")
			m := &Memory{}
			e := m.Load(b)
			So(e, ShouldBeNil)
			So(m.MemTotal, ShouldEqual, 502976)
			So(m.Cached, ShouldEqual, 31436)
		})
	})
}

func mustRead(t *testing.T, p string) []byte {
	b, e := ioutil.ReadFile("fixtures/" + p)
	if e != nil {
		t.Fatal(e)
	}
	return b
}
