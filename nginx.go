package main

import (
	"bufio"
	"bytes"
	"strings"
)

const NGINX = "nginx"

func init() {
	parser.Add(NGINX, "http://127.0.0.1:8080", "Collect nginx metrics")
}

type Nginx struct {
	Address string
	Raw     []byte
}

const (
	nginxActiveConnectionsPrefix = "Active connections: "
	nginxSecondLinePrefix        = "server accepts"
	nginxSecondLineState         = "secondLine"
	nginxKeyActiveConnections    = "ActiveConnections"
	nginxKeyAccepts              = "Accepts"
	nginxKeyHandled              = "Handled"
	nginxKeyRequests             = "Requests"
	nginxKeyReading              = "Reading"
	nginxKeyWriting              = "Writing"
	nginxKeyWaiting              = "Waiting"
)

var nginxKeyMapping = map[string]string{
	"Reading:": nginxKeyReading,
	"Writing:": nginxKeyWriting,
	"Waiting:": nginxKeyWaiting,
}

func (self *Nginx) Collect(c *MetricsCollection) (e error) {
	if len(self.Raw) == 0 {
		self.Raw, e = FetchURL(self.Address)
		if e != nil {
			return
		}
	}
	scanner := bufio.NewScanner(bytes.NewReader(self.Raw))
	state := ""
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(txt, nginxActiveConnectionsPrefix) {
			e = c.MustAddString(nginxKeyActiveConnections, strings.TrimPrefix(txt, nginxActiveConnectionsPrefix))
			if e != nil {
				return e
			}
		} else if strings.HasPrefix(txt, nginxSecondLinePrefix) {
			state = nginxSecondLineState
		} else if state == nginxSecondLineState {
			fields := strings.Fields(txt)
			if len(fields) == 3 {
				e := c.MustAddString(nginxKeyAccepts, fields[0])
				if e != nil {
					return e
				}
				e = c.MustAddString(nginxKeyHandled, fields[1])
				if e != nil {
					return e
				}
				e = c.MustAddString(nginxKeyRequests, fields[2])
				if e != nil {
					return e
				}
			}
			state = ""
		} else if strings.HasPrefix(txt, "Reading") {
			fields := strings.Fields(txt)
			last := ""
			for _, f := range fields {
				if key, ok := nginxKeyMapping[last]; ok {
					e := c.MustAddString(key, f)
					if e != nil {
						return e
					}
				}
				last = f
			}

		}

	}
	return scanner.Err()
}

func (self *Nginx) Prefix() string {
	return "nginx"
}

func (self *Nginx) Keys() []string {
	return []string{
		"ActiveConnections",
		"Accepts",
		"Handled",
		"Requests",
		"Reading",
		"Writing",
		"Waiting",
	}
}
