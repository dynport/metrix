package metrix

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func LoadMeminfo() (*Meminfo, error) {
	f, e := os.Open("/proc/meminfo")
	if e != nil {
		return nil, e
	}
	defer f.Close()
	m := &Meminfo{}
	return m, m.Load(f)
}

type Meminfo struct {
	MemTotal          int64
	MemFree           int64
	Buffers           int64
	Cached            int64
	SwapCached        int64
	Active            int64
	Inactive          int64
	ActiveAnon        int64
	InactiveAnon      int64
	ActiveFile        int64
	InactiveFile      int64
	Unevictable       int64
	Mlocked           int64
	SwapTotal         int64
	SwapFree          int64
	Dirty             int64
	Writeback         int64
	AnonPages         int64
	Mapped            int64
	Shmem             int64
	Slab              int64
	SReclaimable      int64
	SUnreclaim        int64
	KernelStack       int64
	PageTables        int64
	NFS_Unstable      int64
	Bounce            int64
	WritebackTmp      int64
	CommitLimit       int64
	Committed_AS      int64
	VmallocTotal      int64
	VmallocUsed       int64
	VmallocChunk      int64
	HardwareCorrupted int64
	AnonHugePages     int64
	HugePages_Total   int64
	HugePages_Free    int64
	HugePages_Rsvd    int64
	HugePages_Surp    int64
	Hugepagesize      int64
	DirectMap4k       int64
	DirectMap2M       int64
	DirectMap1G       int64
}

func (m *Meminfo) Load(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 {
			value, e := parseInt64(fields[1])
			if e != nil {
				dbg.Printf("unable to parse %q in %q as int64", fields[1], scanner.Text())
				continue
			}
			switch fields[0] {
			case "MemTotal:":
				m.MemTotal = value
			case "MemFree:":
				m.MemFree = value
			case "Buffers:":
				m.Buffers = value
			case "Cached:":
				m.Cached = value
			case "SwapCached:":
				m.SwapCached = value
			case "Active:":
				m.Active = value
			case "Inactive:":
				m.Inactive = value
			case "Active(anon):":
				m.ActiveAnon = value
			case "Inactive(anon):":
				m.InactiveAnon = value
			case "Active(file):":
				m.ActiveFile = value
			case "Inactive(file):":
				m.InactiveFile = value
			case "Unevictable:":
				m.Unevictable = value
			case "Mlocked:":
				m.Mlocked = value
			case "SwapTotal:":
				m.SwapTotal = value
			case "SwapFree:":
				m.SwapFree = value
			case "Dirty:":
				m.Dirty = value
			case "Writeback:":
				m.Writeback = value
			case "AnonPages:":
				m.AnonPages = value
			case "Mapped:":
				m.Mapped = value
			case "Shmem:":
				m.Shmem = value
			case "Slab:":
				m.Slab = value
			case "SReclaimable:":
				m.SReclaimable = value
			case "SUnreclaim:":
				m.SUnreclaim = value
			case "KernelStack:":
				m.KernelStack = value
			case "PageTables:":
				m.PageTables = value
			case "NFS_Unstable:":
				m.NFS_Unstable = value
			case "Bounce:":
				m.Bounce = value
			case "WritebackTmp:":
				m.WritebackTmp = value
			case "CommitLimit:":
				m.CommitLimit = value
			case "Committed_AS:":
				m.Committed_AS = value
			case "VmallocTotal:":
				m.VmallocTotal = value
			case "VmallocUsed:":
				m.VmallocUsed = value
			case "VmallocChunk:":
				m.VmallocChunk = value
			case "HardwareCorrupted:":
				m.HardwareCorrupted = value
			case "AnonHugePages:":
				m.AnonHugePages = value
			case "HugePages_Total:":
				m.HugePages_Total = value
			case "HugePages_Free:":
				m.HugePages_Free = value
			case "HugePages_Rsvd:":
				m.HugePages_Rsvd = value
			case "HugePages_Surp:":
				m.HugePages_Surp = value
			case "Hugepagesize:":
				m.Hugepagesize = value
			case "DirectMap4k:":
				m.DirectMap4k = value
			case "DirectMap2M:":
				m.DirectMap2M = value
			case "DirectMap1G:":
				m.DirectMap1G = value
			}
		}
	}
	return scanner.Err()

}
