package main

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const PGBOUNCER = "pgbouncer"

func init() {
	parser.Add(PGBOUNCER, "127.0.0.1:6432", "Collect pgbouncer metrics")
}

type PgBouncer struct {
	Address string
	Port    int
}

func (self *PgBouncer) Collect(c *MetricsCollection) (e error) {
	chunks := strings.Split(self.Address, ":")
	if len(chunks) == 2 {
		self.Address = chunks[0]
		self.Port, _ = strconv.Atoi(chunks[1])
	} else {
		self.Address = chunks[0]
		self.Port = 6432
	}
	es := []string{}

	if e = self.CollectServers(c); e != nil {
		es = append(es, e.Error())
	}

	if e = self.CollectClients(c); e != nil {
		es = append(es, e.Error())
	}

	if e = self.CollectStats(c); e != nil {
		es = append(es, e.Error())
	}

	if e = self.CollectSockets(c); e != nil {
		es = append(es, e.Error())
	}

	if e = self.CollectPools(c); e != nil {
		es = append(es, e.Error())
	}

	if e = self.CollectMemory(c); e != nil {
		es = append(es, e.Error())
	}
	if len(es) > 0 {
		e = errors.New(strings.Join(es, ", "))
	}
	return
}

func (self *PgBouncer) Keys() []string {
	return []string{
		"connections.Total",
		"memory.Free",
		"memory.MemTotal",
		"memory.Size",
		"memory.Used",
		"pools.ClientsActive",
		"pools.ClientsWaiting",
		"pools.MaxWait",
		"pools.ServersActive",
		"pools.ServersIdle",
		"pools.ServersLogin",
		"pools.ServersTested",
		"pools.ServersUsed",
		"sockets.PktAvail",
		"sockets.PktPos",
		"sockets.PktRemain",
		"sockets.RecvPos",
		"sockets.SendAvail",
		"sockets.SendPos",
		"sockets.SendRemain",
		"stats.AvgQuery",
		"stats.AvgRecv",
		"stats.AvgReq",
		"stats.AvgSent",
		"stats.TotalQueryTime",
		"stats.TotalReceived",
		"stats.TotalRequests",
		"stats.TotalSent",
	}
}

func (self *PgBouncer) Prefix() string {
	return "pgbouncer"
}

type PgBouncerConnection struct {
	Type, User, Database, State, Address, LocalAddress string
	Port, LocalPort                                    int64
	ConnectTime, RequestTime                           time.Time
}

type PgBouncerStat struct {
	Database                                                                                    string
	TotalRequests, TotalReceived, TotalSent, TotalQueryTime, AvgReq, AvgRecv, AvgSent, AvgQuery int64
}

func (self *PgBouncer) CollectStats(c *MetricsCollection) (e error) {
	allStats, e := self.Stats()
	if e == nil {
		for _, s := range allStats {
			tags := map[string]string{"database": s.Database}
			c.AddWithTags("stats.TotalRequests", s.TotalRequests, tags)
			c.AddWithTags("stats.TotalReceived", s.TotalReceived, tags)
			c.AddWithTags("stats.TotalSent", s.TotalSent, tags)
			c.AddWithTags("stats.TotalQueryTime", s.TotalQueryTime, tags)
			c.AddWithTags("stats.AvgReq", s.AvgReq, tags)
			c.AddWithTags("stats.AvgRecv", s.AvgRecv, tags)
			c.AddWithTags("stats.AvgSent", s.AvgSent, tags)
			c.AddWithTags("stats.AvgQuery", s.AvgQuery, tags)
		}
	}
	return
}

type PgBouncerFD struct {
	Task, User, Database, Address, Cancel string
	Fd, Port, Link                        int64
}

type PgBouncerSocket struct {
	Type, User, Database, State, Address, LocalAddress                                    string
	ConnectTime, RequestTime                                                              time.Time
	Port, LocalPort, RecvPos, PktPos, PktRemain, SendPos, SendRemain, PktAvail, SendAvail int64
}

type PgBouncerMemory struct {
	Name                       string
	Size, Used, Free, MemTotal int64
}

func (self *PgBouncer) CollectClients(c *MetricsCollection) (e error) {
	return self.CollectConnections("clients", c)
}

