package main

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

const (
	nginxActiveConnectionsPrefix = "Active connections: "
	nginxSecondLineState         = "secondLine"
	NGINX                        = "nginx"
)

func init() {
	parser.Add(NGINX, "http://127.0.0.1:8080", "Collect nginx metrics")
}

type Nginx struct {
	Address string
	Raw     []byte

	ActiveConnections int64
	Accepts           int64
	Handled           int64
	Requests          int64
	Reading           int64
	Writing           int64
	Waiting           int64
}

func (n *Nginx) Metrics() []*Metric {
	return []*Metric{
		{Key: "ActiveConnections", Value: n.ActiveConnections},
	}
}

func parseIntE(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func (nginx *Nginx) Load(raw []byte) error {
	var e error
	scanner := bufio.NewScanner(bytes.NewReader(raw))
	state := ""
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(txt, nginxActiveConnectionsPrefix) {
			nginx.ActiveConnections, e = parseIntE(strings.TrimPrefix(txt, nginxActiveConnectionsPrefix))
			if e != nil {
				return e
			}
		} else if strings.HasPrefix(txt, "server accepts") {
			state = nginxSecondLineState
		} else if state == nginxSecondLineState {
			fields := strings.Fields(txt)
			if len(fields) == 3 {
				nginx.Accepts, e = parseIntE(fields[0])
				if e != nil {
					return e
				}
				nginx.Handled, e = parseIntE(fields[1])
				if e != nil {
					return e
				}
				nginx.Requests, e = parseIntE(fields[2])
				if e != nil {
					return e
				}
			}
			state = ""
		} else if strings.HasPrefix(txt, "Reading") {
			fields := strings.Fields(txt)
			last := ""
			for _, f := range fields {
				switch last {
				case "Reading:":
					nginx.Reading, e = parseIntE(f)
					if e != nil {
						return e
					}
				case "Writing:":
					nginx.Writing, e = parseIntE(f)
					if e != nil {
						return e
					}
				case "Waiting:":
					nginx.Waiting, e = parseIntE(f)
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

func (nginx *Nginx) Collect(c *MetricsCollection) error {
	var e error
	if len(nginx.Raw) == 0 {
		nginx.Raw, e = FetchURL(nginx.Address)
		if e != nil {
			return e
		}
	}
	e = nginx.Load(nginx.Raw)
	if e != nil {
		return e
	}
	h := MetricsMap{
		"ActiveConnections": nginx.ActiveConnections,
		"Accepts":           nginx.Accepts,
		"Handled":           nginx.Handled,
		"Requests":          nginx.Requests,
		"Reading":           nginx.Reading,
		"Writing":           nginx.Writing,
		"Waiting":           nginx.Waiting,
	}
	for k, v := range h {
		c.Add(k, v)
	}
	return nil
}
