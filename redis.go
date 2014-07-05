package main

import (
	"net"
	"regexp"
	"strings"
)

const REDIS = "redis"

func init() {
	parser.Add(REDIS, "127.0.0.1:6379", "Collect redis metrics")
}

type Redis struct {
	Address string
	Raw     []byte
}

func (self *Redis) Prefix() string {
	return "redis"
}

var logger = &Logger{Prefix: "redis"}

func (self *Redis) Collect(c *MetricsCollection) (e error) {
	logger.Info("collecting redis")
	b, e := self.ReadInfo()
	if e != nil {
		logger.Error("reading info", e.Error())
		return
	}
	str := string(b)
	values := map[string]string{}
	dbRegexp := regexp.MustCompile("^db(\\d+):keys=(\\d+),expires=(\\d+)")
	pid := values["process_id"]
	tcp_port := values["tcp_port"]
	for _, line := range strings.Split(str, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || len(line) < 2 {
			continue
		}
		res := dbRegexp.FindStringSubmatch(line)
		if len(res) == 4 {
			dbTags := map[string]string{
				"db":   res[1],
				"pid":  pid,
				"port": tcp_port,
			}
			c.AddWithTags("db.Keys", parseInt64(res[2]), dbTags)
			c.AddWithTags("db.Expires", parseInt64(res[3]), dbTags)
			continue
		}
		chunks := strings.Split(line, ":")
		if len(chunks) > 1 {
			values[chunks[0]] = strings.Join(chunks[1:], ":")
		}
	}
	tags := map[string]string{
		"pid":  pid,
		"port": tcp_port,
	}

	c.AddWithTags("UptimeInSeconds", parseInt64(values["uptime_in_seconds"]), tags)
	c.AddWithTags("memory.UsedMemory", parseInt64(values["used_memory"]), tags)
	c.AddWithTags("memory.UsedMemoryRSS", parseInt64(values["used_memory_rss"]), tags)
	c.AddWithTags("memory.UsedMemoryPeak", parseInt64(values["used_memory_peak"]), tags)
	c.AddWithTags("memory.UsedMemoryLua", parseInt64(values["used_memory_lua"]), tags)
	c.AddWithTags("clients.ConnectedClients", parseInt64(values["connected_clients"]), tags)
	c.AddWithTags("clients.ClientLongestOutputList", parseInt64(values["client_longest_output_list"]), tags)
	c.AddWithTags("clients.ClientBiggestInputBuf", parseInt64(values["client_biggest_input_buf"]), tags)
	c.AddWithTags("clients.BlockedClients", parseInt64(values["blocked_clients"]), tags)
	c.AddWithTags("stats.TotalConnectionsReceived", parseInt64(values["total_connections_received"]), tags)
	c.AddWithTags("stats.TotalCommandsProcessed", parseInt64(values["total_commands_processed"]), tags)
	c.AddWithTags("stats.InstantaneousOpsPerSec", parseInt64(values["instantaneous_ops_per_sec"]), tags)
	c.AddWithTags("stats.RejectedConnections", parseInt64(values["rejected_connections"]), tags)
	c.AddWithTags("stats.ExpiredKeys", parseInt64(values["expired_keys"]), tags)
	c.AddWithTags("stats.EvictedKeys", parseInt64(values["evicted_keys"]), tags)
	c.AddWithTags("stats.KeyspaceHits", parseInt64(values["keyspace_hits"]), tags)
	c.AddWithTags("stats.KeyspaceMisses", parseInt64(values["keyspace_misses"]), tags)
	c.AddWithTags("stats.PubsubChannels", parseInt64(values["pubsub_channels"]), tags)
	c.AddWithTags("stats.PubsubPatterns", parseInt64(values["pubsub_patterns"]), tags)
	c.AddWithTags("stats.LatestForkUsec", parseInt64(values["latest_fork_usec"]), tags)
	c.AddWithTags("replication.ConnectedSlaves", parseInt64(values["connected_slaves"]), tags)
	return
}

func (self *Redis) ReadInfo() (b []byte, e error) {
	if len(self.Raw) == 0 {
		var con net.Conn
		logger.Info("connecting", self.Address)
		con, e = net.Dial("tcp", self.Address)
		if e != nil {
			logger.Error("connecting", self.Address, e.Error())
			return
		}
		defer con.Close()

		con.Write([]byte("INFO\r\n"))
		b = make([]byte, 4096)
		var i int
		i, e = con.Read(b)
		logger.Finfo("read %d bytes", i)
		if e != nil {
			logger.Error("reading", e.Error())
			return
		}
		self.Raw = b
	}
	b = self.Raw
	return
}
