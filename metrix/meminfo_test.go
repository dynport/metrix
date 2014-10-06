package metrix

import (
	"os"
	"testing"

	"github.com/dynport/dgtk/expect"
)

func TestMeminfo(t *testing.T) {
	expect := expect.New(t)
	f, e := os.Open("fixtures/meminfo.txt")
	expect(e).ToBeNil()
	defer f.Close()

	m := &Meminfo{}
	e = m.Load(f)
	expect(e).ToBeNil()
	expect(m.MemTotal).ToEqual(4041052)
	expect(m.VmallocTotal).ToEqual(34359738367)
	expect(m.InactiveFile).ToEqual(169752)
	expect(m.DirectMap1G).ToEqual(3145728)
}
