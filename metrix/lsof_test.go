package metrix

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOpenFiles(t *testing.T) {
	Convey("Load", t, func() {
		f, e := os.Open("fixtures/lsof.txt")
		So(e, ShouldBeNil)
		defer f.Close()

		files := &OpenFiles{}
		e = files.Load(f)
		So(e, ShouldBeNil)
		So(files.Files, ShouldNotBeNil)
		So(len(files.Files), ShouldEqual, 265)

		selectFile := func(pid string) *File {
			for _, f := range files.Files {
				if f.ProcessId == pid {
					return f
				}
			}
			return nil
		}

		file := selectFile("1")
		So(file, ShouldNotBeNil)
		So(file.FileInodeNumber, ShouldEqual, "6964")
		So(file.FileName, ShouldEqual, "/dev/ptmx")
	})
}
