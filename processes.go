package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const PROCESSES = "processes"

func init() {
	parser.Add(PROCESSES, "true", "Collect metrics for processes")
}

var procStatsMapping = map[int]string{
	13: "Utime",
	14: "Stime",
	15: "Cutime",
	16: "Sctime",
	17: "Priority",
	18: "Nice",
	19: "NumThreads",
	20: "Itrealvalue",
	21: "Starttime",
	22: "Vsize",
	23: "RSS",
	24: "RSSlim",
	27: "Startstac",
	42: "GuestTime",
	43: "CguestTime",
}

type Processes struct {
}

func (self *Processes) Prefix() string {
	return "processes"
}

func (self *Processes) Keys() []string {
	return []string{
		"Pid",
		"Ppid",
		"Pgrp",
		"Session",
		"TtyNr",
		"Tpgid",
		"Flags",
		"Minflt",
		"Cminflt",
		"Majflt",
		"Cmajflt",
		"Utime",
		"Stime",
		"Cutime",
		"Sctime",
		"Priority",
		"Nice",
		"NumThreads",
		"Itrealvalue",
		"Starttime",
		"Vsize",
		"RSS",
		"RSSlim",
		"Startcode",
		"Endcode",
		"Startstac",
		"GuestTime",
		"CguestTime",
	}
}

var matchBrackets = regexp.MustCompile("(^\\(|\\)$)")

func NormalizeProcessName(comm string) string {
	withoutBrackes := matchBrackets.ReplaceAllString(comm, "")
	return strings.Split(withoutBrackes, "/")[0]
}

var matchNums = regexp.MustCompile("[0-9]+")

func generateProcfiles() (matches chan string, e error) {
	procmount := ProcRoot()  + "/proc"
	d, e := os.Open(procmount)
	if e != nil {
		return nil, e
	}

	matches = make(chan string, 100)
	go func(dir *os.File) {
		for e := error(nil); e == error(nil); {
			names, e := dir.Readdirnames(100)
			if e != nil {
				break
			}
			for _, name := range names {
				if matchNums.MatchString(name) {
					matches <- procmount + "/" + name + "/stat"
				}
			}
		}
		close(matches)
	}(d)
	return matches, nil
}

func (self *Processes) Collect(c *MetricsCollection) (e error) {
	matches, e := generateProcfiles()
	if e != nil {
		return
	}
	for path := range matches {
		if data, e := ioutil.ReadFile(path); e == nil {
			chunks := strings.Split(string(data), " ")
			tags := map[string]string{
				"pid":   chunks[0],
				"ppid":  chunks[3],
				"comm":  chunks[1],
				"name":  NormalizeProcessName(chunks[1]),
				"state": chunks[2],
			}
			for idx, v := range chunks {
				if i, e := strconv.ParseInt(v, 10, 64); e == nil {
					if k, ok := procStatsMapping[idx]; ok {
						c.AddWithTags(k, i, tags)
					}
				}
			}
		}
	}
	return
}
