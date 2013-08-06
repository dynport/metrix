package main

import (
	"errors"
	"os/exec"
	"strings"
	"regexp"
	"strconv"
)

type Df struct {
	RawSpace []byte
	RawInode []byte
}

func (self *Df) Prefix() string {
	return "df"
}

const (
	SPACE = "space"
	INODE = "inode"
	DF = "df"
)

func init() {
	parser.Add(DF, "true", "Collect disk free space metrics")
}

var dfFlgaMapping = map[string]string {
	SPACE: "-k",
	INODE: "-i",
}


func (self *Df) Keys() []string {
	return []string {
		SPACE + ".Total",
		SPACE + ".Used",
		SPACE + ".Available",
		SPACE + ".Use",
		INODE + ".Total",
		INODE + ".Used",
		INODE + ".Available",
		INODE + ".Use",
	}
}

func CollectDf(t string, raw []byte, c* MetricsCollection) (e error) {
	if len(raw) == 0 {
		if flag, ok := dfFlgaMapping[t]; ok {
			raw, e = ReadDf(flag)
		} else {
			return errors.New("no mapping for df key " + t + " defined")
		}
	}
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
	re := regexp.MustCompile("\\s+")
	for _, line := range lines[1:] {
		if !strings.HasPrefix(line, "/") {
			continue
		}
		chunks := re.Split(line, 6)
		if len(chunks) == 6 {
			tags := map[string]string {
				"file_system": chunks[0],
				"mounted_on": chunks[5],
			}
			if v, e := strconv.ParseInt(chunks[1], 10, 64); e == nil {
				c.AddWithTags(t + ".Total", v, tags)
			}
			if v, e := strconv.ParseInt(chunks[2], 10, 64); e == nil {
				c.AddWithTags(t + ".Used", v, tags)
			}
			if v, e := strconv.ParseInt(chunks[3], 10, 64); e == nil {
				c.AddWithTags(t + ".Available", v, tags)
			}
			if v, e := strconv.ParseInt(strings.Replace(chunks[4], "%", "", 1), 10, 64); e == nil {
				c.AddWithTags(t + ".Use", v, tags)
			}
		}
	}
	return
}

func (self *Df) Collect(c* MetricsCollection) (e error) {
	e = CollectDf(SPACE, self.RawSpace, c)
	if e != nil {
		return
	}
	e = CollectDf(INODE, self.RawInode, c)
	return
}

func ReadDf(flag string) (b []byte, e error) {
	logger.Debug("reading df with", flag)
	b, e = exec.Command("df", flag).Output()
	if e != nil {
		return
	}
	if len(b) == 0 {
		e = errors.New("Reading df returned an empty string")
	}
	return
}
