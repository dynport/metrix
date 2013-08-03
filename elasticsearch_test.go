package main

func (t *testSuite) TestParseElasticSearch() {
	mh := &MetricHandler{}
	es := &ElasticSearch{ RawStatus: readFile("fixtures/elasticsearch_status.json") }

	all, _ := mh.Collect(es)

	mapped := map[string]*Metric {}
	for _, m := range all {
		k := m.Key
		if name, ok := m.Tags["index_name"]; ok {
			k = name + "." + k
		}
		mapped[k] = m
	}
	t.Equal(mapped["elasticsearch.shards.Total"].Value, int64(20))
	t.Equal(mapped["elasticsearch.shards.Successful"].Value, int64(10))
	t.Equal(mapped["index.elasticsearch.indices.index.SizeInBytes"].Value, int64(2161))
	t.Equal(mapped["index.elasticsearch.indices.merges.TotalDocs"].Value, int64(1))
	t.True(len(mapped) > 5)
}
