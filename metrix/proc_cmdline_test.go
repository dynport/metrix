package metrix

import (
	"os"
	"testing"

	"github.com/dynport/dgtk/expect"
)

func TestProcCmdline(t *testing.T) {
	expect := expect.New(t)
	f, err := os.Open("fixtures/proc_cmdline.txt")
	expect(err).ToBeNil()
	defer f.Close()

	cmd := &ProcCmdline{}
	err = cmd.Load(f)
	expect(err).ToBeNil()
	expect(cmd.Cmd).ToEqual("/opt/postgresql-9.3.5/bin/postgres")
	expect(cmd.Args).ToHaveLength(2)
	expect(cmd.Args[0]).ToEqual("-D")
	expect(cmd.Args[1]).ToEqual("/data/postgres")
}
