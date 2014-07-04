package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseNginx(t *testing.T) {
	Convey("Nginx", t, func() {
		logger.LogLevel = INFO
		mh := &MetricHandler{}
		nginx := &Nginx{Raw: readFile("fixtures/nginx.status")}

		all, e := mh.Collect(nginx)
		So(e, ShouldBeNil)

		for _, m := range all {
			t.Logf("%s: %v", m.Key, m.Value)
		}

		So(len(all), ShouldEqual, 7)

		So(all[0].Key, ShouldEqual, "nginx.ActiveConnections")
		So(all[0].Value, ShouldEqual, 10)

		So(all[6].Key, ShouldEqual, "nginx.Waiting")
		So(all[6].Value, ShouldEqual, 70)

	})
}
