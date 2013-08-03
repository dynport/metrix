package main

import (
	"fmt"
	"os"
	"flag"
	"net/http"
	"io/ioutil"
)

func AbortWith(message string) {
	fmt.Println("ERROR:", message)
	flag.PrintDefaults()
	os.Exit(1)
}

func FetchURL(url string) (b []byte, e error) {
	var rsp *http.Response
	rsp, e = http.Get(url)
	if e != nil {
		return
	}

	b, e = ioutil.ReadAll(rsp.Body)
	return
}
