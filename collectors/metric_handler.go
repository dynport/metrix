package collectors

type MetricHandler struct {
}

type MetricCollector interface {
	Prefix() string
	Collect(*MetricsCollection) error
}

func (h *MetricHandler) Collect(c MetricCollector) (all []*Metric, e error) {
	mc := &MetricsCollection{Prefix: c.Prefix()}
	if e = c.Collect(mc); e != nil {
		return
	}
	all = mc.Metrics
	return all, nil
}
