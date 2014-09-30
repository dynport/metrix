package metrix

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseDf(t *testing.T) {
	Convey("ParseDf", t, func() {
		f, e := os.Open("fixtures/df.txt")
		So(e, ShouldBeNil)
		defer f.Close()

		disks, e := ParseDf(f)
		So(e, ShouldBeNil)
		So(len(disks), ShouldEqual, 7)
		So(disks[0].Filesystem, ShouldEqual, "/dev/xvda1")
		So(disks[0].Blocks, ShouldEqual, 51466360)
		So(disks[0].Used, ShouldEqual, 32923572)
		So(disks[0].Available, ShouldEqual, 16346296)
		So(disks[0].Use, ShouldEqual, 67)
		So(disks[0].MountedOn, ShouldEqual, "/")
	})
}
