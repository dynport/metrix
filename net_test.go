package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNet(t *testing.T) {
	mh := new(MetricHandler)
	net := &Net{RawStatus: readFile("fixtures/netstat.txt")}
	stats, _ := mh.Collect(net)
	assert.True(t, len(stats) > 4)
	assert.Equal(t, len(stats), 21)

	assert.Equal(t, stats[0].Key, "net.ip.TotalPacketsReceived")
	assert.Equal(t, stats[0].Value, int64(162673))

	assert.Equal(t, stats[20].Key, "net.ip.OutOctets")
	assert.Equal(t, stats[20].Value, int64(667161104))
}
