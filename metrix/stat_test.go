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
		So(s.Cpu.User, ShouldEqual, 13586)
		So(s.Cpu.System, ShouldEqual, 32308)
		So(s.Cpu.IOWait, ShouldEqual, 762)
		So(len(s.Cpus), ShouldEqual, 4)
		So(s.Cpus[0].User, ShouldEqual, 4488)
		So(s.Cpus[0].System, ShouldEqual, 23009)
		So(s.BootTime, ShouldEqual, 1412216124)
		So(s.ContextSwitches, ShouldEqual, 4087538)
		So(s.Processes, ShouldEqual, 20491)
	})
}
