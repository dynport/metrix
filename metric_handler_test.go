package main

func (t *testSuite) TestMetricKeys() {
	col := &MetricHandler{}
	es := &ElasticSearch{}
	keys := col.Keys(es)
	t.True(len(keys) > 0)
	t.Equal(keys[0], "elasticsearch.shards.Total")
}
