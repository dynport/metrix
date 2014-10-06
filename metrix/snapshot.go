package metrix

import (
	"compress/gzip"
	"encoding/json"
	"os"
	"path"
	"time"
)

type Snapshot struct {
	Hostname    string       `json:"hostname,omitempty"`
	Ec2Metadata *Ec2Metadata `json:"ec2_metadata,omitempty"`
	Disks       []*Disk      `json:"disks,omitempty"`
	LoadAvg     *LoadAvg     `json:"load_avg,omitempty"`
	Meminfo     *Meminfo     `json:"meminfo,omitempty"`
	ProcStats   []*ProcStat  `json:"proc_stats,omitempty"`
	Stat        *Stat        `json:"stat,omitempty"`
}

func (s *Snapshot) Load() error {
	funcs := []func() error{
		s.loadEc2Metadata,
		s.loadDisks,
		s.loadLoadAvg,
		s.loadMeminfo,
		s.loadProcStats,
		s.loadStat,
	}
	for _, f := range funcs {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Snapshot) Store(rootDir string) error {
	var err error
	if s.Hostname == "" {
		s.Hostname, err = os.Hostname()
		if err != nil {
			return err
		}
	}
	localPath := rootDir + "/snapshots/" + time.Now().UTC().Format("2006/01/02/15/2006-01-02T150405") + "-" + s.Hostname + ".json.gz"
	logger.Printf("writing snapshot to %s", localPath)
	tmpPath := localPath + ".tmp"
	err = os.MkdirAll(path.Dir(tmpPath), 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(tmpPath)
	if err != nil {
		return err
	}
	defer f.Close()
	gz := gzip.NewWriter(f)
	defer gz.Close()
	enc := json.NewEncoder(gz)
	err = enc.Encode(s)
	if err != nil {
		return err
	}
	return os.Rename(tmpPath, localPath)
}

func (s *Snapshot) loadEc2Metadata() (err error) {
	defer benchmark("load ec2 metadata")()
	em, err := LoadEc2Metadata()
	if err != nil {
		dbg.Printf("ERROR loading ec2 metadata: %q", err)
	} else {
		s.Ec2Metadata = em
	}
	return nil
}

func (s *Snapshot) loadProcStats() (err error) {
	s.ProcStats, err = LoadProcStats()
	return err
}

func (s *Snapshot) loadMeminfo() (err error) {
	s.Meminfo, err = LoadMeminfo()
	return err
}

func (s *Snapshot) loadLoadAvg() (err error) {
	s.LoadAvg, err = LoadLoadAvg()
	return err
}

func (s *Snapshot) loadDisks() (err error) {
	s.Disks, err = LoadDisks()
	return err
}

func (s *Snapshot) loadStat() (err error) {
	s.Stat, err = LoadStat()
	return err
}
