package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseRedis(t *testing.T) {
	Convey("Redis", t, func() {
		mh := &MetricHandler{}
		es := &Redis{Raw: readFile("fixtures/redis_info.txt")}

		all, _ := mh.Collect(es)
		So(len(all), ShouldBeGreaterThan, 0)

	})
}
