// +build !linux

package main

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
)

const NET = "net"

func init() {
	parser.Add(NET, "true", "Collect network metrics")
}

var mapStatMapping = map[string]string{
	"total packets received":            "ip.TotalPacketsReceived",
	"forwarded":                         "ip.Forwarded",
	"incoming packets discarded":        "ip.IncomingPacketsDiscarded",
	"incoming packets delivered":        "ip.IncomingPacketsDelivered",
	"requests sent out":                 "ip.RequestsSentOut",
	"active connections openings":       "tcp.ActiveConnectionsOpenings",
	"passive connection openings":       "tcp.PassiveConnectionsOpenings",
	"failed connection attempts":        "tcp.FailedConnectionAttempts",
	"connection resets received":        "tcp.ConnectionResetsReceived",
	"connections established":           "tcp.ConnectionsEstablished",
	"segments received":                 "tcp.SegmentsReceived",
	"segments send out":                 "tcp.SegmentsSendOut",
	"segments retransmited":             "tcp.SegmentsTransmitted",
	"bad segments received.":            "tcp.BadSegmentsReceived",
	"resets sent":                       "tcp.ResetsSent",
	"packets received":                  "udp.PacketsReceived",
	"packets to unknown port received.": "udp.PacketsToUnknownPortRecived",
	"packet receive errors":             "udp.PacketReceiveErrors",
	"packets sent":                      "udp.PacketsSent",
}

type Net struct {
	RawStatus []byte
}

func (self *Net) fetch() (b []byte, e error) {
	if len(self.RawStatus) == 0 {
		self.RawStatus, _ = exec.Command("netstat", "-s").Output()
		if len(self.RawStatus) == 0 {
			e = errors.New("netstat returned empty output")
			return
		}
	}
	b = self.RawStatus
	return
}

func (self *Net) Prefix() string {
	return "net"
}

func (self *Net) Collect(c *MetricsCollection) (e error) {
	s, e := self.fetch()
	if e != nil {
		return
	}
	raw := string(s)
	re := regexp.MustCompile("(\\d+) ([\\w\\ \\.]+)")
	for _, v := range re.FindAllStringSubmatch(raw, -1) {
		if value, e := strconv.ParseInt(v[1], 10, 64); e == nil {
			if k, ok := mapStatMapping[v[2]]; ok {
				c.Add(k, value)
			}
		}
	}

	re = regexp.MustCompile("(\\w+): (\\d+)")
	for _, v := range re.FindAllStringSubmatch(raw, -1) {
		if value, e := strconv.ParseInt(v[2], 10, 64); e == nil {
			switch v[1] {
			case "InOctets":
				c.Add("ip.InOctets", value)
			case "OutOctets":
				c.Add("ip.OutOctets", value)
			}
		}
	}
	return
}

func (self *Net) Keys() []string {
	return []string{
		"ip.TotalPacketsReceived",
		"ip.Forwarded",
		"ip.IncomingPacketsDiscarded",
		"ip.IncomingPacketsDelivered",
		"ip.RequestsSentOut",
		"tcp.ActiveConnectionsOpenings",
		"tcp.PassiveConnectionsOpenings",
		"tcp.FailedConnectionAttempts",
		"tcp.ConnectionResetsReceived",
		"tcp.ConnectionsEstablished",
		"tcp.SegmentsReceived",
		"tcp.SegmentsSendOut",
		"tcp.SegmentsTransmitted",
		"tcp.BadSegmentsReceived",
		"tcp.ResetsSent",
		"udp.PacketsReceived",
		"udp.PacketsToUnknownPortRecived",
		"udp.PacketReceiveErrors",
		"udp.PacketsSent",
		"ip.InOctets",
		"ip.OutOctets",
	}
}
