package metrix

import (
	"os"
	"testing"

	"github.com/dynport/dgtk/expect"
)

func TestParseProcStat(t *testing.T) {
	expect := expect.New(t)
	f, e := os.Open("fixtures/proc_stat.txt")
	expect(e).ToBeNil()
	defer f.Close()
	p := &ProcStat{}
	expect(p).ToNotBeNil()
	expect(p.Load(f)).ToBeNil()
	expect(p.Pid).ToEqual(1153)
	expect(p.Comm).ToEqual("(cron)")
	expect(p.State).ToEqual("S")
	expect(p.Ppid).ToEqual(1)
	expect(p.Pgrp).ToEqual(1153)
	expect(p.Utime).ToEqual(17)
	expect(p.Stime).ToEqual(115)
	expect(p.Cutime).ToEqual(3549)
	expect(p.Cstime).ToEqual(2293)
	expect(p.NumThreads).ToEqual(1)
	expect(p.RSS).ToEqual(262)
	expect(p.VSize).ToEqual(24223744)
	expect(p.RSSlim).ToEqual(0)
	expect(p.StatStartTime).ToEqual(1301)
}

func TestNumeric(t *testing.T) {
	expect := expect.New(t)
	expect(numeric("1234")).ToEqual(true)
	expect(numeric("1")).ToEqual(true)
	expect(numeric("")).ToEqual(false)
	expect(numeric("12a")).ToEqual(false)
}
