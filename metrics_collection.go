package main

import (
	"errors"
	"fmt"
)

type MetricsCollection struct {
	Prefix  string
	Metrics []*Metric
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
		m.Metrics = append(m.Metrics, &Metric{Key: m.Prefix + "." + realKey, Value: v, Tags: tags})
	} else {
		e = errors.New("no mapping defined for " + key)
		fmt.Println("ERROR", e.Error())
	}
	return
}

func (m *MetricsCollection) Add(key string, v int64) (e error) {
	return m.AddWithTags(key, v, map[string]string{})
}
