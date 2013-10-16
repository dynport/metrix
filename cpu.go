package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var statMapping = map[string]string{
	"ctxt":          "Ctxt",
	"btime":         "Btime",
	"processes":     "Processes",
	"procs_running": "ProcsRunning",
	"procs_blocked": "ProcsBlocked",
}

var cpuLineMapping = map[int]string{
	1: "User",
	2: "Nice",
	3: "System",
	4: "Idle",
	5: "IOWait",
	6: "IRC",
	7: "SoftIRQ",
}

const CPU = "cpu"

func init() {
	parser.Add(CPU, "true", "Collect cpu metrics")
}

type Cpu struct {
}

func (*Cpu) Prefix() string {
	return "cpu"
}

func (*Cpu) Keys() []string {
	return []string{
		"Ctxt",
		"Btime",
		"Processes",
		"ProcsRunning",
		"ProcsBlocked",
		"User",
		"Nice",
		"System",
		"Idle",
		"IOWait",
		"IRC",
		"SoftIRQ",
	}
}

func ProcRoot() string {
	return os.Getenv("PROC_ROOT")
}

func ReadProcFile(path string) (r string) {
	if b, e := ioutil.ReadFile(ProcRoot() + "/proc/" + path); e == nil {
		r = string(b)
	}
	return
}

func (cpu *Cpu) Collect(c *MetricsCollection) (e error) {
	str := ReadProcFile("stat")
	for _, line := range strings.Split(str, "\n") {
		chunks := strings.Split(line, " ")
		if strings.HasPrefix(chunks[0], "cpu") {
			cpu.CollectCpu(c, line)
			continue
		}
		if k, ok := statMapping[chunks[0]]; ok {
			if v, e := strconv.ParseInt(chunks[1], 10, 64); e == nil {
				c.Add(k, v)
			}
		}
	}
	return
}

func (cpu *Cpu) CollectCpu(c *MetricsCollection, line string) (metrics []*Metric) {
	chunks := strings.Fields(line)

	re := regexp.MustCompile("^cpu(\\d+)")
	mets := re.FindStringSubmatch(chunks[0])
	tags := make(map[string]string)
	if len(mets) == 2 {
		tags["cpu_id"] = mets[1]
	} else {
		tags["total"] = "true"
	}
	for i, v := range chunks {
		if k, ok := cpuLineMapping[i]; ok {
			if i, e := strconv.ParseInt(v, 10, 64); e == nil {
				c.AddWithTags(k, i, tags)
			}
		}
	}
	return
}
