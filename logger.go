package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", 0)

func logError(e error) {
	logger.Printf("ERROR: %q", e)
}

func debugStream() io.Writer {
	if os.Getenv("DEBUG") == "true" {
		return os.Stderr
	}
	return ioutil.Discard
}

var dbg = log.New(debugStream(), "[DEBUG] ", log.Lshortfile)
