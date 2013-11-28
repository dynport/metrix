package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRiakStatus(t *testing.T) {
	m := &Riak{}
	raw := readFile("fixtures/riak.json")
	status, e := m.ParseRiakStatus(raw)
	if e != nil {
		t.Fatal(e.Error())
	}
	assert.Equal(t, status.VNodeGets, int64(241))
	assert.Equal(t, status.NodePutFsmActive, int64(0))
	assert.Equal(t, status.CpuAvg1, int64(274))

	assert.Equal(t, len(status.ConnectedNodes), 4)
	assert.Equal(t, status.ConnectedNodes[0], "riak@192.168.0.16")

	assert.Equal(t, len(status.RingMembers), 5)
	assert.Equal(t, status.RingMembers[0], "riak@192.168.0.16")
}

func TestSerializeMetric(t *testing.T) {
	m := &Metric{}
	b, e := json.Marshal(m)
	assert.Nil(t, e)
	assert.NotContains(t, string(b), "Tags")

	m.Tags = map[string]string{"a": "b"}
	b, e = json.Marshal(m)
	assert.Nil(t, e)
	assert.Contains(t, string(b), "Tags")
}
