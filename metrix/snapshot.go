package metrix

import (
	"compress/gzip"
	"encoding/json"
	"os"
	"path"
	"time"
)

type Snapshot struct {
	Disks        []*Disk        `json:"disks,omitempty"`
	Ec2Metadata  *Ec2Metadata   `json:"ec2_metadata,omitempty"`
	Hostname     string         `json:"hostname,omitempty"`
	LoadAvg      *LoadAvg       `json:"load_avg,omitempty"`
	Meminfo      *Meminfo       `json:"meminfo,omitempty"`
	ProcCmdlines []*ProcCmdline `json:"proc_cmdlines,omitempty"`
	ProcStats    []*ProcStat    `json:"proc_stats,omitempty"`
	Stat         *Stat          `json:"stat,omitempty"`
	TakenAt      time.Time      `json:"taken_at,omitempty"`
}

func (s *Snapshot) Load() error {
	funcs := []func() error{
		s.loadEc2Metadata,
		s.loadDisks,
		s.loadLoadAvg,
		s.loadMeminfo,
		s.loadProcStats,
		s.loadProcCmdlines,
		s.loadStat,
		s.loadTakenAt,
		s.loadHostname,
	}
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Snapshot) loadHostname() (err error) {
	s.Hostname, err = os.Hostname()
	return err
}

func (s Snapshot) loadTakenAt() error {
	s.TakenAt = time.Now().UTC()
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

func (s *Snapshot) loadProcCmdlines() (err error) {
	s.ProcCmdlines, err = LoadProcCmdlines()
	return err
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
