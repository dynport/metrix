package main

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRiakStatus(t *testing.T) {
	Convey("RiakStatus", t, func() {
		m := &Riak{}
		raw := readFile("fixtures/riak.json")
		status, e := m.ParseRiakStatus(raw)
		if e != nil {
			t.Fatal(e.Error())
		}
		So(status.VNodeGets, ShouldEqual, int64(241))
		So(status.NodePutFsmActive, ShouldEqual, int64(0))
		So(status.CpuAvg1, ShouldEqual, int64(274))

		So(len(status.ConnectedNodes), ShouldEqual, 4)
		So(status.ConnectedNodes[0], ShouldEqual, "riak@192.168.0.16")

		So(len(status.RingMembers), ShouldEqual, 5)
		So(status.RingMembers[0], ShouldEqual, "riak@192.168.0.16")

	})
}

func TestSerializeMetric(t *testing.T) {
	Convey("SerializeMetric", t, func() {
		m := &Metric{}
		b, e := json.Marshal(m)
		So(e, ShouldBeNil)
		So(string(b), ShouldNotContainSubstring, "Tags")

		m.Tags = map[string]string{"a": "b"}
		b, e = json.Marshal(m)
		So(e, ShouldBeNil)
		So(string(b), ShouldContainSubstring, "Tags")
	})
}
