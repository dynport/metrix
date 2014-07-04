package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFiles(t *testing.T) {
	Convey("Files", t, func() {
		files := &Files{RawStatus: readFile("fixtures/lsof.txt")}

		So(files.Prefix(), ShouldEqual, "files")
		So(len(files.Keys()), ShouldNotEqual, 0)

		mh := new(MetricHandler)
		stats, _ := mh.Collect(files)
		So(len(stats) > 4, ShouldBeTrue)

		So(len(stats), ShouldEqual, 74)

		names := map[string]int{}
		for _, s := range stats {
			names[s.Tags["name"]]++
		}
		So(names["kworker/0"], ShouldEqual, 0)
		So(names["kworker"], ShouldEqual, 8)

	})
}
