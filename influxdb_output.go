package main

import (
	"fmt"
	"time"
	"net/url"

	influxClient "github.com/influxdb/influxdb/client"
)



func PublishMetricsWithInfluxDB(address string, metrics []*Metric, hostname string) (e error) {

	started := time.Now()
	taken := started.UTC().Unix()
	series := []*influxClient.Series{}
	//TODO, test address componets ...
	u, err := url.Parse("addmissing://"+address)
    if err != nil {
        panic(err)
    }
    p, _ := u.User.Password()

	c, err := influxClient.NewClient(&influxClient.ClientConfig{
		Username: u.User.Username(),
		Password: p,
		Host: u.Host,
		Database: u.Path[1:],
	})
	if err != nil {
		panic(err)
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
			fmt.Printf("ERROR: %s\n",err)
		} else {
			fmt.Printf("sent %d metrics in %.06f to influxdb://%s/%s \n", len(series), time.Now().Sub(started).Seconds(),u.Host,u.Path[1:])
		}
	}
	return nil
}