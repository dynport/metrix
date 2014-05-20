package collectors

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func readNetstats() (metrics []*Metric, e error) {
	for _, name := range []string{"net/snmp", "net/netstat"} {
		f, e := OpenProcFile(name)
		if e != nil {
			return nil, e
		}
		defer f.Close()
		metrics = append(metrics, parseNetstats(f)...)
	}
	return metrics, nil
}

func parseNetstats(r io.Reader) []*Metric {
	metrics := []*Metric{}
	scanner := bufio.NewScanner(r)
	line := 0
	fields := []string{}
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		if len(tmp) < 2 {
			continue
		}

		if line%2 == 0 {
			fields = tmp
		} else {
			for i, value := range tmp[1:] {
				prefix := strings.TrimSuffix(fields[0], ":")
				field := prefix + "." + fields[i+1]
				i, e := strconv.ParseInt(value, 10, 64)
				if e == nil {
					metrics = append(metrics, &Metric{Key: field, Value: i})
				}
			}
		}
		line++
	}
	return metrics
}
