package metrix

import (
	"os"
	"testing"

	"github.com/dynport/dgtk/expect"
)

func TestOpenFiles(t *testing.T) {
	expect := expect.New(t)
	f, e := os.Open("fixtures/lsof.txt")
	expect(e).ToBeNil()
	defer f.Close()

	files := &OpenFiles{}
	expect(files.Load(f)).ToBeNil()
	expect(files.Files).ToNotBeNil()
	expect(files.Files).ToHaveLength(265)

	selectFile := func(pid string) *File {
		for _, f := range files.Files {
			if f.ProcessId == pid {
				return f
			}
		}
		return nil
	}

	file := selectFile("1")
	expect(file).ToNotBeNil()
	expect(file.FileInodeNumber).ToEqual("6964")
	expect(file.FileName).ToEqual("/dev/ptmx")
}
