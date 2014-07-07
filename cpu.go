package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
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

const (
	cpuFieldUser = iota + 1
	cpuFieldNice
	cpuFieldSystem
	cpuFieldIdle
	cpuFieldIOWait
	cpuFieldIrq
	cpuFieldSoftIrq
	cpuFieldSteal
	cpuFieldGuest
	cpuFieldGuestNice
)

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
	Cpus         []*CpuStat
	Ctxt         int64
	Btime        int64
	Processes    int64
	ProcsRunning int64
	ProcsBlocked int64
}

func (c *Cpu) Load(b []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	c.Cpus = []*CpuStat{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			// continue
		}
		key := parts[0]
		if strings.HasPrefix(key, "cpu") {
			s := &CpuStat{Id: strings.TrimPrefix(key, "cpu")}
			for i, v := range parts[1:] {
				value, e := strconv.ParseInt(v, 10, 64)
				if e != nil {
					return e
				}
				switch i + 1 {
				case cpuFieldUser:
					s.User = value
				case cpuFieldNice:
					s.Nice = value
				case cpuFieldSystem:
					s.System = value
				case cpuFieldIdle:
					s.Idle = value
				case cpuFieldIOWait:
					s.IOWait = value
				case cpuFieldIrq:
					s.IRQ = value
				case cpuFieldSoftIrq:
					s.SoftIRQ = value
				case cpuFieldSteal:
					s.Steal = value
				case cpuFieldGuest:
					s.Guest = value
				case cpuFieldGuestNice:
					s.GuestNice = value
				}
			}
			c.Cpus = append(c.Cpus, s)
		} else {
			var e error
			switch key {
			case "ctxt":
				c.Ctxt, e = parseIntE(parts[1])
				if e != nil {
					return e
				}
			case "btime":
				c.Btime, e = parseIntE(parts[1])
				if e != nil {
					return e
				}
			case "processes":
				c.Processes, e = parseIntE(parts[1])
				if e != nil {
					return e
				}
			case "procs_running":
				c.ProcsRunning, e = parseIntE(parts[1])
				if e != nil {
					return e
				}
			case "procs_blocked":
				c.ProcsBlocked, e = parseIntE(parts[1])
				if e != nil {
					return e
				}
			}
		}
	}
	return scanner.Err()
}

type CpuStat struct {
	Id        string
	User      int64
	Nice      int64
	System    int64
	Idle      int64
	IOWait    int64
	IRQ       int64
	SoftIRQ   int64
	Steal     int64
	Guest     int64
	GuestNice int64
}

func (*Cpu) Prefix() string {
	return "cpu"
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
	e = cpu.Load([]byte(str))
	if e != nil {
		return e
	}

	values := map[string]int64{
		"ctxt":          cpu.Ctxt,
		"btime":         cpu.Btime,
		"processes":     cpu.Processes,
		"procs_running": cpu.ProcsRunning,
		"procs_blocked": cpu.ProcsBlocked,
	}
	for k, v := range values {
		c.Add(k, v)
	}

	for _, cpu := range cpu.Cpus {
		tags := map[string]string{}
		if cpu.Id != "" {
			tags["cpu_id"] = cpu.Id
		} else {
			tags["total"] = "true"
		}
		values := map[string]int64{
			"user":       cpu.User,
			"nice":       cpu.Nice,
			"system":     cpu.System,
			"idle":       cpu.Idle,
			"io_wait":    cpu.IOWait,
			"irq":        cpu.IRQ,
			"soft_irq":   cpu.SoftIRQ,
			"steal":      cpu.Steal,
			"guest":      cpu.Guest,
			"guest_nice": cpu.GuestNice,
		}
		for k, v := range values {
			c.AddWithTags(k, v, tags)
		}
	}
	return nil
}
