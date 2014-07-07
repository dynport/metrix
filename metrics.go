package main

import (
	"fmt"
	"sort"
	"strings"
)

type MetricsMap map[string]int64

type Metrics []*Metric

func (list Metrics) Len() int {
	return len(list)
}

func (list Metrics) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func (list Metrics) Less(a, b int) bool {
	if list[a].Key == list[b].Key {
		return flattenTags(list[a].Tags) < flattenTags(list[b].Tags)
	} else {
		return list[a].Key < list[b].Key
	}
}

func flattenTags(tags map[string]string) string {
	out := []string{}
	for k, v := range tags {
		out = append(out, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(out)
	return strings.Join(out, " ")
}
