package main

import (
	"time"
)

func (t *testSuite) TestOpenTSDB() {
	theTime := time.Unix(11, 0)
	m := &Metric{Key: "metric", Value: int64(10)}
	t.Equal(m.OpenTSDB(theTime, "test.host"), "put metric 11 10 host=test.host")

	m = &Metric{Key: "os.cpu.User", Value: int64(10), Tags: map[string]string { "cpu_id": "1" }}
	t.Equal(m.OpenTSDB(theTime, "test.host"), "put os.cpu.User 11 10 host=test.host cpu_id=1")

	m = &Metric{Key: "os.cpu.User", Value: int64(10), Tags: map[string]string { "name": "(kworker/0:2)" }}
	t.Equal(m.OpenTSDB(theTime, "test.host"), "put os.cpu.User 11 10 host=test.host name=kworker_0_2")
}
