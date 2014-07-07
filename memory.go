package main

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
)

const MEMORY = "memory"

func init() {
	parser.Add(MEMORY, "true", "Collect memory metrics")
}

type Memory struct {
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
}

func (m *Memory) Prefix() string {
	return "memory"
}

func (m *Memory) Load(b []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := parts[0]
			valueString := strings.Fields(parts[1])[0]
			value, e := strconv.ParseInt(valueString, 10, 64)
			if e != nil {
				log.Printf("error parsing %q to int64: %e", valueString, e)
				continue
			}
			switch key {
			case "MemTotal":
				m.MemTotal = value
			case "MemFree":
				m.MemFree = value
			case "Buffers":
				m.Buffers = value
			case "Cached":
				m.Cached = value
			case "SwapCached":
				m.SwapCached = value
			case "Active":
				m.Active = value
			case "Inactive":
				m.Inactive = value
			case "Active(anon)":
				m.ActiveAnon = value
			case "Inactive(anon)":
				m.InactiveAnon = value
			case "Active(file)":
				m.ActiveFile = value
			case "Inactive(file)":
				m.InactiveFile = value
			case "Unevictable":
				m.Unevictable = value
			case "Mlocked":
				m.Mlocked = value
			case "SwapTotal":
				m.SwapTotal = value
			case "SwapFree":
				m.SwapFree = value
			case "Dirty":
				m.Dirty = value
			case "Writeback":
				m.Writeback = value
			case "AnonPages":
				m.AnonPages = value
			case "Mapped":
				m.Mapped = value
			case "Shmem":
				m.Shmem = value
			case "Slab":
				m.Slab = value
			case "SReclaimable":
				m.SReclaimable = value
			case "SUnreclaim":
				m.SUnreclaim = value
			case "KernelStack":
				m.KernelStack = value
			case "PageTables":
				m.PageTables = value
			case "NFS_Unstable":
				m.NFS_Unstable = value
			case "Bounce":
				m.Bounce = value
			case "WritebackTmp":
				m.WritebackTmp = value
			case "CommitLimit":
				m.CommitLimit = value
			case "Committed_AS":
				m.Committed_AS = value
			case "VmallocTotal":
				m.VmallocTotal = value
			case "VmallocUsed":
				m.VmallocUsed = value
			case "VmallocChunk":
				m.VmallocChunk = value
			case "HardwareCorrupted":
				m.HardwareCorrupted = value
			case "AnonHugePages":
				m.AnonHugePages = value
			case "HugePages_Total":
				m.HugePages_Total = value
			case "HugePages_Free":
				m.HugePages_Free = value
			case "HugePages_Rsvd":
				m.HugePages_Rsvd = value
			case "HugePages_Surp":
				m.HugePages_Surp = value
			case "Hugepagesize":
				m.Hugepagesize = value
			case "DirectMap4k":
				m.DirectMap4k = value
			case "DirectMap2M":
				m.DirectMap2M = value
			default:
				log.Printf("ERROR: key=%q unknown", key)
			}
		}
	}
	return scanner.Err()
}

func (m *Memory) Collect(c *MetricsCollection) (e error) {
	s := ReadProcFile("meminfo")
	e = m.Load([]byte(s))
	if e != nil {
		return e
	}

	values := map[string]int64{
		"mem_total":          m.MemTotal,
		"mem_free":           m.MemFree,
		"buffers":            m.Buffers,
		"cached":             m.Cached,
		"swap_cached":        m.SwapCached,
		"active":             m.Active,
		"inactive":           m.Inactive,
		"active_anon":        m.ActiveAnon,
		"inactive_anon":      m.InactiveAnon,
		"active_file":        m.ActiveFile,
		"inactive_file":      m.InactiveFile,
		"unevictable":        m.Unevictable,
		"mlocked":            m.Mlocked,
		"swap_total":         m.SwapTotal,
		"swap_free":          m.SwapFree,
		"dirty":              m.Dirty,
		"writeback":          m.Writeback,
		"anon_pages":         m.AnonPages,
		"mapped":             m.Mapped,
		"shmem":              m.Shmem,
		"slab":               m.Slab,
		"s_reclaimable":      m.SReclaimable,
		"s_unreclaim":        m.SUnreclaim,
		"kernel_stack":       m.KernelStack,
		"page_tables":        m.PageTables,
		"nfs_unstable":       m.NFS_Unstable,
		"bounce":             m.Bounce,
		"writeback_tmp":      m.WritebackTmp,
		"commit_limit":       m.CommitLimit,
		"committed_as":       m.Committed_AS,
		"vmalloc_total":      m.VmallocTotal,
		"vmalloc_used":       m.VmallocUsed,
		"vmalloc_chunk":      m.VmallocChunk,
		"hardware_corrupted": m.HardwareCorrupted,
		"anon_huge_pages":    m.AnonHugePages,
		"hugepages_total":    m.HugePages_Total,
		"hugepages_free":     m.HugePages_Free,
		"hugepages_rsvd":     m.HugePages_Rsvd,
		"hugepages_surp":     m.HugePages_Surp,
		"hugepagesize":       m.Hugepagesize,
		"direct_map4k":       m.DirectMap4k,
		"direct_map2_m":      m.DirectMap2M,
	}

	for k, v := range values {
		c.Add(k, v)
	}
	return nil
}
