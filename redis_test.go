package main

func (t *testSuite) TestParseRedis() {
	logger.LogLevel = WARN
	mh := &MetricHandler{}
	es := &Redis{ Raw: readFile("fixtures/redis_info.txt") }

	all, _ := mh.Collect(es)
	t.True(len(all) > 0)

	return
}
