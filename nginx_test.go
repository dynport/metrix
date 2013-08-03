package main

func (t *testSuite) TestParseNginx() {
	logger.LogLevel = INFO
	mh := &MetricHandler{}
	nginx := &Nginx{ Raw: readFile("fixtures/nginx.status") }

	all, _ := mh.Collect(nginx)
	t.Equal(len(all), 7)

	t.Equal(all[0].Key, "nginx.ActiveConnections")
	t.Equal(all[0].Value, int64(10))

	t.Equal(all[6].Key, "nginx.Waiting")
	t.Equal(all[6].Value, int64(70))

	return
}
