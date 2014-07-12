package main

import (
	"strconv"
	"strings"
)

// Parses the self documenting format of the linux net statistics interface.
// The interface is in /proc/net and contains two consecutive lines starting
// with the same prefix like "prefix:". The first line contains the header
// and the second line the actual signed 64 bit values.
// A value < 0 means, that this statistic is not supported.
func parse2lines(headers, values string) map[string]int64 {
	keys := strings.Fields(headers)
	vals := strings.Fields(values)
	result := make(map[string]int64, len(keys))

	if len(keys) != len(vals) || len(keys) <= 1 || keys[0] != vals[0] {
		return result
	}

	// strip the ":" of "foo:" ...
	topic := keys[0][:len(keys[0])-1]
	// .. and just get the actual header entries and values
	keys = keys[1:]
	vals = vals[1:]

	for i, k := range keys {
		if v, e := strconv.ParseInt(vals[i], 10, 64); e == nil && v >= 0 {
			result[topic+"."+k] = v
		}
	}
	return result
}
