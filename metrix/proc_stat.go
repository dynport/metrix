package metrix

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	fieldProcStatPid = iota
	fieldProcStatComm
	fieldProcStatState
	fieldProcStatPpid
	fieldProcStatPgrp
	fieldProcStatSession               // %d  (6) The session ID of the process.
	fieldProcStatTtyNr                 // %d   (7) The controlling terminal of the process.  (The minor device number is contained in the combination of bits 31
	fieldProcStatTpgid                 // %d    (8) The ID of the foreground process group of the controlling terminal of the process.
	fieldProcStatFlags                 // %u (%lu before Linux 2.6.22)
	fieldProcStatMinflt                // %lu  (10) The number of minor faults the process has made which have not required loading a memory page from disk.
	fieldProcStatCminflt               // %lu (11) The number of minor faults that the process's waited-for children have made.
	fieldProcStatMajflt                // %lu  (12) The number of major faults the process has made which have required loading a memory page from disk.
	fieldProcStatCmajflt               // %lu (13) The number of major faults that the process's waited-for children have made.
	fieldProcStatUtime                 // %lu   (14) Amount of time that this process has been scheduled in  user  mode,  measured  in  clock  ticks  (divide  by
	fieldProcStatStime                 // %lu   (15) Amount of time that this process has been scheduled in kernel mode,  measured  in  clock  ticks  (divide  by
	fieldProcStatCutime                // %ld  (16)  Amount  of time that this process's waited-for children have been scheduled in user mode, measured in clock
	fieldProcStatCstime                // %ld  (17) Amount of time that this process's waited-for children have been scheduled in kernel mode, measured in clock
	fieldProcStatPriority              // %ld
	fieldProcStatNice                  // %ld    (19) The nice value (see setpriority(2)), a value in the range 19 (low priority) to -20 (high priority).
	fieldProcStatNumThreads            // %ld
	fieldProcStatItRealValue           // %ld
	fieldProcStatStartTime             // %llu (was %lu before Linux 2.6)
	fieldProcStatVSize                 // %lu   (23) Virtual memory size in bytes.
	fieldProcStatRSS                   // %ld     (24)  Resident  Set  Size:  number  of  pages the process has in real memory.  This is just the pages which count
	fieldProcStatRSSlim                // %lu  (25) Current soft limit in bytes on the rss of the process; see the description of RLIMIT_RSS in getrlimit(2).
	fieldProcStatStartCode             // %lu (26) The address above which program text can run.
	fieldProcStatEndCode               // %lu (27) The address below which program text can run.
	fieldProcStatStartStack            // %lu (28) The address of the start (i.e., bottom) of the stack.
	fieldProcStatKstkesp               // %lu (29) The current value of ESP (stack pointer), as found in the kernel stack page for the process.
	fieldProcStatKstkeip               // %lu (30) The current EIP (instruction pointer).
	fieldProcStatSignal                // %lu  (31)  The bitmap of pending signals, displayed as a decimal number.  Obsolete, because it does not provide infor‐
	fieldProcStatBlocked               // %lu (32) The bitmap of blocked signals, displayed as a decimal number.  Obsolete, because it does not provide  infor‐
	fieldProcStatSigignore             // %lu (33)  The bitmap of ignored signals, displayed as a decimal number.
	fieldProcStatSigcatch              // %lu (34) The bitmap of caught signals, displayed as a decimal number. Obsolete, because it does not provide informa‐
	fieldProcStatWchan                 // %lu   (35)  This  is  the  "channel"  in  which the process is waiting.  It is the address of a system call, and can be
	fieldProcStatNswap                 // %lu   (36) Number of pages swapped (not maintained).
	fieldProcStatCnswap                // %lu  (37) Cumulative nswap for child processes (not maintained).
	fieldProcStatExit_signal           // %d (since Linux 2.1.22) (38) Signal to be sent to parent when we die.
	fieldProcStatProcessor             // %d (since Linux 2.2.8) (39) CPU number last executed on.
	fieldProcStatRt_priority           // %u (since Linux 2.5.19; was %lu before Linux 2.6.22) (40)  Real-time scheduling priority,
	fieldProcStatPolicy                // %u (since Linux 2.5.19; was %lu before Linux 2.6.22) (41) Scheduling policy (see sched_setscheduler(2)).
	fieldProcStatDelayacct_blkio_ticks // %llu (since Linux 2.6.18) (42) Aggregated block I/O delays, measured in clock ticks (centiseconds).
	fieldProcStatGuest_time            // %lu (since Linux 2.6.24) (43) Guest time of the process (time spent running a virtual CPU for a guest operating
	fieldProcStatCguest_time           // %ld (since Linux 2.6.24) (44) Guest time of the process's children, measured in clock ticks
)

func numeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func eachProcDir(fun func(p string) error) error {
	files, err := filepath.Glob("/proc/*")
	if err != nil {
		return err
	}
	for _, f := range files {
		if numeric(path.Base(f)) {
			err := fun(f)
			if err != nil {
				logger.Printf("ERROR: %q", err)
			}
		}
	}
	return nil
}

func LoadProcStats() ([]*ProcStat, error) {
	defer benchmark("load proc stat")()
	out := []*ProcStat{}

	err := eachProcDir(func(f string) error {
		p, err := LoadProcStat(f + "/stat")
		if err != nil {
			return err
		}
		out = append(out, p)
		return nil
	})
	return out, err
}

func LoadProcStat(path string) (*ProcStat, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	p := &ProcStat{}
	return p, p.Load(f)
}

type ProcStat struct {
	Pid   int64  `json:"pid,omitempty,omitempty"`
	Comm  string `json:"comm,omitempty,omitempty"`
	State string `json:"state,omitempty,omitempty"`
	Ppid  int64  `json:"ppid,omitempty,omitempty"`
	Pgrp  int64  `json:"pgrp,omitempty,omitempty"`
	//Session int64 `json:"//session,omitempty"`
	//TtyNr   int64 `json:"//tty_nr,omitempty"`
	//Tpgid   int64 `json:"//tpgid,omitempty"`
	//Flags   int64 `json:"//flags,omitempty"`
	Minflt  int64 `json:"minflt,omitempty,omitempty"`
	Cminflt int64 `json:"cminflt,omitempty,omitempty"`
	Majflt  int64 `json:"majflt,omitempty,omitempty"`
	Cmajflt int64 `json:"cmajflt,omitempty,omitempty"`
	Utime   int64 `json:"utime,omitempty,omitempty"`
	Stime   int64 `json:"stime,omitempty,omitempty"`
	Cutime  int64 `json:"cutime,omitempty,omitempty"`
	Cstime  int64 `json:"cstime,omitempty,omitempty"`
	//Priority   int64 `json:"//priority,omitempty"`
	//Nice       int64 `json:"//nice,omitempty"`
	NumThreads              int64 `json:"num_threads,omitempty,omitempty"`
	VSize                   int64 `json:"v_size,omitempty,omitempty"`
	RSS                     int64 `json:"rss,omitempty,omitempty"`
	RSSlim                  int64 `json:"rs_slim,omitempty,omitempty"`
	StatStartTime           int64 `json:"stat_start_time,omitempty"`
	StatDelayacctBlkioTicks int64 `json:"stat_delayacct_blkio_ticks,omitempty"`
}

func (p *ProcStat) Load(in io.Reader) error {
	b, e := ioutil.ReadAll(in)
	if e != nil {
		return e
	}
	for i, f := range strings.Fields(string(b)) {
		switch i {
		case fieldProcStatComm:
			p.Comm = f
		case fieldProcStatState:
			p.State = f
		default:
			value, e := strconv.ParseInt(f, 10, 64)
			if e != nil {
				continue
			}
			switch i {
			case fieldProcStatPid:
				p.Pid = value
			case fieldProcStatPpid:
				p.Ppid = value
			case fieldProcStatPgrp:
				p.Pgrp = value
			case fieldProcStatMinflt:
				p.Minflt = value
			case fieldProcStatCminflt:
				p.Cminflt = value
			case fieldProcStatMajflt:
				p.Majflt = value
			case fieldProcStatCmajflt:
				p.Cmajflt = value
			case fieldProcStatUtime:
				p.Utime = value
			case fieldProcStatStime:
				p.Stime = value
			case fieldProcStatCutime:
				p.Cutime = value
			case fieldProcStatCstime:
				p.Cstime = value
			case fieldProcStatNumThreads:
				p.NumThreads = value
			case fieldProcStatRSS:
				p.RSS = value
			case fieldProcStatDelayacct_blkio_ticks:
				p.StatDelayacctBlkioTicks = value
			case fieldProcStatRSSlim:
				p.RSSlim = value
			case fieldProcStatStartTime:
				p.StatStartTime = value
			case fieldProcStatVSize:
				p.VSize = value
			}
		}
	}
	return nil
}
