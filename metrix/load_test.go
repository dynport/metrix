package metrix

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadAVG(t *testing.T) {
	Convey("parse", t, func() {
		f, e := os.Open("fixtures/loadavg.txt")
		So(e, ShouldBeNil)
		defer f.Close()

		l := &LoadAvg{}
		e = l.Load(f)
		So(e, ShouldBeNil)
		So(l.Min1, ShouldEqual, 0.00)
		So(l.Min5, ShouldEqual, 0.01)
		So(l.Min15, ShouldEqual, 0.05)
		So(l.RunnableEntities, ShouldEqual, 1)
		So(l.ExistingEntities, ShouldEqual, 258)
		So(l.MostRecentPid, ShouldEqual, 8518)
	})
}
