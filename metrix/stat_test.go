package metrix

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStat(t *testing.T) {
	Convey("parse", t, func() {
		f, e := os.Open("fixtures/stat.txt")
		So(e, ShouldBeNil)
		defer f.Close()

		s := &Stat{}
		e = s.Load(f)
		So(e, ShouldBeNil)
		So(s.Cpu, ShouldNotBeNil)
		So(s.Cpu.User, ShouldEqual, 6703)
		So(s.Cpu.System, ShouldEqual, 7497)
		So(s.Cpu.IOWait, ShouldEqual, 311)
		So(len(s.Cpus), ShouldEqual, 4)
		So(s.Cpus[0].User, ShouldEqual, 2191)
		So(s.Cpus[0].System, ShouldEqual, 4023)
		So(s.BootTime, ShouldEqual, 1412045635)
		So(s.ContextSwitches, ShouldEqual, 1058203)
		So(s.Processes, ShouldEqual, 8497)
	})
}
