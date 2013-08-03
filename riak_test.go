package main

func (t *testSuite) TestRiakMetrics() {
	mh := &MetricHandler{}
	col := &Riak{ Raw: readFile("fixtures/riak.json") }
	metrics, _ := mh.Collect(col)

	t.Equal(len(metrics), 126)
	m1 := metrics[0]
	t.Equal(m1.Key, "riak.VNodeGets")
	t.Equal(m1.Value, int64(241))
}
