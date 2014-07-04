package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProcesses(t *testing.T) {
	Convey("Processes", t, func() {
		mh := new(MetricHandler)
		col := &Processes{}
		stats, _ := mh.Collect(col)
		So(len(stats), ShouldBeGreaterThan, 0)
		So(len(stats), ShouldEqual, 980)
		for _, stat := range stats {
			if stat.Tags["name"] == "init" && stat.Key == "processes.Utime" {
				So(stat.Key, ShouldEqual, "processes.Utime")
				So(stat.Value, ShouldEqual, 25)
				So(stat.Tags["comm"], ShouldEqual, "(init)")
				So(stat.Tags["name"], ShouldEqual, "init")
				So(stat.Tags["state"], ShouldEqual, "S")
				So(stat.Tags["pid"], ShouldEqual, "1")
			}
		}

	})
}

func TestNormlizeProcessName(t *testing.T) {
	Convey("NormalizeProcessName", t, func() {
		So(NormalizeProcessName("(int)"), ShouldEqual, "int")
		So(NormalizeProcessName("(kworker/2:1H)"), ShouldEqual, "kworker")

	})
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
