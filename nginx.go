package main

import (
	"bufio"
	"bytes"
	"strconv"
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
	ac            = "Active connections: "
	nginxSecond   = "server accepts"
	nginxAc       = "ActiveConnections"
	nginxAccepts  = "Accepts"
	nginxHandled  = "Handled"
	nginxRequests = "Requests"
	nginxReading  = "Reading"
	nginxWriting  = "Writing"
	nginxWaiting  = "Waiting"
)

var nginxKeyMapping = map[string]string{
	"Reading:": nginxReading,
	"Writing:": nginxWriting,
	"Waiting:": nginxWaiting,
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
		if strings.HasPrefix(txt, ac) {
			acs, e := strconv.ParseInt(strings.TrimPrefix(txt, ac), 10, 64)
			if e != nil {
				return e
			}
			c.Add("ActiveConnections", acs)
		} else if strings.HasPrefix(txt, nginxSecond) {
			state = "secondLine"
		} else if state == "secondLine" {
			fields := strings.Fields(txt)
			if len(fields) == 3 {
				e := c.MustAddString(nginxAccepts, fields[0])
				if e != nil {
					return e
				}
				e = c.MustAddString(nginxHandled, fields[1])
				if e != nil {
					return e
				}
				e = c.MustAddString(nginxRequests, fields[2])
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
