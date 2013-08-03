package main

func (t *testSuite) TestProcesss() {
	mh := new(MetricHandler)
	col := &Processes{}
	stats, _ := mh.Collect(col)
	t.True(len(stats) > 0)
	t.Equal(len(stats), 980)
	t.Equal(stats[0].Key, "processes.Utime")
	t.Equal(stats[0].Value, int64(25))
	t.Equal(stats[0].Tags["comm"], "(init)")
	t.Equal(stats[0].Tags["state"], "S")
	t.Equal(stats[0].Tags["pid"], "1")
}