func (self *PgBouncer) CollectServers(c *MetricsCollection) (e error) {
	return self.CollectConnections("servers", c)
}

func (self *PgBouncer) CollectConnections(t string, c *MetricsCollection) (e error) {
	connections, e := self.Connections(t)
	if e != nil {
		return
	}
	for _, connection := range connections {
		tags := map[string]string{
			"database":      connection.Database,
			"state":         connection.State,
			"address":       connection.Address,
			"port":          strconv.FormatInt(connection.Port, 10),
			"local_address": connection.LocalAddress,
			"local_port":    strconv.FormatInt(connection.LocalPort, 10),
			"type":          connection.Type,
		}
		c.AddWithTags("connections.Total", 1, tags)
	}
	return
}

func (b *PgBouncer) Connections(t string) (s []*PgBouncerConnection, e error) {
	f := func(chunks []string) {
		ser := &PgBouncerConnection{
			Type:         chunks[0],
			User:         chunks[1],
			Database:     chunks[2],
			State:        chunks[3],
			Address:      chunks[4],
			Port:         parseInt64(chunks[5]),
			LocalAddress: chunks[6],
			LocalPort:    parseInt64(chunks[7]),
		}
		if time, e := time.Parse(timeFormat, chunks[8]+"Z"); e == nil {
			ser.ConnectTime = time
		}
		if time, e := time.Parse(timeFormat, chunks[9]+"Z"); e == nil {
			ser.RequestTime = time
		}
		s = append(s, ser)
	}
	e = b.Execute(t, f)
	return
}

func (b *PgBouncer) Execute(s string, f func([]string)) (e error) {
	cmd := exec.Command("psql", "-h", b.Address, "-p", strconv.Itoa(b.Port), "pgbouncer", "-c", "SHOW "+s, "-t", "-A")
	out, e := cmd.CombinedOutput()
	if e != nil {
		return
	}
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if len(line) > 0 {
			chunks := strings.Split(line, "|")
			f(chunks)
		}
	}
	return
}

type PgBouncerPool struct {
	Database, User                                                                                               string
	ClientsActive, ClientsWaiting, ServersActive, ServersIdle, ServersUsed, ServersTested, ServersLogin, MaxWait int64
}

func (self *PgBouncer) CollectPools(c *MetricsCollection) (e error) {
	pools, e := self.Pools()
	if e != nil {
		return
	}
	for _, pool := range pools {
		tags := map[string]string{
			"database": pool.Database,
		}
		for _, pool := range pools {
			c.AddWithTags("pools.ClientsActive", pool.ClientsActive, tags)
			c.AddWithTags("pools.ClientsWaiting", pool.ClientsWaiting, tags)
			c.AddWithTags("pools.ServersActive", pool.ServersActive, tags)
			c.AddWithTags("pools.ServersIdle", pool.ServersIdle, tags)
			c.AddWithTags("pools.ServersUsed", pool.ServersUsed, tags)
			c.AddWithTags("pools.ServersTested", pool.ServersTested, tags)
			c.AddWithTags("pools.ServersLogin", pool.ServersLogin, tags)
			c.AddWithTags("pools.MaxWait", pool.MaxWait, tags)
		}
	}
	return
}

func (b *PgBouncer) Pools() (ret []*PgBouncerPool, e error) {
	f := func(c []string) {
		p := &PgBouncerPool{}
		p.Database = c[0]
		p.User = c[1]
		p.ClientsActive = parseInt64(c[2])
		p.ClientsWaiting = parseInt64(c[3])
		p.ServersActive = parseInt64(c[4])
		p.ServersIdle = parseInt64(c[5])
		p.ServersUsed = parseInt64(c[6])
		p.ServersTested = parseInt64(c[7])
		p.ServersLogin = parseInt64(c[8])
		p.MaxWait = parseInt64(c[9])
		ret = append(ret, p)
	}
	b.Execute("pools", f)
	return
}

func (b *PgBouncer) Stats() (s []*PgBouncerStat, e error) {
	f := func(chunks []string) {
		ser := &PgBouncerStat{
			Database:       chunks[0],
			TotalRequests:  parseInt64(chunks[1]),
			TotalReceived:  parseInt64(chunks[2]),
			TotalSent:      parseInt64(chunks[3]),
			TotalQueryTime: parseInt64(chunks[4]),
			AvgReq:         parseInt64(chunks[5]),
			AvgRecv:        parseInt64(chunks[6]),
			AvgSent:        parseInt64(chunks[7]),
			AvgQuery:       parseInt64(chunks[8]),
		}
		s = append(s, ser)
	}
	b.Execute("stats", f)
	return
}

