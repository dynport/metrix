package collectors

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
	Convey("TestMemory", t, func() {
		m := &Memory{}
		mh := &MetricHandler{}
		stats, _ := mh.Collect(m)

		So(len(stats), ShouldEqual, 42)
		passed := 0
		for _, s := range stats {
			switch s.Key {
			case "memory.MemTotal":
				So(s.Value, ShouldEqual, 502976)
				passed++
			case "memory.Committed_AS":
				So(s.Value, ShouldEqual, 350856)
				passed++
			case "memory.ActiveAnon":
				So(s.Value, ShouldEqual, 16568)
				passed++
			}
		}
		So(passed, ShouldEqual, 3)
	})
}
