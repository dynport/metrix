package collectors

import (
	"bufio"
	"strconv"
	"strings"
)

const MEMORY = "memory"

type Memory struct {
}

func (m *Memory) Prefix() string {
	return MEMORY
}

var memoryMapping = map[string]string{
	"Active(anon)":   "ActiveAnon",
	"Inactive(anon)": "InactiveAnon",
	"Active(file)":   "ActiveFile",
	"Inactive(file)": "InactiveFile",
}

func (memory *Memory) Collect(c *MetricsCollection) error {
	f, e := OpenProcFile("meminfo")
	if e != nil {
		return e
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) > 1 && strings.HasSuffix(fields[0], ":") {
			key := strings.TrimSuffix(fields[0], ":")
			if newKey, ok := memoryMapping[key]; ok {
				key = newKey
			}
			v, e := strconv.ParseInt(fields[1], 10, 64)
			if e == nil {
				c.Add(key, v)

			}
		}
	}
	return nil
}
