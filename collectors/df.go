package collectors

import (
	"bufio"
	"io"
	"os"
	"strings"
	"syscall"
)

type FsStat struct {
	MountedOn       string
	Type            string
	BlockSize       uint32
	Blocks          uint64
	BlocksFree      uint64
	BlocksAvailable uint64
	Files           uint64
	FilesFree       uint64
}

func (stat *FsStat) Load(mountedOn string) error {
	//stat.BlockSize = t.Bsize
	//stat.Blocks = t.Blocks
	//stat.BlocksFree = t.Bfree
	//stat.BlocksAvailable = t.Bavail
	//stat.Files = t.Files
	//stat.FilesFree = t.Ffree
	return nil
}

type FileSystem struct {
}

func (fs *FileSystem) Collect(c *MetricsCollection) error {
	f, e := os.Open("/etc/mtab")
	if e != nil {
		return e
	}
	defer f.Close()
	metrics, e := parseMtab(f)
	if e != nil {
		return e
	}
	for _, m := range metrics {
		c.AddWithTags(m.Key, m.Value, m.Tags)
	}
	return nil
}

func parseMtab(f io.Reader) ([]*Metric, error) {
	metrics := []*Metric{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 2 {
			mountedOn := fields[1]
			tags := Tags{
				"file_system": fields[2],
				"mounted_on":  mountedOn,
			}
			t := &syscall.Statfs_t{}
			e := syscall.Statfs(mountedOn, t)
			if e != nil {
				return nil, e
			}
			oneK := int64(t.Bsize)
			metrics = append(metrics, []*Metric{
				{Key: "space.total", Value: int64(t.Blocks) * oneK, Tags: tags},
				{Key: "space.free", Value: int64(t.Bfree) * oneK, Tags: tags},
				{Key: "space.available", Value: int64(t.Bavail) * oneK, Tags: tags},
				{Key: "files.total", Value: int64(t.Files), Tags: tags},
				{Key: "files.free", Value: int64(t.Ffree), Tags: tags},
			}...)
		}
	}
	return metrics, nil
}
