package main

import (
	"regexp"
	"strings"
	"strconv"
)

var diskStatFields = map[int]string {
	0: "ReadsCompleted",
	1: "ReadsMerged",
	2: "SectorsRead",
	3: "MillisecondsRead",
	4: "WritesCompleted",
	5: "WritesMerged",
	6: "SectorsWritten",
	7: "MillisecondsWritten",
	8: "IosInProgress",
	9: "MillisecondsIO",
	10: "WeightedMillisecondsIO",
}

type Disk struct {
}

func (self *Disk) Keys() ([]string) {
	return []string {
		"ReadsCompleted",
		"ReadsMerged",
		"SectorsRead",
		"MillisecondsRead",
		"WritesCompleted",
		"WritesMerged",
		"SectorsWritten",
		"MillisecondsWritten",
		"IosInProgress",
		"MillisecondsIO",
		"WeightedMillisecondsIO",
	}
}

func (self *Disk) Prefix() string {
	return "disk"
}

func (self *Disk) Collect(c* MetricsCollection) (e error) {
	str := ReadProcFile("diskstats")
	re := regexp.MustCompile("\\d+\\s+\\d+ (\\w+) (\\d+.*)")
	matches := re.FindAllStringSubmatch(str, -1)
	for _, m1 := range matches {
		name := m1[1]
		tags := map[string]string { "name": name }
		if strings.HasPrefix(name, "ram") || strings.HasPrefix(name, "loop") || strings.HasPrefix(name, "sr") {
			continue
		}
		chunks := strings.Split(m1[2], " ")
		for idx, v := range chunks {
			if i, e := strconv.ParseInt(v, 10, 64); e == nil {
				if k, ok := diskStatFields[idx]; ok {
					c.AddWithTags(k, i, tags)
				}
			}
		}
	}
	return
}
