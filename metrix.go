package main

import (
	"io/ioutil"
	"os"
)

type Metrixd struct {
}

func (m* Metrixd) ProcRoot() (s string) {
	return os.Getenv("PROC_ROOT")
}

func (m* Metrixd) ReadProcFile(path string) (r string) {
	if b, e := ioutil.ReadFile(m.ProcRoot() + "/proc/" + path); e == nil {
		r = string(b)
	}
	return
}
