package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDf(t *testing.T) {
	Convey("Df", t, func() {
		mh := new(MetricHandler)
		col := &Df{
			RawSpace: readFile("fixtures/df.txt"),
			RawInode: readFile("fixtures/df_inode.txt"),
		}
		stats, e := mh.Collect(col)
		So(e, ShouldBeNil)

		So(len(stats), ShouldEqual, 8)

		So(stats[0].Key, ShouldEqual, "df.space.Total")
		So(stats[0].Value, ShouldEqual, int64(20511356))
		So(stats[0].Tags["file_system"], ShouldEqual, "/dev/sda")
		So(stats[0].Tags["mounted_on"], ShouldEqual, "/")

		So(stats[4].Key, ShouldEqual, "df.inode.Total")
		So(stats[4].Value, ShouldEqual, int64(1310720))
		So(stats[4].Tags["file_system"], ShouldEqual, "/dev/sda")
		So(stats[4].Tags["mounted_on"], ShouldEqual, "/")
	})
}
