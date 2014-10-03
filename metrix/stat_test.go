package metrix

import (
	"os"
	"testing"
)

func TestStat(t *testing.T) {
	expect := New(t)
	f, e := os.Open("fixtures/stat.txt")
	expect(e).ToBeNil()
	defer f.Close()

	s := &Stat{}
	e = s.Load(f)
	expect(e).ToBeNil()
	expect(s.Cpu).ToNotBeNil()
	expect(s.Cpu.User).ToEqual(13586)
	expect(s.Cpu.System).ToEqual(32308)
	expect(s.Cpu.IOWait).ToEqual(762)
	expect(len(s.Cpus)).ToEqual(4)
	expect(s.Cpus[0].User).ToEqual(4488)
	expect(s.Cpus[0].System).ToEqual(23009)
	expect(s.BootTime).ToEqual(1412216124)
	expect(s.ContextSwitches).ToEqual(4087538)
	expect(s.Processes).ToEqual(20491)
}
