package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

const (
	FREE             = "free"
	FREE_MEM_TOTAL   = "mem.total"
	FREE_MEM_USED    = "mem.used"
	FREE_MEM_FREE    = "mem.free"
	FREE_MEM_BUFFERS = "mem.buffers"
	FREE_MEM_CACHED  = "mem.cached"
	FREE_SWAP_TOTAL  = "swap.total"
	FREE_SWAP_USED   = "swap.used"
	FREE_SWAP_FREE   = "swap.free"
)

var FREE_MEMORY_REGEXP = regexp.MustCompile("Mem:\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)")
var FREE_SWAP_REGEXP = regexp.MustCompile("Swap:\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)")

func init() {
	parser.Add(FREE, "true", "Collect free metrics")
}

type Free struct {
	RawStatus []byte
}

func (free *Free) fetch() (b []byte, e error) {
	if len(free.RawStatus) == 0 {
		free.RawStatus, _ = exec.Command("free").Output()
		if len(free.RawStatus) == 0 {
			return b, fmt.Errorf("free returned empty output")
		}
	}
	return free.RawStatus, nil
}

func (free *Free) Collect(c *MetricsCollection) error {
	s, e := free.fetch()
	if e != nil {
		return e
	}
	raw := string(s)
	parts := FREE_MEMORY_REGEXP.FindStringSubmatch(raw)
	if len(parts) > 0 {
		c.Add(FREE_MEM_TOTAL, parseInt64(parts[1]))
		c.Add(FREE_MEM_USED, parseInt64(parts[2]))
		c.Add(FREE_MEM_FREE, parseInt64(parts[3]))
		c.Add(FREE_MEM_BUFFERS, parseInt64(parts[5]))
		c.Add(FREE_MEM_CACHED, parseInt64(parts[6]))
	}

	parts = FREE_SWAP_REGEXP.FindStringSubmatch(raw)
	if len(parts) > 0 {
		c.Add(FREE_SWAP_TOTAL, parseInt64(parts[1]))
		c.Add(FREE_SWAP_USED, parseInt64(parts[2]))
		c.Add(FREE_SWAP_FREE, parseInt64(parts[3]))
	}
	return nil
}

func (free *Free) Prefix() string {
	return "free"
}
