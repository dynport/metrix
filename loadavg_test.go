package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadAvg(t *testing.T) {
	Convey("TestLoadAvg", t, func() {
		mh := &MetricHandler{}
		stats, _ := mh.Collect(&LoadAvg{})
		So(len(stats), ShouldEqual, 3)
		So(stats[0].Value, ShouldEqual, 3)
		So(stats[1].Value, ShouldEqual, 7)
		So(stats[2].Value, ShouldEqual, 8)
	})
}
