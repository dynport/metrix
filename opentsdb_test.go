package main

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOpenTSDB(t *testing.T) {
	Convey("Opentsdb", t, func() {
		theTime := time.Unix(11, 0)
		m := &Metric{Key: "metric", Value: int64(10)}
		So(m.OpenTSDB(theTime, "test.host"), ShouldEqual, "put metric 11 10 host=test.host")

		m = &Metric{Key: "os.cpu.User", Value: int64(10), Tags: map[string]string{"cpu_id": "1"}}
		So(m.OpenTSDB(theTime, "test.host"), ShouldEqual, "put os.cpu.User 11 10 host=test.host cpu_id=1")

		m = &Metric{Key: "os.cpu.User", Value: int64(10), Tags: map[string]string{"name": "(kworker/0:2)"}}
		So(m.OpenTSDB(theTime, "test.host"), ShouldEqual, "put os.cpu.User 11 10 host=test.host name=kworker_0_2")

	})
}
