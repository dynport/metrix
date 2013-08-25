package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	FILES = "files"
	FILES_OPEN = "Open"
)

var LSOF_CMD = `lsof | tail -n +2 | tr -s " " "` + "\t" + `" | cut -f 1,2`

func init() {
	parser.Add(FILES, "true", "Collect open files")
}

type Files struct {
	RawStatus []byte
}

func (files *Files) fetch() (b []byte, e error) {
	if len(files.RawStatus) == 0 {
		args := []string { "-c" }
		args = append(args, LSOF_CMD)
		files.RawStatus, _ = exec.Command("bash", args...).Output()
		if len(files.RawStatus) == 0 {
			return b, fmt.Errorf("lsof returned empty result")
		}
	}
	return files.RawStatus, nil
}

func (f *Files) Collect(col *MetricsCollection) error {
	b, e := f.fetch()
	if e != nil {
		return e
	}
	stats := map[string]int64{}
	for _, line := range strings.Split(string(b), "\n") {
		stats[line]++
	}
	for k, v := range stats {
		chunks := strings.SplitN(k, "\t", 2)
		if len(chunks) == 2 {
			name, pid := chunks[0], chunks[1]
			tags := map[string]string{"name": NormalizeProcessName(name), "pid": pid}
			col.AddWithTags(FILES_OPEN, v, tags)
		}
	}
	return nil
}

func (f *Files) Keys() []string {
	return []string{FILES_OPEN}
}

func (f *Files) Prefix() string {
	return FILES
}
