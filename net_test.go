// +build !linux

package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNet(t *testing.T) {
	Convey("Net", t, func() {
		mh := new(MetricHandler)
		net := &Net{RawStatus: readFile("fixtures/netstat.txt")}
		stats, _ := mh.Collect(net)
		So(len(stats) > 4, ShouldBeTrue)
		So(len(stats), ShouldEqual, 21)
		So(stats[0].Key, ShouldEqual, "net.ip.TotalPacketsReceived")
		So(stats[0].Value, ShouldEqual, int64(162673))
		So(stats[20].Key, ShouldEqual, "net.ip.OutOctets")
		So(stats[20].Value, ShouldEqual, int64(667161104))
	})
}
