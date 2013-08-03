package main

import (
	"regexp"
	"strconv"
	"errors"
)

type Nginx struct {
	Address string
	Raw []byte
}

var nginxRegexp = regexp.MustCompile("Active connections: ([\\d]+)\\s+\nserver.*?\n\\s+([\\d]+)\\s+([\\d]+)\\s+([\\d]+)\\s+\nReading: ([\\d]+)\\s+Writing: ([\\d]+)\\s+Waiting: ([\\d]+)")

var nginxMapping = map[string]int {
	"ActiveConnections": 1,
	"Accepts": 2,
	"Handled": 3,
	"Requests": 4,
	"Reading": 5,
	"Writing": 6,
	"Waiting": 7,
}

func (self *Nginx) Collect(c* MetricsCollection) (e error) {
	if len(self.Raw) == 0 {
		self.Raw, e = FetchURL(self.Address)
		if e != nil {
			return
		}
	}
	s := string(self.Raw)
	m := nginxRegexp.FindStringSubmatch(s)
	if len(m) == 0 {
		return errors.New("could not parse nginx status")
	}
	for key, idx := range nginxMapping {
		if idx < len(m) {
			i, e := strconv.ParseInt(m[idx], 10, 64)
			if e == nil {
				c.Add(key, i)
			}

		}
	}
	return
}

func (self *Nginx) Prefix() string {
	return "nginx"
}

func (self *Nginx) Keys() []string {
	return []string {
		"ActiveConnections",
		"Accepts",
		"Handled",
		"Requests",
		"Reading",
		"Writing",
		"Waiting",
	}
}
