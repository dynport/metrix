package collectors

import (
	"log"
	"os"
)

var debugger *log.Logger

func init() {
	device := os.Stdout
	if os.Getenv("DEBUG") != "true" {
		var e error
		flag := os.O_RDWR | os.O_CREATE | os.O_APPEND
		device, e = os.OpenFile(os.DevNull, flag, 0700)
		if e != nil {
			panic(e.Error())
		}
	}
	debugger = log.New(device, "", log.Lshortfile)
}
