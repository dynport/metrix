package metrix

import (
	"bufio"
	"io"

	"os/exec"
	"strconv"
	"strings"
)

type Disk struct {
	Filesystem string
	Blocks     int64
	Used       int64
	Available  int64
	Use        int
	MountedOn  string
}

func LoadDisks() ([]*Disk, error) {
	defer benchmark("load disks")()
	c := exec.Command("df", "-k")
	out, e := c.StdoutPipe()
	if e != nil {
		return nil, e
	}
	defer out.Close()
	e = c.Start()
	if e != nil {
		return nil, e
	}
	return ParseDf(out)
}

const (
	fieldDiskFilesystem = iota
	fieldDiskBlocks
	fieldDiskUsed
	fieldDiskAvailable
	fieldDiskUse
	fieldDiskMountedOn
)

func ParseDf(in io.Reader) ([]*Disk, error) {
	scanner := bufio.NewScanner(in)
	disks := []*Disk{}
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[0] == "Filesystem" {
			// header
			continue
		}
		if len(fields) != 6 {
			dbg.Printf("expected 6 fields, got %d", len(fields))
			continue
		}
		d := &Disk{}
		var e error
		e = func() error {
			for i, v := range fields {
				switch i {
				case fieldDiskFilesystem:
					d.Filesystem = v
				case fieldDiskBlocks:
					if d.Blocks, e = parseInt64(v); e != nil {
						return e
					}
				case fieldDiskUsed:
					if d.Used, e = parseInt64(v); e != nil {
						return e
					}
				case fieldDiskAvailable:
					if d.Available, e = parseInt64(v); e != nil {
						return e
					}
				case fieldDiskUse:
					if strings.HasSuffix(v, "%") {
						d.Use, e = strconv.Atoi(strings.TrimSuffix(v, "%"))
						if e != nil {
							return e
						}
					}
				case fieldDiskMountedOn:
					d.MountedOn = v
				}
			}
			return nil
		}()
		if e != nil {
			logger.Printf("ERROR: %q", e)
		} else {
			disks = append(disks, d)
		}
	}
	return disks, scanner.Err()
}

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)

}
