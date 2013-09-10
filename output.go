package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

type OutputHandler struct {
	OpenTSDBAddress string
	GraphiteAddress string
	Hostname        string
}

func (o *OutputHandler) WriteMetrics(all []*Metric) (e error) {
	if o.Hostname == "" {
		hn, e := os.Hostname()
		if e != nil {
			return e
		}
		o.Hostname = hn
	}
	sent := false
	if o.OpenTSDBAddress != "" {
		e = SendMetricsToOpenTSDB(o.OpenTSDBAddress, all, o.Hostname)
		sent = true
	}

	if o.GraphiteAddress != "" {
		e = SendMetricsToGraphite(o.GraphiteAddress, all, o.Hostname)
		sent = true
	}

	if !sent {
		SendMetricsToStdout(all, o.Hostname)
	}
	return
}

func SendMetricsToGraphite(address string, metrics []*Metric, hostname string) (e error) {
	return SendMetricsWith(address, metrics, hostname, func(started time.Time, metric *Metric) string {
		return metric.Graphite(started, hostname)
	},
	)
}

func SendMetricsToOpenTSDB(address string, metrics []*Metric, hostname string) (e error) {
	return SendMetricsWith(address, metrics, hostname, func(started time.Time, metric *Metric) string {
		return metric.OpenTSDB(started, hostname)
	},
	)
}

func SendMetricsWith(address string, metrics []*Metric, hostname string, serializer func(time.Time, *Metric) string) (e error) {
	started := time.Now()
	con, e := net.DialTimeout("tcp", address, 1*time.Second)
	if e != nil {
		return
	}
	defer con.Close()
	fmt.Printf("connected in %.06f\n", time.Now().Sub(started).Seconds())

	started = time.Now()
	debug := os.Getenv("DEBUG") == "true"
	for _, m := range metrics {
		line := serializer(started, m)
		if debug {
			fmt.Println(line)
		}
		fmt.Fprintln(con, line)
	}
	fmt.Printf("sent %d metrics in %.06f\n", len(metrics), time.Now().Sub(started).Seconds())
	return
}

func SendMetricsToStdout(metrics []*Metric, hostname string) {
	now := time.Now()
	for _, m := range metrics {
		fmt.Println(m.Ascii(now, hostname))
	}
}
