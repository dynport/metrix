package metrix

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func LoadProcCmdlines() ([]*ProcCmdline, error) {
	defer benchmark("load proc cmdlines")()
	out := []*ProcCmdline{}
	err := eachProcDir(func(dir string) error {
		p := &ProcCmdline{}
		f, err := os.Open(dir + "/cmdline")
		if err != nil {
			return err
		}
		defer f.Close()
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
	Cmd  string   `json:"cmd,omitempty"`
	Args []string `json:"args,omitempty"`
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
