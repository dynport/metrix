package collectors

import (
	"strconv"
	"strings"
)

const LOADAVG = "loadavg"

type LoadAvg struct {
}

func (l *LoadAvg) Prefix() string {
	return LOADAVG
}

func (l *LoadAvg) Collect(c *MetricsCollection) (e error) {
	s := ReadProcFile("loadavg")
	chunks := strings.Fields(s)
	if len(chunks) >= 3 {
		if v, e := parseLoadAvg(chunks[0]); e == nil {
			c.Add("Load1m", v)
		}
		if v, e := parseLoadAvg(chunks[1]); e == nil {
			c.Add("Load5m", v)
		}
		if v, e := parseLoadAvg(chunks[2]); e == nil {
			c.Add("Load15m", v)
		}
	}
	return
}

func parseLoadAvg(s string) (i int64, e error) {
	if f, e := strconv.ParseFloat(s, 64); e == nil {
		i = int64(f * 100)
	}
	return
}
