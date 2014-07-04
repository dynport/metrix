package main

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGraphite(t *testing.T) {
	Convey("TestGraphite", t, func() {
		theTime := time.Unix(11, 0)
		m := &Metric{Key: "metric", Value: 10}
		So(m.Graphite(theTime, "test.host"), ShouldEqual, "metrix.hosts.test.host.metric 10 11")

		m = &Metric{Key: "os.cpu.User", Value: 10, Tags: map[string]string{"cpu_id": "1"}}
		So(m.Graphite(theTime, "test.host"), ShouldEqual, "metrix.hosts.test.host.os.cpu1.User 10 11")

		m = &Metric{Key: "disk.WeightedMillisecondsIO", Value: int64(10), Tags: map[string]string{"name": "sda"}}
		So(m.Graphite(theTime, "test.host"), ShouldEqual, "metrix.hosts.test.host.disk.sda.WeightedMillisecondsIO 10 11")

		m = &Metric{Key: "df.inodes.Use", Value: 10, Tags: map[string]string{"file_system": "/dev/sda"}}
		So(m.Graphite(theTime, "test.host"), ShouldEqual, "metrix.hosts.test.host.df.dev.sda.inodes.Use 10 11")

		m = &Metric{Key: "processes.Utime", Value: 10, Tags: map[string]string{"name": "init", "pid": "10"}}
		So(m.Graphite(theTime, "test.host"), ShouldEqual, "metrix.hosts.test.host.processes.init.10.Utime 10 11")
	})
}