func (b *PgBouncer) FDs() (ret []*PgBouncerFD, e error) {
	f := func(chunks []string) {
		s := &PgBouncerFD{
			Fd:       parseInt64(chunks[0]),
			Task:     chunks[1],
			User:     chunks[2],
			Database: chunks[3],
			Address:  chunks[4],
			Port:     parseInt64(chunks[5]),
			Cancel:   chunks[6],
			Link:     parseInt64(chunks[7]),
		}
		ret = append(ret, s)
	}
	b.Execute("fds", f)
	return
}

func (b *PgBouncer) Sockets() (ret []*PgBouncerSocket, e error) {
	return b.GenericSockets("sockets")
}

func (b *PgBouncer) GenericSockets(s string) (ret []*PgBouncerSocket, e error) {
	f := func(chunks []string) {
		s := &PgBouncerSocket{
			Type:         chunks[0],
			User:         chunks[1],
			Database:     chunks[2],
			State:        chunks[3],
			Address:      chunks[4],
			Port:         parseInt64(chunks[5]),
			LocalAddress: chunks[6],
			LocalPort:    parseInt64(chunks[7]),
			RecvPos:      parseInt64(chunks[12]),
			PktPos:       parseInt64(chunks[13]),
			PktRemain:    parseInt64(chunks[14]),
			SendPos:      parseInt64(chunks[15]),
			SendRemain:   parseInt64(chunks[16]),
			PktAvail:     parseInt64(chunks[17]),
			SendAvail:    parseInt64(chunks[18]),
		}
		if time, e := time.Parse(timeFormat, chunks[8]+"Z"); e == nil {
			s.ConnectTime = time
		}
		if time, e := time.Parse(timeFormat, chunks[9]+"Z"); e == nil {
			s.RequestTime = time
		}
		ret = append(ret, s)
	}
	e = b.Execute(s, f)
	return
}

func (self *PgBouncer) CollectSockets(c *MetricsCollection) (e error) {
	sockets, e := self.Sockets()
	if e != nil {
		return
	}
	for _, socket := range sockets {
		tags := map[string]string{
			"type":          socket.Type,
			"database":      socket.Database,
			"state":         socket.State,
			"address":       socket.Address,
			"port":          strconv.FormatInt(socket.Port, 10),
			"local_address": socket.LocalAddress,
			"local_port":    strconv.FormatInt(socket.LocalPort, 10),
		}
		c.AddWithTags("sockets.RecvPos", socket.RecvPos, tags)
		c.AddWithTags("sockets.PktPos", socket.PktPos, tags)
		c.AddWithTags("sockets.PktRemain", socket.PktRemain, tags)
		c.AddWithTags("sockets.SendPos", socket.SendPos, tags)
		c.AddWithTags("sockets.SendRemain", socket.SendRemain, tags)
		c.AddWithTags("sockets.PktAvail", socket.PktAvail, tags)
		c.AddWithTags("sockets.SendAvail", socket.SendAvail, tags)
	}
	return
}

func (b *PgBouncer) Memory() (ret []*PgBouncerMemory, e error) {
	f := func(chunks []string) {
		s := &PgBouncerMemory{
			Name:     chunks[0],
			Size:     parseInt64(chunks[1]),
			Used:     parseInt64(chunks[2]),
			Free:     parseInt64(chunks[3]),
			MemTotal: parseInt64(chunks[4]),
		}
		ret = append(ret, s)
	}
	e = b.Execute("mem", f)
	return
}

func (b *PgBouncer) CollectMemory(c *MetricsCollection) (e error) {
	memories, e := b.Memory()
	if e == nil {
		for _, m := range memories {
			tags := map[string]string{"name": m.Name}
			c.AddWithTags("memory.Size", m.Size, tags)
			c.AddWithTags("memory.Used", m.Used, tags)
			c.AddWithTags("memory.Free", m.Free, tags)
			c.AddWithTags("memory.MemTotal", m.MemTotal, tags)
		}
	}
	return
}
