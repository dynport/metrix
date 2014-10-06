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

func LoadProcStats() ([]*ProcStat, error) {
	out := []*ProcStat{}
	e := filepath.Walk("/proc/", func(p string, info os.FileInfo, e error) error {
		dbg.Printf("checking file %s", p)
		if path.Base(p) == "stat" {
			dbg.Printf("stat file found at %s", p)
			parts := strings.Split(strings.TrimPrefix(p, "/"), "/")
			if len(parts) == 3 {
				f, e := os.Open(p)
				if e != nil {
					return e
				}
				defer f.Close()
				p, e := LoadProcStat(p)
				if e != nil {
					return e
				}
				out = append(out, p)
			} else {
				dbg.Printf("parts count is %d, %#v", len(parts), parts)
			}
		}
		return nil
	})
	return out, e
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
	Pid        int64
	Comm       string
	State      string
	Ppid       int64
	Pgrp       int64
	Session    int64
	TtyNr      int64
	Tpgid      int64
	Flags      int64
	Minflt     int64
	Cminflt    int64
	Majflt     int64
	Cmajflt    int64
	Utime      int64
	Stime      int64
	Cutime     int64
	Cstime     int64
	Priority   int64
	Nice       int64
	NumThreads int64
	VSize      int64
	RSS        int64
	RSSlim     int64
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
			case fieldProcStatRSSlim:
				p.RSSlim = value
			case fieldProcStatVSize:
				p.VSize = value
			}
		}
	}
	return nil
}
