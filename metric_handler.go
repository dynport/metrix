package main

type MetricHandler struct {
}

type MetricCollector interface {
	Keys() []string
	Prefix() string
	Collect(*MetricsCollection) error
}

func (h* MetricHandler) Collect(c MetricCollector) (all []*Metric, e error) {
	mc := &MetricsCollection{Prefix: c.Prefix()}
	mc.AddSingularMappings(c.Keys())
	if e = c.Collect(mc); e != nil {
		return
	}
	all = mc.Metrics
	return
}

func (h* MetricHandler) Keys(c MetricCollector) (keys []string) {
	for _, m := range c.Keys() {
		keys = append(keys, c.Prefix() + "." + m)
	}
	return
}
