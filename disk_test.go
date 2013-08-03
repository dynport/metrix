package main

func (t *testSuite) TestDiskStat() {
	mh := &MetricHandler{}
	disk := &Disk{}
	stats, _ := mh.Collect(disk)

	t.Equal(len(stats), 11)
	t.Equal(stats[0].Key, "disk.ReadsCompleted")
	t.Equal(stats[0].Value, int64(9748))
	t.Equal(stats[0].Tags["name"], "sda")
}

