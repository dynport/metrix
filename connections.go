package main

import "strconv"
import "net"
import "fmt"
import "strings"
import "os/exec"
import "regexp"

func parseHexIp(s string) (ip net.IP) {
	a, _ := strconv.ParseInt(s, 16, 64)
	return net.ParseIP(fmt.Sprintf("%d.%d.%d.%d", (a>>0)&255, (a>>8)&255, (a>>16)&255, (a>>24)&255))
}

type Connection struct {
	Protocol      string
	ReceiveQueue  int
	SendQueue     int
	LocalAddress  Address
	RemoteAddress Address
	State         string
	Pid           int
	Name          string
}

type Address struct {
	Ip   string
	Port int
}

var socketStates = map[int]string{
	1:  "ESTABLISHED",
	2:  "SYN_SENT",
	3:  "SYN_RECV",
	4:  "FIN_WAIT1",
	5:  "FIN_WAIT2",
	6:  "TIME_WAIT",
	7:  "CLOSE",
	8:  "CLOSE_WAIT",
	9:  "LAST_ACK",
	10: "LISTEN",
	11: "CLOSING",
}

func ReadConnections() (c []*Connection, err error) {
	out, err := exec.Command("netstat", "-tupna").Output()
	if len(out) > 0 {
		return ParseConnections(string(out)), nil
	}
	return
}

func ParseStringAddress(s string) (a Address) {
	chunks := strings.Split(s, ":")
	a = Address{}
	if port, err := strconv.Atoi(chunks[len(chunks)-1]); err == nil {
		a.Port = port
	}

	if len(chunks) == 2 {
		a.Ip = chunks[0]
	} else {
		a.Ip = ""
	}
	return
}

func ParseConnections(raw string) (connections []*Connection) {
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	connections = make([]*Connection, len(lines)-2)

	splitRe := regexp.MustCompile("\\s+")

	for i, line := range lines[2:] {
		con := &Connection{}
		chunks := splitRe.Split(line, -1)
		if len(chunks) > 5 {
			con.Protocol = chunks[0]
			con.ReceiveQueue, _ = strconv.Atoi(chunks[1])
			con.SendQueue, _ = strconv.Atoi(chunks[2])
			con.LocalAddress = ParseStringAddress(chunks[3])
			con.RemoteAddress = ParseStringAddress(chunks[4])

			offset := 5
			if strings.HasPrefix(con.Protocol, "tcp") {
				con.State = chunks[5]
				offset++
			}

			if len(chunks) > offset {
				pidAndName := strings.Split(chunks[offset], "/")
				if pid, err := strconv.Atoi(pidAndName[0]); err == nil {
					con.Pid = pid
					con.Name = pidAndName[1]
				}
			}
		}
		connections[i] = con
	}
	return
}
