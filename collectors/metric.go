package collectors

type Metric struct {
	Key   string
	Value int64
	Tags  map[string]string
}
