package main

import (
	"encoding/json"
)

const ELASTICSEARCH = "elasticsearch"

func init() {
	parser.Add(ELASTICSEARCH, "http://127.0.0.1:9200/_status", "Collect ElasticSearch metrics")
}

type ElasticSearch struct {
	Url       string
	RawStatus []byte
}

func (es *ElasticSearch) Prefix() string {
	return "elasticsearch"
}

func (es *ElasticSearch) Collect(c *MetricsCollection) (e error) {
	b, e := es.ReadStatus()
	if e != nil {
		return
	}

	s, e := es.ParseStatus(b)
	if e != nil {
		return
	}
	es.CollectMetricsFromStats(c, s)
	return
}

func (es *ElasticSearch) ReadStatus() (b []byte, e error) {
	if len(es.RawStatus) == 0 {
		es.RawStatus, e = FetchURL(es.Url)
		if e != nil {
			return
		}
	}
	b = es.RawStatus
	return
}

func (es *ElasticSearch) ParseStatus(b []byte) (ess *ElasticSearchStatus, e error) {
	ess = &ElasticSearchStatus{}
	e = json.Unmarshal(b, ess)
	if e != nil {
		return nil, e
	}
	return ess, nil
}

func (es *ElasticSearch) CollectMetricsFromStats(mc *MetricsCollection, s *ElasticSearchStatus) {
	mc.Add("shards.Total", s.Shards.Total)
	mc.Add("shards.Successful", s.Shards.Successful)
	mc.Add("shards.Failed", s.Shards.Failed)
	for name, index := range s.Indices {
		tags := map[string]string{"index_name": name}
		mc.AddWithTags("indices.index.SizeInBytes", index.Index.SizeInBytes, tags)
		mc.AddWithTags("indices.index.PrimarySizeInBytes", index.Index.PrimarySizeInBytes, tags)
		mc.AddWithTags("indices.translog.Operations", index.Translog.Operations, tags)
		mc.AddWithTags("indices.docs.NumDocs", index.Docs.NumDocs, tags)
		mc.AddWithTags("indices.docs.MaxDoc", index.Docs.MaxDoc, tags)
		mc.AddWithTags("indices.docs.DeletedDocs", index.Docs.DeletedDocs, tags)
		mc.AddWithTags("indices.merges.Current", index.Merges.Current, tags)
		mc.AddWithTags("indices.merges.CurrentDocs", index.Merges.CurrentDocs, tags)
		mc.AddWithTags("indices.merges.CurrentSizeInBytes", index.Merges.CurrentSizeInBytes, tags)
		mc.AddWithTags("indices.merges.Total", index.Merges.Total, tags)
		mc.AddWithTags("indices.merges.TotalTimeInMillis", index.Merges.TotalTimeInMillis, tags)
		mc.AddWithTags("indices.merges.TotalDocs", index.Merges.TotalDocs, tags)
		mc.AddWithTags("indices.merges.TotalSizeInBytes", index.Merges.TotalSizeInBytes, tags)
		mc.AddWithTags("indices.refresh.Total", index.Refresh.Total, tags)
		mc.AddWithTags("indices.refresh.TotalTimeInMillis", index.Refresh.TotalTimeInMillis, tags)
		mc.AddWithTags("indices.flush.Total", index.Flush.Total, tags)
		mc.AddWithTags("indices.flush.TotalTimeInMillis", index.Flush.TotalTimeInMillis, tags)

	}
	return
}

type ElasticSearchIndexMerges struct {
	Current            int64 `json:"current"`
	CurrentDocs        int64 `json:"current_docs"`
	CurrentSizeInBytes int64 `json:"current_size_in_bytes"`
	Total              int64 `json:"total"`
	TotalTimeInMillis  int64 `json:"total_time_in_millis"`
	TotalDocs          int64 `json:"total_docs"`
	TotalSizeInBytes   int64 `json:"total_size_in_bytes"`
}

type ElasticSearchFlushOrRefresh struct {
	Total             int64 `json:"total"`
	TotalTimeInMillis int64 `json:"total_time_in_millis"`
}

type ElasticSearchDocs struct {
	NumDocs     int64 `json"num_docs"`
	MaxDoc      int64 `json:"max_doc"`
	DeletedDocs int64 `json:"deleted_docs"`
}

type ElasticSearchTranslog struct {
	Operations int64 `json:"operations"`
}

type ElasticSearchIndexIndexStats struct {
	SizeInBytes        int64 `json:"size_in_bytes"`
	PrimarySizeInBytes int64 `json:"primary_size_in_bytes"`
}

type ElasticSearchIndexStats struct {
	Translog ElasticSearchTranslog        `json:"translog"`
	Index    ElasticSearchIndexIndexStats `json:"index"`
	Docs     ElasticSearchDocs            `json:"docs"`
	Merges   ElasticSearchIndexMerges     `json:"merges"`
	Refresh  ElasticSearchFlushOrRefresh  `json:"refresh"`
	Flush    ElasticSearchFlushOrRefresh  `json:"flush"`
}

type ElasticSearchShards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Failed     int64 `json:"failed"`
}

type ElasticSearchStatus struct {
	Shards  ElasticSearchShards                `json:"_shards"`
	Indices map[string]ElasticSearchIndexStats `json:"indices"`
}
