package main

import (
	"io/ioutil"
	"os"
)

func init() {
	dir, _ := os.Getwd()
	os.Setenv("PROC_ROOT", dir+"/fixtures")
}

func readFile(path string) []byte {
	b, e := ioutil.ReadFile(path)
	if e != nil {
		panic(e.Error())
	}
	return b
}
