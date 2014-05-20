package collectors

import (
	"bufio"
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

type Cpu struct {
}

func (*Cpu) Prefix() string {
	return CPU
}

func (cpu *Cpu) Collect(c *MetricsCollection) error {
	f, e := OpenProcFile("stat")
	if e != nil {
		return e
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, " ")
		if strings.HasPrefix(chunks[0], "cpu") {
			cpu.CollectCpu(c, line)
			continue
		} else if k, ok := statMapping[chunks[0]]; ok {
			if v, e := strconv.ParseInt(chunks[1], 10, 64); e == nil {
				c.Add(k, v)
			}
		}
	}
	return nil
}

func (cpu *Cpu) CollectCpu(c *MetricsCollection, line string) (metrics []*Metric) {
	chunks := strings.Fields(line)
	cpuId := strings.TrimPrefix(chunks[0], "cpu")
	tags := make(map[string]string)
	if cpuId != "" {
		tags["cpu_id"] = cpuId
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
