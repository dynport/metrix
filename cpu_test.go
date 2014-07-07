package main

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCpu(t *testing.T) {
	Convey("MetricHandler", t, func() {
		mh := &MetricHandler{}
		stats, e := mh.Collect(&Cpu{})
		So(e, ShouldBeNil)
		So(len(stats), ShouldEqual, 25)
		sort.Sort(stats)

		So(stats[0].Key, ShouldEqual, "cpu.btime")

		user := stats[23]

		So(user.Key, ShouldEqual, "cpu.user")
		So(user.Tags["total"], ShouldEqual, "")
		So(user.Tags["cpu_id"], ShouldEqual, "0")

		userCpu := stats[24]
		So(userCpu.Key, ShouldEqual, "cpu.user")
		So(userCpu.Tags["total"], ShouldEqual, "true")
		So(userCpu.Tags["cpu_id"], ShouldEqual, "")

		Convey("Load", func() {
			c := &Cpu{}
			b := mustRead(t, "proc/stat")
			e := c.Load(b)
			So(e, ShouldBeNil)
			So(len(c.Cpus), ShouldEqual, 2)
			So(c.Cpus[0].Id, ShouldEqual, "")
			So(c.Cpus[0].User, ShouldEqual, 3700)
			So(c.Cpus[1].Id, ShouldEqual, "0")
			So(c.Cpus[1].User, ShouldEqual, 3701)

			So(c.Ctxt, ShouldEqual, 957584)
			So(c.Btime, ShouldEqual, 1372481686)
			So(c.Processes, ShouldEqual, 42476)
			So(c.ProcsRunning, ShouldEqual, 1)
			So(c.ProcsBlocked, ShouldEqual, 3)

		})
	})
}
