package main

import (
	"regexp"
	"strconv"
)

const MEMORY = "memory"

func init() {
	parser.Add(MEMORY, "true", "Collect memory metrics")
}

type Memory struct {
}

func (m *Memory) Prefix() string {
	return "memory"
}

func (m *Memory) Keys() []string {
	return []string{
		"MemTotal", "MemFree", "Buffers", "Cached", "SwapCached", "Active", "Inactive", "Active(anon)", "Inactive(anon)", "Active(file)",
		"Inactive(file)", "Unevictable", "Mlocked", "SwapTotal", "SwapFree", "Dirty", "Writeback", "AnonPages", "Mapped", "Shmem",
		"Slab", "SReclaimable", "SUnreclaim", "KernelStack", "PageTables", "NFS_Unstable", "Bounce", "WritebackTmp", "CommitLimit",
		"Committed_AS", "VmallocTotal", "VmallocUsed", "VmallocChunk", "HardwareCorrupted", "AnonHugePages", "HugePages_Total",
		"HugePages_Free", "HugePages_Rsvd", "HugePages_Surp", "Hugepagesize", "DirectMap4k", "DirectMap2M",
	}
}

func (memory *Memory) Collect(c *MetricsCollection) (e error) {
	s := ReadProcFile("meminfo")
	re := regexp.MustCompile("(.*?)\\:\\s+(\\d+)")
	matches := re.FindAllStringSubmatch(s, -1)
	for _, a := range matches {
		k := a[1]
		v, e := strconv.ParseInt(a[2], 10, 64)
		if e == nil {
			c.Add(k, v)
		}
	}
	return
}
