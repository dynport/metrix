package metrix

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	fieldLoad1 = iota
	fieldLoad5
	fieldLoad15
	fieldLoadEntities
	fieldLoadMostRecentPid
)

type LoadAvg struct {
	Min1             float64 `json:"min1,omitempty"`
	Min5             float64 `json:"min5,omitempty"`
	Min15            float64 `json:"min15,omitempty"`
	RunnableEntities int64   `json:"runnable_entities,omitempty"`
	ExistingEntities int64   `json:"existing_entities,omitempty"`
	MostRecentPid    int     `json:"most_recent_pid,omitempty"`
}

func LoadLoadAvg() (*LoadAvg, error) {
	defer benchmark("load loadavg")()
	f, e := os.Open("/proc/loadavg")
	if e != nil {
		return nil, e
	}
	defer f.Close()
	l := &LoadAvg{}
	return l, l.Load(f)
}

func (l *LoadAvg) Load(in io.Reader) error {
	b, e := ioutil.ReadAll(in)
	if e != nil {
		return e
	}
	fields := strings.Fields(string(b))
	if len(fields) != 5 {
		return fmt.Errorf("expected 5 fields, found %d", len(fields))
	}
	l.Min1, e = strconv.ParseFloat(fields[fieldLoad1], 64)
	if e != nil {
		return e
	}
	l.Min5, e = strconv.ParseFloat(fields[fieldLoad5], 64)
	if e != nil {
		return e
	}
	l.Min15, e = strconv.ParseFloat(fields[fieldLoad15], 64)
	if e != nil {
		return e
	}
	parts := strings.Split(fields[fieldLoadEntities], "/")
	if len(parts) == 2 {
		l.RunnableEntities, e = parseInt64(parts[0])
		if e != nil {
			return e
		}
		l.ExistingEntities, e = parseInt64(parts[1])
		if e != nil {
			return e
		}
	}
	l.MostRecentPid, e = strconv.Atoi(fields[fieldLoadMostRecentPid])
	return e
}
