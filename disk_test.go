package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDiskStat(t *testing.T) {
	Convey("DiskStat", t, func() {
		mh := &MetricHandler{}
		disk := &Disk{}
		stats, _ := mh.Collect(disk)

		So(len(stats), ShouldEqual, 11)
		So(stats[0].Key, ShouldEqual, "disk.ReadsCompleted")
		So(stats[0].Value, ShouldEqual, int64(9748))
		So(stats[0].Tags["name"], ShouldEqual, "sda")
	})
}
