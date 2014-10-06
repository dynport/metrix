package metrix

import (
	"io"
	"io/ioutil"
	"strings"
)

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
