package main

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseNginx(t *testing.T) {
	Convey("Nginx", t, func() {
		logger.LogLevel = INFO
		mh := &MetricHandler{}
		nginx := &Nginx{Raw: readFile("fixtures/nginx.status")}

		all, e := mh.Collect(nginx)
		sort.Sort(all)
		So(e, ShouldBeNil)

		So(len(all), ShouldEqual, 7)

		So(all[1].Key, ShouldEqual, "nginx.ActiveConnections")
		So(all[1].Value, ShouldEqual, 10)

		So(all[5].Key, ShouldEqual, "nginx.Waiting")
		So(all[5].Value, ShouldEqual, 70)

	})
}
