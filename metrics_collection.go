package main

import "strconv"

type MetricsCollection struct {
	Prefix  string
	Metrics Metrics
	Mapping map[string]string
}

func (m *MetricsCollection) AddSingularMappings(mappings []string) {
	for _, k := range mappings {
		m.AddSingularMapping(k)
	}
}

func (m *MetricsCollection) AddSingularMapping(from string) {
	m.AddMapping(from, from)
}

func (m *MetricsCollection) AddMapping(from, to string) {
	if m.Mapping == nil {
		m.Mapping = map[string]string{}
	}
	m.Mapping[from] = to
}

func (m *MetricsCollection) AddWithTags(key string, v int64, tags map[string]string) (e error) {
	if realKey, ok := m.Mapping[key]; ok {
		key = realKey
	}
	m.Metrics = append(m.Metrics, &Metric{Key: m.Prefix + "." + key, Value: v, Tags: tags})
	return
}

func (m *MetricsCollection) MustAddString(key, value string) error {
	i, e := strconv.ParseInt(value, 10, 64)
	if e != nil {
		return e
	}
	return m.Add(key, i)
}

func (m *MetricsCollection) Add(key string, v int64) (e error) {
	return m.AddWithTags(key, v, map[string]string{})
}
