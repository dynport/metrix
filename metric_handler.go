package main

type MetricHandler struct {
}

type MetricCollector interface {
	Prefix() string
	Collect(*MetricsCollection) error
}

func (h *MetricHandler) Collect(c MetricCollector) (Metrics, error) {
	mc := &MetricsCollection{Prefix: c.Prefix()}
	e := c.Collect(mc)
	return mc.Metrics, e
}
