package collectors

import (
	"io"
	"io/ioutil"
	"os"
)

func ProcRoot() string {
	return os.Getenv("PROC_ROOT")
}

func OpenProcFile(path string) (io.ReadCloser, error) {
	return os.Open(ProcRoot() + "/proc/" + path)
}

func ReadProcFile(path string) (r string) {
	f, e := OpenProcFile(path)
	if e != nil {
		return ""
	}
	defer f.Close()
	b, e := ioutil.ReadAll(f)
	if e != nil {
		return ""
	}
	return string(b)
}
