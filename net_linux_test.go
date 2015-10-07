package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNet(t *testing.T) {
	SkipConvey("Net", t, func() {
		t.Skip("not implemented yet")
		//return
		//mh := new(MetricHandler)
		//net := &Net{}
		//stats, _ := mh.Collect(net)
		//So(len(stats), ShouldBeGreaterThan, 4)
		//So(len(stats), ShouldEqual, 21)

		//So(stats[0].Key, ShouldEqual, "net.ip.TotalPacketsReceived")
		//So(stats[0].Value, ShouldEqual, int64(162673))

		//So(stats[20].Key, ShouldEqual, "net.ip.OutOctets")
		//So(stats[20].Value, ShouldEqual, int64(667161104))
	})
}
