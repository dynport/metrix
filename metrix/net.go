package metrix

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type NetCon struct {
	Slot       int
	LocalIP    net.IP
	LocalPort  int
	RemoteIP   net.IP
	RemotePort int
	Status     string
	UUID       string
	Timeout    int
}

var socketStates = map[string]string{
	"01": "ESTABLISHED",
	"02": "SYN_SENT",
	"03": "SYN_RECV",
	"04": "FIN_WAIT1",
	"05": "FIN_WAIT2",
	"06": "TIME_WAIT",
	"07": "CLOSE",
	"08": "CLOSE_WAIT",
	"09": "LAST_ACK",
	"0A": "LISTEN",
	"0B": "CLOSING",
}

const (
	fieldNetSlot = iota
	fieldNetLocalAddress
	fieldNetRemoteAddress
	fieldNetSocketState
	fieldTxQueue
	fieldRxQueue
	fieldTr
	fieldTmWhen
	fieldRetrnSmt
	fieldUUID
	fieldTimeout
)

func ParseNetCon(raw string) (c *NetCon, err error) {
	c = &NetCon{}
	fields := strings.Fields(raw)
	for i, f := range fields {
		switch i {
		case fieldNetSlot:
			c.Slot, err = strconv.Atoi(strings.TrimSuffix(f, ":"))
			if err != nil {
				return nil, fmt.Errorf("parsing slot from %q: %s", f, err)
			}
		case fieldNetLocalAddress:
			c.LocalIP, c.LocalPort, err = parseAddress(f)
			if err != nil {
				return nil, fmt.Errorf("error parsing local address: %s", err)
			}
		case fieldNetRemoteAddress:
			c.RemoteIP, c.RemotePort, err = parseAddress(f)
			if err != nil {
				return nil, fmt.Errorf("error parsing remote address: %s", err)
			}
		case fieldNetSocketState:
			c.Status = socketStates[f]
		case fieldUUID:
			c.UUID = f
		case fieldTimeout:
			c.Timeout, err = strconv.Atoi(f)
			if err != nil {
				return nil, fmt.Errorf("error converting timeout %q to string: %s", f, err)
			}

		}
	}
	return c, nil
}

func parseAddress(s string) (ip net.IP, port int, err error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return nil, 0, fmt.Errorf("unable to split %q into 2 parts", s)
	}
	bs := make([]byte, 4)
	i, err := strconv.ParseInt(parts[0], 16, 64)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing ip %q: %s", parts[0], err)
	}
	binary.LittleEndian.PutUint32(bs, uint32(i))
	port64, err := strconv.ParseInt(parts[1], 16, 64)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing port %q: %s", parts[1], err)
	}
	return net.IP(bs), int(port64), nil
}
