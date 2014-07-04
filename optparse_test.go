package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOptParse(t *testing.T) {
	Convey("OptParse", t, func() {
		parser := &OptParser{}
		parser.Add("redis", "127.0.0.1:6379", "")
		parser.AddKey("postgres", "")

		parser.ProcessAll([]string{"--redis"})
		v := parser.Get("redis")
		So(v, ShouldEqual, "127.0.0.1:6379")

		parser.ProcessAll([]string{"--redis", "1.2.3.4:1234"})
		v = parser.Get("redis")
		So(v, ShouldEqual, "1.2.3.4:1234")

		parser.ProcessAll([]string{"--redis=1.2.3.4:1234"})
		v = parser.Get("redis")
		So(v, ShouldEqual, "1.2.3.4:1234")

		parser.ProcessAll([]string{"-redis", "1.2.3.4:1234", "-postgres", "1.2.3.4"})
		v = parser.Get("postgres")
		So(v, ShouldEqual, "1.2.3.4")

	})
}
