package collectors

import (
	"log"
	"os"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParseNetstat(t *testing.T) {
	Convey("ParseNetstat", t, func() {
		f, e := os.Open("fixtures/proc/net/snmp")
		if e != nil {
			log.Fatal(e.Error())
		}
		defer f.Close()
		stats := parseNetstats(f)
		So(len(stats), ShouldEqual, 75)
		cnt := 0
		for _, m := range stats {
			switch m.Key {
			case "Ip.Forwarding":
				So(m.Value, ShouldEqual, 2)
				cnt++
			case "Icmp.InMsgs":
				So(m.Value, ShouldEqual, 9)
				cnt++
			}
		}
		So(cnt, ShouldEqual, 2)
	})
}
