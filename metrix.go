package main

import (
	"io/ioutil"
	"os"
)

type Metrix struct {
}

func (m* Metrix) ProcRoot() (s string) {
	return os.Getenv("PROC_ROOT")
}

func (m* Metrix) ReadProcFile(path string) (r string) {
	if b, e := ioutil.ReadFile(m.ProcRoot() + "/proc/" + path); e == nil {
		r = string(b)
	}
	return
}
