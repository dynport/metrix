package main

func (t *testSuite) TestLoadAvg() {
	mh := &MetricHandler{}
	stats, _ := mh.Collect(&LoadAvg{})
	t.Equal(len(stats), 3)
	t.Equal(stats[0].Value, int64(3))
	t.Equal(stats[1].Value, int64(7))
	t.Equal(stats[2].Value, int64(8))
}

