package main

func (t *testSuite) TestNet() {
	mh := new(MetricHandler)
	net := &Net{RawStatus: readFile("fixtures/netstat.txt")}
	stats, _ := mh.Collect(net)
	t.True(len(stats) > 4)
	t.Equal(len(stats), 21)

	t.Equal(stats[0].Key, "net.ip.TotalPacketsReceived")
	t.Equal(stats[0].Value, int64(162673))

	t.Equal(stats[20].Key, "net.ip.OutOctets")
	t.Equal(stats[20].Value, int64(667161104))
}
