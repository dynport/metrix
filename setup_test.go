package main

import (
	"github.com/remogatto/prettytest"
	"testing"
	"os"
	"io/ioutil"
)

type testSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	dir, _ := os.Getwd()
	os.Setenv("PROC_ROOT", dir + "/fixtures")
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testSuite),
	)
}

func readFile(path string) []byte {
	b, e := ioutil.ReadFile(path)
	if e != nil {
		panic(e.Error())
	}
	return b
}
