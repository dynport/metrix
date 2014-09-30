package metrix

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMeminfo(t *testing.T) {
	Convey("parse", t, func() {
		f, e := os.Open("fixtures/meminfo.txt")
		So(e, ShouldBeNil)
		defer f.Close()

		m := &Meminfo{}
		e = m.Load(f)
		So(e, ShouldBeNil)
		So(m.MemTotal, ShouldEqual, 4041052)
		So(m.VmallocTotal, ShouldEqual, 34359738367)
		So(m.InactiveFile, ShouldEqual, 169752)
		So(m.DirectMap1G, ShouldEqual, 3145728)
	})
}
