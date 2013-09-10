package main

import (
	"strconv"
	"strings"
)

const LOADAVG = "loadavg"

func init() {
	parser.Add(LOADAVG, "true", "Collect loadvg metrics")
}

type LoadAvg struct {
}

func (l *LoadAvg) Keys() []string {
	return []string{
		"Load1m",
		"Load5m",
		"Load15m",
	}
}

func (l *LoadAvg) Prefix() string {
	return "load"
}

func (l *LoadAvg) Collect(c *MetricsCollection) (e error) {
	s := ReadProcFile("loadavg")
	chunks := strings.Split(s, " ")
	if len(chunks) >= 3 {
		if v, e := l.ParseLoadAvg(chunks[0]); e == nil {
			c.Add("Load1m", v)
		}
		if v, e := l.ParseLoadAvg(chunks[1]); e == nil {
			c.Add("Load5m", v)
		}
		if v, e := l.ParseLoadAvg(chunks[2]); e == nil {
			c.Add("Load15m", v)
		}
	}
	return
}

func (l *LoadAvg) ParseLoadAvg(s string) (i int64, e error) {
	if f, e := strconv.ParseFloat(s, 64); e == nil {
		i = int64(f * 100)
	}
	return
}
