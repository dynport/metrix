package main

import (
	"fmt"
	"time"
	"net/url"
	"syscall"

	influxClient "github.com/influxdb/influxdb/client"
)



func PublishMetricsWithInfluxDB(address string, metrics []*Metric, hostname string) (e error) {

	started := time.Now()
    var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	taken := (int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3)
	series := []*influxClient.Series{}
	//TODO, test address componets ...
	u, err := url.Parse("addmissing://"+address)
	if err != nil {
		return err
	}
	p, _ := u.User.Password()

	c, err := influxClient.NewClient(&influxClient.ClientConfig{
		Username: u.User.Username(),
		Password: p,
		Host: u.Host,
		Database: u.Path[1:],
	})
	if err != nil {
		return err
	}
	for _, m := range metrics {
		columns := []string{"time", "value", "host"}
		_points := []interface{}{taken, m.Value, hostname}
		
		for k,v := range m.Tags{
			columns = append(columns,k)
			_points = append(_points,v)
		}
		points := [][]interface{}{_points}
		series = append(series, &influxClient.Series{
				Name:    m.Key,
				Columns: columns,
				Points:  points,
			})
	}
	if len(series) > 0 {
		if err := c.WriteSeries(series); err != nil {
			return err
		} else {
			fmt.Printf("sent %d metrics in %.06f to influxdb://%s/%s \n", len(series), time.Now().Sub(started).Seconds(),u.Host,u.Path[1:])
		}
	}
	return nil
}