package metrix

import (
	"bufio"
	"io"
	"os/exec"
	"strings"
)

func LoadEc2Metadata() (*Ec2Metadata, error) {
	c := exec.Command("ec2metadata")
	out, err := c.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer out.Close()
	err = c.Start()
	if err != nil {
		return nil, err
	}
	em := &Ec2Metadata{}
	if err := em.Load(out); err != nil {
		return nil, err
	}
	return em, nil
}

type Ec2Metadata struct {
	AvailabilityZone string `json:"availability_zone,omitempty"`
	InstanceId       string `json:"instance_id,omitempty"`
	InstanceType     string `json:"instance_type,omitempty"`
	AmiId            string `json:"ami_id,omitempty"`
}

func (e *Ec2Metadata) Load(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), ": ", 2)
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			switch key {
			case "availability-zone":
				e.AvailabilityZone = value
			case "instance-id":
				e.InstanceId = value
			case "ami-id":
				e.AmiId = value
			}
		}
	}
	return scanner.Err()
}
