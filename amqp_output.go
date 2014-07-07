package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const AMQP_EXCHANGE = "metrix"

func PublishMetricsWithAMQP(address string, metrics []*Metric, hostname string) (e error) {
	con, e := amqp.Dial("amqp://" + address)
	if e != nil {
		return e
	}
	defer con.Close()
	channel, e := con.Channel()
	if e != nil {
		return e
	}
	defer channel.Close()
	e = channel.ExchangeDeclare(AMQP_EXCHANGE, "fanout", false, false, false, false, nil)
	if e != nil {
		return e
	}
	started := time.Now()
	for _, m := range metrics {
		amqpKey := hostname + "." + m.Key
		b, e := json.Marshal(m)
		if e != nil {
			logger.Printf(e.Error())
			continue
		}
		e = channel.Publish(AMQP_EXCHANGE, amqpKey, false, false, amqp.Publishing{
			Body:        b,
			ContentType: "application/json",
			Timestamp:   time.Now(),
		})
		if e != nil {
			logger.Printf("ERROR: error publishing " + e.Error())
		}
	}
	fmt.Printf("sent %d metrics in %.06f\n", len(metrics), time.Now().Sub(started).Seconds())
	return nil
}
