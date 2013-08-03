package main

func (t *testSuite) TestCpu() {
	mh := &MetricHandler{}
	stats, _ := mh.Collect(&Cpu{})
	t.Equal(len(stats), 19)
	t.Equal(stats[0].Key, "cpu.User")
	t.Equal(stats[0].Tags["total"], "true")
	t.Equal(stats[0].Value, int64(3700))

	t.Equal(stats[14].Key, "cpu.Ctxt")
	t.Equal(stats[14].Value, int64(957584))
}
