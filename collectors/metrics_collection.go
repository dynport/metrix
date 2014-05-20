package collectors

type MetricsCollection struct {
	Prefix  string
	Metrics []*Metric
	Mapping map[string]string
}

func (m *MetricsCollection) AddWithTags(key string, v int64, tags map[string]string) (e error) {
	if m.Prefix != "" {
		key = m.Prefix + "." + key
	}
	m.Metrics = append(m.Metrics, &Metric{Key: key, Value: v, Tags: tags})
	return nil
}

func (m *MetricsCollection) Add(key string, v int64) (e error) {
	return m.AddWithTags(key, v, map[string]string{})
}
