package metrix

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	fieldStatUser = iota
	fieldStatNice
	fieldStatSystem
	fieldStatIdle
	fieldStatIOWait
	fieldStatIRQ
	fieldStatSoftIRQ
	fieldStatSteal
	fieldStatGuest
	fieldStatQuestNice
)

type Stat struct {
	Cpu              *CpuStat
	Cpus             []*CpuStat
	ContextSwitches  int64
	BootTime         int64
	Processes        int64
	ProcessRunning   int64
	ProcessesBlocked int64
}

func LoadStat() (*Stat, error) {
	defer benchmark("load stat")()
	f, e := os.Open("/proc/stat")
	if e != nil {
		return nil, e
	}
	defer f.Close()
	s := &Stat{}
	return s, s.Load(f)
}

func (s *Stat) Load(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		switch {
		case fields[0] == "cpu":
			s.Cpu = &CpuStat{}
			e := s.Cpu.Load(fields[1:])
			if e != nil {
				return e
			}
		case strings.HasPrefix(fields[0], "cpu"):
			id, e := strconv.Atoi(strings.TrimPrefix(fields[0], "cpu"))
			if e != nil {
				return e
			}
			stat := &CpuStat{Id: id}
			e = stat.Load(fields[1:])
			if e != nil {
				return e
			}
			s.Cpus = append(s.Cpus, stat)
		default:
			v, e := parseInt64(fields[1])
			if e == nil {
				switch fields[0] {
				case "ctxt":
					s.ContextSwitches = v
				case "btime":
					s.BootTime = v
				case "processes":
					s.Processes = v
				case "procs_running":
					s.ProcessRunning = v
				case "procs_blocked":
					s.ProcessesBlocked = v
				}
			}
		}
	}
	return scanner.Err()
}

type CpuStat struct {
	Id        int
	User      int64
	Nice      int64
	System    int64
	Idle      int64
	IOWait    int64
	IRQ       int64
	SoftIRQ   int64
	Steal     int64
	Guest     int64
	QuestNice int64
}

func (stat *CpuStat) Load(parts []string) error {
	for i, raw := range parts {
		v, e := parseInt64(raw)
		if e != nil {
			return e
		}
		switch i {
		case fieldStatUser:
			stat.User = v
		case fieldStatNice:
			stat.Nice = v
		case fieldStatSystem:
			stat.System = v
		case fieldStatIdle:
			stat.Idle = v
		case fieldStatIOWait:
			stat.IOWait = v
		case fieldStatIRQ:
			stat.IRQ = v
		case fieldStatSoftIRQ:
			stat.SoftIRQ = v
		case fieldStatSteal:
			stat.Steal = v
		case fieldStatGuest:
			stat.Guest = v
		case fieldStatQuestNice:
			stat.QuestNice = v
		}

	}
	return nil
}
