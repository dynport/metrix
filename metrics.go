package main

type MetricsMap map[string]int64

type Metrics []*Metric

func (list Metrics) Len() int {
	return len(list)
}

func (list Metrics) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func (list Metrics) Less(a, b int) bool {
	return list[a].Key < list[b].Key
}
