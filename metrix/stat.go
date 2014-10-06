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
	Cpu              *CpuStat   `json:"cpu,omitempty"`
	Cpus             []*CpuStat `json:"cpus,omitempty"`
	ContextSwitches  int64      `json:"context_switches,omitempty"`
	BootTime         int64      `json:"boot_time,omitempty"`
	Processes        int64      `json:"processes,omitempty"`
	ProcessRunning   int64      `json:"process_running,omitempty"`
	ProcessesBlocked int64      `json:"processes_blocked,omitempty"`
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
	Id        int   `json:"id,omitempty"`
	User      int64 `json:"user,omitempty"`
	Nice      int64 `json:"nice,omitempty"`
	System    int64 `json:"system,omitempty"`
	Idle      int64 `json:"idle,omitempty"`
	IOWait    int64 `json:"io_wait,omitempty"`
	IRQ       int64 `json:"irq,omitempty"`
	SoftIRQ   int64 `json:"soft_irq,omitempty"`
	Steal     int64 `json:"steal,omitempty"`
	Guest     int64 `json:"guest,omitempty"`
	QuestNice int64 `json:"quest_nice,omitempty"`
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
