package main

func (t *testSuite) TestMemory() {
	m := &Memory{}
	mh := &MetricHandler{}
	stats, _ := mh.Collect(m)

	t.True(len(stats) > 10)
	t.Equal(len(stats), 42)
	t.Equal(stats[0].Key, "memory.MemTotal")
	t.Equal(stats[0].Value, int64(502976))
}
