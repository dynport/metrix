package metrix

import (
	"os"
	"testing"

	"github.com/dynport/dgtk/expect"
)

func TestLoadAVG(t *testing.T) {
	expect := expect.New(t)
	f, e := os.Open("fixtures/loadavg.txt")
	expect(e).ToBeNil()
	defer f.Close()
	l := &LoadAvg{}
	expect(l.Load(f)).ToBeNil()
	expect(l.Min1).ToEqual(0.00)
	expect(l.Min5).ToEqual(0.01)
	expect(l.Min15).ToEqual(0.05)
	expect(l.RunnableEntities).ToEqual(1)
	expect(l.ExistingEntities).ToEqual(258)
	expect(l.MostRecentPid).ToEqual(8518)
}
