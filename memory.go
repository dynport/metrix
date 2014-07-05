package main

import (
	"regexp"
	"strconv"
)

const MEMORY = "memory"

func init() {
	parser.Add(MEMORY, "true", "Collect memory metrics")
}

type Memory struct {
}

func (m *Memory) Prefix() string {
	return "memory"
}

func (memory *Memory) Collect(c *MetricsCollection) (e error) {
	s := ReadProcFile("meminfo")
	re := regexp.MustCompile("(.*?)\\:\\s+(\\d+)")
	matches := re.FindAllStringSubmatch(s, -1)
	for _, a := range matches {
		k := a[1]
		v, e := strconv.ParseInt(a[2], 10, 64)
		if e == nil {
			c.Add(k, v)
		}
	}
	return
}
