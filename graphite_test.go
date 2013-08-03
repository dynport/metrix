package main

import (
	"time"
)

func (t *testSuite) TestGraphite() {
	theTime := time.Unix(11, 0)
	m := &Metric{Key: "metric", Value: int64(10)}
	t.Equal(m.Graphite(theTime, "test.host"), "metrix.hosts.test.host.metric 10 11")

	m = &Metric{Key: "os.cpu.User", Value: int64(10), Tags: map[string]string { "cpu_id": "1" }}
	t.Equal(m.Graphite(theTime, "test.host"), "metrix.hosts.test.host.os.cpu1.User 10 11")

	m = &Metric{Key: "disk.WeightedMillisecondsIO", Value: int64(10), Tags: map[string]string { "name": "sda" }}
	t.Equal(m.Graphite(theTime, "test.host"), "metrix.hosts.test.host.disk.sda.WeightedMillisecondsIO 10 11")

	m = &Metric{Key: "df.inodes.Use", Value: int64(10), Tags: map[string]string { "file_system": "/dev/sda" }}
	t.Equal(m.Graphite(theTime, "test.host"), "metrix.hosts.test.host.df.dev.sda.inodes.Use 10 11")
}
