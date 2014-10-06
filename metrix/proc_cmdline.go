package metrix

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func LoadProcCmdlines() ([]*ProcCmdline, error) {
	defer benchmark("load proc cmdlines")()
	out := []*ProcCmdline{}
	var err error
	err = eachProcDir(func(dir string) error {
		p := &ProcCmdline{}
		localPath := dir + "/cmdline"
		p.Pid, err = strconv.Atoi(path.Base(dir))
		if err != nil {
			return err
		}
		f, err := os.Open(localPath)
		if err != nil {
			return err
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			return err
		}
		p.StartedAt = stat.ModTime().UTC()
		err = p.Load(f)
		if err != nil {
			return err
		}
		if p.Cmd != "" {
			out = append(out, p)
		}
		return nil
	})
	return out, err
}

type ProcCmdline struct {
	Pid       int       `json:"pid,omitempty"`
	Cmd       string    `json:"cmd,omitempty"`
	Args      []string  `json:"args,omitempty"`
	StartedAt time.Time `json:"started_at,omitempty"`
}

func (p *ProcCmdline) Load(in io.Reader) error {
	b, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	fields := strings.Split(string(b), "\x00")
	for i, s := range fields {
		if s == "" {
			continue
		}
		if i == 0 {
			p.Cmd = s
		} else {
			p.Args = append(p.Args, s)
		}
	}
	return nil
}
