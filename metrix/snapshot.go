package metrix

type Snapshot struct {
	Disks     []*Disk     `json:"disks,omitempty"`
	LoadAvg   *LoadAvg    `json:"load_avg,omitempty"`
	Meminfo   *Meminfo    `json:"meminfo,omitempty"`
	ProcStats []*ProcStat `json:"proc_stats,omitempty"`
	Stat      *Stat       `json:"stat,omitempty"`
}

func (s *Snapshot) Load() error {
	funcs := []func() error{
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
