package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCpu(t *testing.T) {
	Convey("MetricHandler", t, func() {
		mh := &MetricHandler{}
		stats, _ := mh.Collect(&Cpu{})
		So(len(stats), ShouldEqual, 19)
		So(stats[0].Key, ShouldEqual, "cpu.User")
		So(stats[0].Key, ShouldEqual, "cpu.User")
		So(stats[0].Tags["total"], ShouldEqual, "true")
		So(stats[0].Value, ShouldEqual, 3700)

		So(stats[14].Key, ShouldEqual, "cpu.Ctxt")
		So(stats[14].Value, ShouldEqual, 957584)
	})
}
