package metrix

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func LoadMeminfo() (*Meminfo, error) {
	defer benchmark("load meminfo")()
	f, e := os.Open("/proc/meminfo")
	if e != nil {
		return nil, e
	}
	defer f.Close()
	m := &Meminfo{}
	return m, m.Load(f)
}

type Meminfo struct {
	MemTotal          int64 `json:"mem_total,omitempty"`
	MemFree           int64 `json:"mem_free,omitempty"`
	Buffers           int64 `json:"buffers,omitempty"`
	Cached            int64 `json:"cached,omitempty"`
	SwapCached        int64 `json:"swap_cached,omitempty"`
	Active            int64 `json:"active,omitempty"`
	Inactive          int64 `json:"inactive,omitempty"`
	ActiveAnon        int64 `json:"active_anon,omitempty"`
	InactiveAnon      int64 `json:"inactive_anon,omitempty"`
	ActiveFile        int64 `json:"active_file,omitempty"`
	InactiveFile      int64 `json:"inactive_file,omitempty"`
	Unevictable       int64 `json:"unevictable,omitempty"`
	Mlocked           int64 `json:"mlocked,omitempty"`
	SwapTotal         int64 `json:"swap_total,omitempty"`
	SwapFree          int64 `json:"swap_free,omitempty"`
	Dirty             int64 `json:"dirty,omitempty"`
	Writeback         int64 `json:"writeback,omitempty"`
	AnonPages         int64 `json:"anon_pages,omitempty"`
	Mapped            int64 `json:"mapped,omitempty"`
	Shmem             int64 `json:"shmem,omitempty"`
	Slab              int64 `json:"slab,omitempty"`
	SReclaimable      int64 `json:"s_reclaimable,omitempty"`
	SUnreclaim        int64 `json:"s_unreclaim,omitempty"`
	KernelStack       int64 `json:"kernel_stack,omitempty"`
	PageTables        int64 `json:"page_tables,omitempty"`
	NFS_Unstable      int64 `json:"nfs_unstable,omitempty"`
	Bounce            int64 `json:"bounce,omitempty"`
	WritebackTmp      int64 `json:"writeback_tmp,omitempty"`
	CommitLimit       int64 `json:"commit_limit,omitempty"`
	Committed_AS      int64 `json:"committed_as,omitempty"`
	VmallocTotal      int64 `json:"vmalloc_total,omitempty"`
	VmallocUsed       int64 `json:"vmalloc_used,omitempty"`
	VmallocChunk      int64 `json:"vmalloc_chunk,omitempty"`
	HardwareCorrupted int64 `json:"hardware_corrupted,omitempty"`
	AnonHugePages     int64 `json:"anon_huge_pages,omitempty"`
	HugePages_Total   int64 `json:"huge_pages_total,omitempty"`
	HugePages_Free    int64 `json:"huge_pages_free,omitempty"`
	HugePages_Rsvd    int64 `json:"huge_pages_rsvd,omitempty"`
	HugePages_Surp    int64 `json:"huge_pages_surp,omitempty"`
	Hugepagesize      int64 `json:"hugepagesize,omitempty"`
	DirectMap4k       int64 `json:"direct_map4k,omitempty"`
	DirectMap2M       int64 `json:"direct_map2_m,omitempty"`
	DirectMap1G       int64 `json:"direct_map1_g,omitempty"`
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
