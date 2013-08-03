package main

func (t *testSuite) TestDf() {
	mh := new(MetricHandler)
	col := &Df{
		RawSpace: readFile("fixtures/df.txt"),
		RawInode: readFile("fixtures/df_inode.txt"),
	}
	stats, e := mh.Collect(col)
	if e != nil {
		t.T.Log("ERROR", e.Error())
		t.Failed()
		return
	}

	t.Equal(len(stats), 8)

	t.Equal(stats[0].Key, "df.space.Total")
	t.Equal(stats[0].Value, int64(20511356))
	t.Equal(stats[0].Tags["file_system"], "/dev/sda")
	t.Equal(stats[0].Tags["mounted_on"], "/")

	t.Equal(stats[4].Key, "df.inode.Total")
	t.Equal(stats[4].Value, int64(1310720))
	t.Equal(stats[4].Tags["file_system"], "/dev/sda")
	t.Equal(stats[4].Tags["mounted_on"], "/")
}
