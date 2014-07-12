package main

import (
	"io/ioutil"
	"strings"
)

const NET = "net"

func init() {
	parser.Add(NET, "true", "Collect network metrics")
}

var mapStatMapping = map[string]string{
	"Ip.InReceives":    "ip.TotalPacketsReceived",
	"Ip.ForwDatagrams": "ip.Forwarded",
	"Ip.InDiscards":    "ip.IncomingPacketsDiscarded",
	"Ip.InDelivers":    "ip.IncomingPacketsDelivered",
	"Ip.OutRequests":   "ip.RequestsSentOut",
	"Tcp.ActiveOpens":  "tcp.ActiveConnectionsOpenings",
	"Tcp.PassiveOpens": "tcp.PassiveConnectionsOpenings",
	"Tcp.AttemptFails": "tcp.FailedConnectionAttempts",
	"Tcp.EstabResets":  "tcp.ConnectionResetsReceived",
	"Tcp.CurrEstab":    "tcp.ConnectionsEstablished",
	"Tcp.InSegs":       "tcp.SegmentsReceived",
	"Tcp.OutSegs":      "tcp.SegmentsSendOut",
	"Tcp.RetransSegs":  "tcp.SegmentsTransmitted",
	"Tcp.InErrs":       "tcp.BadSegmentsReceived",
	"Tcp.OutRsts":      "tcp.ResetsSent",
	"Udp.InDatagrams":  "udp.PacketsReceived",
	"Udp.NoPorts.":     "udp.PacketsToUnknownPortRecived",
	"Udp.InErrors":     "udp.PacketReceiveErrors",
	"Udp.OutDatagrams": "udp.PacketsSent",
	"IpExt.InOctets":   "ip.InOctets",
	"IpExt.OutOctets":  "ip.OutOctets",
}

type Net struct {
	RawStatus []byte
}

func (self *Net) fetch(key string) (b []byte, e error) {
	return ioutil.ReadFile(ProcRoot() + "/proc/net/" + key)
}

func (self *Net) Prefix() string {
	return "net"
}

func (self *Net) collect2LineFile(c *MetricsCollection, name string) (e error) {
	b, e := self.fetch(name)
	if e != nil {
		return e
	}
	raw := string(b)
	lines := strings.Split(raw, "\n")
	for i := 0; i < len(lines)-1; i += 2 {
		for k, v := range parse2lines(lines[i], lines[i+1]) {
			if mapped, ok := mapStatMapping[k]; ok {
				c.Add(mapped, v)
			}
		}
	}
	return nil
}

func (self *Net) Collect(c *MetricsCollection) error {
	for _, name := range []string{"snmp", "netstat"} {
		if e := self.collect2LineFile(c, name); e != nil {
			return e
		}
	}
	return nil
}
