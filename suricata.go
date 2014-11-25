package main

import (
	"fmt"
	"net"
	"io"
	"encoding/json"
	"errors"
)

const SURICATA = "suricata"

func init() {
	parser.Add(SURICATA, "/usr/local/var/run/suricata/suricata-command.socket", "dump Suricata's performance counters")
}


type Suricata struct {
	SocketName string
}

func (self *Suricata) Prefix() string {
	return SURICATA
}

func (self *Suricata) Collect(c *MetricsCollection) (e error) {
	name := "testTag"
	tags := map[string]string{"name": name}
	c.AddWithTags("suricata.test", 1024, tags)
	r,e := getCountersFromSocket(self.SocketName)
	if e != nil {
		return
	}
	for k,v := range r {
		for kk,vv := range v {
			// add thread name as tag
			if k != "Detect" && k != "FlowManagerThread" {
				tags := map[string]string{"thread": k}
				c.AddWithTags(kk, vv, tags)
			} else {
				c.Add(kk, vv)
			}
		}
	}
	return
}

type Response struct {
    ReturnCode string         `json:"return"`
    Message map[string]map[string]int64 `json:"message"`
}

func getCountersFromSocket(socketName string) (counts map[string]map[string]int64, err error) {
	conn, err := net.Dial("unix", socketName)	
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// see https://github.com/inliniac/suricata/blob/94571c5dd28858ff68c44b648fd41c5d87c0e28d/src/unix-manager.c#L288
	_, err = fmt.Fprint(conn, "{\"version\": \"0.1\"}")
	if err != nil {
			return nil, errors.New("can not send version :: "+err.Error())
	}
	var buf [128]byte
	_, err = conn.Read(buf[0:])
	if err != nil {
			// see https://github.com/inliniac/suricata/blob/94571c5dd28858ff68c44b648fd41c5d87c0e28d/src/unix-manager.c#L319
			return nil, errors.New("got kick on version :: "+err.Error())
	}
	_, err = fmt.Fprint(conn, "{\"command\": \"dump-counters\"}")
	if err != nil {
		return nil, errors.New("can not send command :: "+err.Error())
	}
	
	data := ""
	// read anser for command
	//  see https://github.com/inliniac/suricata/blob/672049632431bb695f56798c9c5f196afcf2fb27/scripts/suricatasc/src/suricatasc.py#L83
	for i := 0; i < 32; i++ {
		var buf2 [4096]byte
		result2, err := conn.Read(buf2[0:])
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
		}
		data = data + string(buf2[:result2])
		var isJson  Response 
		jerr := json.Unmarshal([]byte(data), &isJson)
		if jerr == nil {
			if isJson.ReturnCode == "OK" {
				return isJson.Message, nil
			} else {
				return nil, errors.New(isJson.ReturnCode)
			}

		} else {
			expect := "unexpected end of JSON input"
			if jerr.Error() != expect {
				// is json but not our counters
				return nil, errors.New(data) 
			}
		}
	}
	return nil, errors.New("max limit for loop")
}