package metrix

import (
	"os"
	"testing"
)

func TestParseDf(t *testing.T) {
	expect := New(t)
	f, e := os.Open("fixtures/df.txt")
	expect(e).ToBeNil()
	defer f.Close()

	disks, e := ParseDf(f)
	expect(e).ToBeNil()
	expect(disks).ToHaveLength(7)
	expect(disks[0].Filesystem).ToEqual("/dev/xvda1")
	expect(disks[0].Blocks).ToEqual(51466360)
	expect(disks[0].Used).ToEqual(32923572)
	expect(disks[0].Available).ToEqual(16346296)
	expect(disks[0].Use).ToEqual(67)
	expect(disks[0].MountedOn).ToEqual("/")
}
