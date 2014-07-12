package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse2lines(t *testing.T) {
	Convey("Parse2Lines", t, func() {
		header := "Ip: Forwarding DefaultTTL InReceives InHdrErrors InAddrErrors ForwDatagrams InUnknownProtos InDiscards InDelivers OutRequests OutDiscards OutNoRoutes ReasmTimeout ReasmReqds ReasmOKs ReasmFails FragOKs FragFails FragCreates"
		value := "Ip: 2 64 1594741 0 0 0 0 0 1572047 2228197 5479 1154 0 390 78 0 0 0 0"
		want := map[string]int64{
			"Ip.Forwarding": 2,
			"Ip.DefaultTTL": 64,

			"Ip.InReceives":   1594741,
			"Ip.InHdrErrors":  0,
			"Ip.InAddrErrors": 0,

			"Ip.ForwDatagrams": 0,

			"Ip.InUnknownProtos": 0,
			"Ip.InDiscards":      0,
			"Ip.InDelivers":      1572047,

			"Ip.OutRequests": 2228197,
			"Ip.OutDiscards": 5479,
			"Ip.OutNoRoutes": 1154,

			"Ip.ReasmTimeout": 0,
			"Ip.ReasmReqds":   390,
			"Ip.ReasmOKs":     78,
			"Ip.ReasmFails":   0,

			"Ip.FragOKs":     0,
			"Ip.FragFails":   0,
			"Ip.FragCreates": 0,
		}

		got := parse2lines(header, value)
		So(got, ShouldEqual, want)
	})
}

func TestNet(t *testing.T) {
	SkipConvey("Net", t, func() {
		t.Skip("not implemented yet")
		return

		mh := new(MetricHandler)
		net := &Net{}
		stats, _ := mh.Collect(net)
		So(len(stats), ShouldBeGreaterThan, 4)
		So(len(stats), ShouldEqual, 21)

		So(stats[0].Key, ShouldEqual, "net.ip.TotalPacketsReceived")
		So(stats[0].Value, ShouldEqual, int64(162673))

		So(stats[20].Key, ShouldEqual, "net.ip.OutOctets")
		So(stats[20].Value, ShouldEqual, int64(667161104))
	})
}
