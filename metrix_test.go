package main

func NewTestMetrix() (m* Metrix) {
	return &Metrix{}
}

func (t *testSuite) TestRiakStatus() {
	m := &Riak{}
	raw := readFile("fixtures/riak.json")
	status, e := m.ParseRiakStatus(raw)
	if e != nil {
		t.T.Fatal(e.Error())
	}
	t.Equal(status.VNodeGets, int64(241))
	t.Equal(status.NodePutFsmActive, int64(0))
	t.Equal(status.CpuAvg1, int64(274))

	t.Equal(len(status.ConnectedNodes), 4)
	t.Equal(status.ConnectedNodes[0], "riak@192.168.0.16")

	t.Equal(len(status.RingMembers), 5)
	t.Equal(status.RingMembers[0], "riak@192.168.0.16")
}

