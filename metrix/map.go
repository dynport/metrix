package metrix

import "sort"

type Value struct {
	Key   string
	Value int
}

type Map map[string]*Value

type Values []*Value

func (list Values) TopN(n int) Values {
	sort.Sort(sort.Reverse(list))
	out := make(Values, 0, n)
	for _, v := range list {
		out = append(out, v)
		if len(out) >= n {
			break
		}
	}
	return out
}

func (list Values) Len() int {
	return len(list)
}

func (list Values) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func (list Values) Less(a, b int) bool {
	return list[a].Value < list[b].Value
}

func (m Map) Values() Values {
	values := make(Values, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m Map) Inc(key string) int {
	return m.IncBy(key, 1)
}

func (m Map) IncBy(key string, value int) int {
	if m[key] == nil {
		m[key] = &Value{Key: key}
	}
	m[key].Value += value
	return m[key].Value
}
