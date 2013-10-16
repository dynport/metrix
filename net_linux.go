package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const NET = "net"

func init() {
	parser.Add(NET, "true", "Collect network metrics")
}

var mapStatMapping = map[string]string{
	"Ip.InReceives":            "ip.TotalPacketsReceived",
	"Ip.ForwDatagrams":                         "ip.Forwarded",
	"Ip.InDiscards":        "ip.IncomingPacketsDiscarded",
	"Ip.InDelivers":        "ip.IncomingPacketsDelivered",
	"Ip.OutRequests":                 "ip.RequestsSentOut",
	"Tcp.ActiveOpens":       "tcp.ActiveConnectionsOpenings",
	"Tcp.PassiveOpens":       "tcp.PassiveConnectionsOpenings",
	"Tcp.AttemptFails":        "tcp.FailedConnectionAttempts",
	"Tcp.EstabResets":        "tcp.ConnectionResetsReceived",
	"Tcp.CurrEstab":           "tcp.ConnectionsEstablished",
	"Tcp.InSegs":                 "tcp.SegmentsReceived",
	"Tcp.OutSegs":                 "tcp.SegmentsSendOut",
	"Tcp.RetransSegs":             "tcp.SegmentsTransmitted",
	"Tcp.InErrs":            "tcp.BadSegmentsReceived",
	"Tcp.OutRsts":                       "tcp.ResetsSent",
	"Udp.InDatagrams":                  "udp.PacketsReceived",
	"Udp.NoPorts.": "udp.PacketsToUnknownPortRecived",
	"Udp.InErrors":             "udp.PacketReceiveErrors",
	"Udp.OutDatagrams":                      "udp.PacketsSent",
	"IpExt.InOctets": "ip.InOctets",
	"IpExt.OutOctets": "ip.OutOctets",
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

// Parses the self documenting format of the linux net statistics interface.
// The interface is in /proc/net and contains two consecutive lines starting
// with the same prefix like "prefix:". The first line contains the header
// and the second line the actual signed 64 bit values.
// A value < 0 means, that this statistic is not supported.
func parse2lines(headers, values string) map[string]int64 {
	keys := strings.Fields(headers)
	vals := strings.Fields(values)
	result := make(map[string]int64, len(keys))

	if len(keys) != len(vals) || len(keys) <= 1 || keys[0] != vals[0] {
		return result
	}

	// strip the ":" of "foo:" ...
	topic := keys[0][:len(keys[0])-1]
	// .. and just get the actual header entries and values
	keys = keys[1:]
	vals = vals[1:]

	for i, k := range keys {
		if v, e := strconv.ParseInt(vals[i], 10, 64); e == nil && v >= 0 {
			result[topic+"."+k] = v
		}
	}
	return result
}

func (self *Net) collect2LineFile(c *MetricsCollection, name string) (e error) {
	b, e := self.fetch(name)
	if e != nil {
		return e
	}
	raw := string(b)
	lines := strings.Split(raw, "\n")
	for i := 0; i < len(lines) - 1; i+=2 {
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
